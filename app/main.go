package pioj

import (
	"context"
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LocalizedStrings map[string]string

type SMTPConfiguration struct {
	From     string
	Username string
	Password string
	Host     string
	Port     int
}

type Configuration struct {
	SMTP SMTPConfiguration
}

type App struct {
	Configuration
	ServeMux *http.ServeMux
	Database *mongo.Database
	Redis    *redis.Client
	Root     string
	Fallback string
}

func SetUnique(ctx context.Context, collection *mongo.Collection, field string) (string, error) {
	return collection.Indexes().CreateOne(
		ctx,
		mongo.IndexModel{
			Keys:    bson.D{{Key: field, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
}

func SendMail(config SMTPConfiguration, to []string, subject string, body string) error {
	auth := smtp.PlainAuth("", config.Username, config.Password, config.Host)
	err := smtp.SendMail(
		fmt.Sprint(config.Host, ":", config.Port),
		auth,
		config.Username,
		to,
		[]byte(fmt.Sprintf("From: %s\r\nSubject: %s\r\n\r\nHello, world!", config.From, "Hello")),
	)
	return err
}

type NotFoundFallbackRespWr struct {
	http.ResponseWriter // We embed http.ResponseWriter
	status              int
}

func (w *NotFoundFallbackRespWr) WriteHeader(status int) {
	w.status = status // Store the status for our own use
	if status != http.StatusNotFound {
		w.ResponseWriter.WriteHeader(status)
	}
}

func (w *NotFoundFallbackRespWr) Write(p []byte) (int, error) {
	if w.status != http.StatusNotFound {
		return w.ResponseWriter.Write(p)
	}
	return len(p), nil // Lie that we successfully written it
}

func wrapHandler(h http.Handler, fallback string, contentType string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		nfrw := &NotFoundFallbackRespWr{ResponseWriter: w}
		h.ServeHTTP(nfrw, r)
		if nfrw.status == http.StatusNotFound {
			w.Header().Set("Content-Type", contentType)
			http.ServeFile(w, r, fallback)
		}
	}
}

func (app App) Handle() {
	fs := wrapHandler(http.FileServer(http.Dir(app.Root)), app.Fallback, "text/html; charset=utf-8")
	app.ServeMux.Handle("/", http.StripPrefix("/", fs))

	app.HandleProblems()
	app.HandleUsers()
}

func NewApp(config Configuration, database *mongo.Database, rdb *redis.Client) App {
	app := App{
		Configuration: config,
		Database:      database,
		Redis:         rdb,
		ServeMux:      http.NewServeMux(),
		Root:          "./dist/",
		Fallback:      "./dist/index.html",
	}
	app.Handle()
	return app
}
