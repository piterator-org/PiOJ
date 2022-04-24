package pioj

import (
	"context"
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IOExample [2]string

type ProblemTag string

type TestCase [2]string

type Subtask []TestCase

type Problem struct {
	ID           int              `json:"id"`
	Title        LocalizedStrings `json:"title"`
	Difficulty   int              `json:"difficulty"`
	InputFile    string           `json:"input_file"`
	OutputFile   string           `json:"output_file"`
	TimeLimit    int              `json:"time_limit"`
	MemoryLimit  int              `json:"memory_limit"`
	Background   LocalizedStrings `json:"background"`
	Description  LocalizedStrings `json:"description"`
	InputFormat  LocalizedStrings `json:"input_format"`
	OutputFormat LocalizedStrings `json:"output_format"`
	Examples     []IOExample      `json:"examples"`
	Hints        LocalizedStrings `json:"hints"`
	Tags         []ProblemTag     `json:"tags"`
	Subtasks     []Subtask        `json:"subtasks"`
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
			w.Header().Set("Content-Type", "application/json")
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
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(problem)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}
