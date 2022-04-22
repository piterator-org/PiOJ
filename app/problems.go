package pioj

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

func (app App) HandleProblems() {
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
}
