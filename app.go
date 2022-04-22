package main

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Problem struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type App struct {
	ServeMux *http.ServeMux
	Database *mongo.Database
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

func (app App) Handle() *http.ServeMux {
	fs := wrapHandler(http.FileServer(http.Dir("dist/")), "dist/index.html", "text/html; charset=utf-8")
	app.ServeMux.Handle("/", http.StripPrefix("/", fs))

	app.ServeMux.HandleFunc("/api/problem/create", func(w http.ResponseWriter, r *http.Request) {
		var problem Problem
		if err := json.NewDecoder(r.Body).Decode(&problem); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var last Problem
		switch err := app.Database.Collection("problem").FindOne(
			context.TODO(), bson.D{}, options.FindOne().SetSort(map[string]int{"id": -1}),
		).Decode(&last); err {
		case mongo.ErrNoDocuments:
			problem.ID = 1
		case nil:
			problem.ID = last.ID + 1
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if _, err := app.Database.Collection("problem").InsertOne(context.TODO(), problem); err != nil {
			http.Error(w, err.Error(), http.StatusInsufficientStorage)
		} else {
			json.NewEncoder(w).Encode(problem)
		}
	})

	app.ServeMux.HandleFunc("/api/problem/get", func(w http.ResponseWriter, r *http.Request) {
		var body struct{ ID int }
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var problem Problem
		switch err := app.Database.Collection("problem").FindOne(
			context.TODO(), bson.D{{Key: "id", Value: body.ID}},
		).Decode(&problem); err {
		case mongo.ErrNoDocuments:
			http.Error(w, err.Error(), http.StatusNotFound)
		case nil:
			json.NewEncoder(w).Encode(problem)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
	return app.ServeMux
}
