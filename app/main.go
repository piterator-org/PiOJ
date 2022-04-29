package pioj

import (
	"context"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LocalizedStrings map[string]string

type App struct {
	ServeMux *http.ServeMux
	Database *mongo.Database
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
}

func NewApp(database *mongo.Database) App {
	app := App{
		Database: database,
		ServeMux: http.NewServeMux(),
		Root:     "./dist/",
		Fallback: "./dist/index.html",
	}
	app.Handle()
	return app
}
