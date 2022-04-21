package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestProblem(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	database := client.Database("test")
	database.Drop(ctx)
	mux := http.NewServeMux()
	App{mux, database}.Handle()

	request := func(path string, body []byte, i int) (Problem, error) {
		req := httptest.NewRequest(http.MethodPost, path, bytes.NewBuffer(body))
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)
		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			return Problem{}, fmt.Errorf("[%d] Status code of %s: %d", i, path, resp.StatusCode)
		}
		var problem Problem
		if err := json.NewDecoder(resp.Body).Decode(&problem); err != nil {
			return problem, errors.New(err.Error())
		}
		if problem.ID != i {
			return problem, fmt.Errorf("[%d] Unexpected problem ID: %d", i, problem.ID)
		}
		return problem, nil
	}

	n := 3

	t.Run("create", func(t *testing.T) {
		for i := 1; i <= n; i++ {
			body, _ := json.Marshal(map[string]string{"title": fmt.Sprint("Problem ", i), "content": "Content"})
			if _, err := request("/problem/create", body, i); err != nil {
				t.Error(err.Error())
			}
		}
	})

	t.Run("get", func(t *testing.T) {
		for i := 1; i <= n; i++ {
			body, _ := json.Marshal(map[string]int{"id": i})
			if problem, err := request("/problem/get", body, i); err != nil {
				t.Error(err.Error())
			} else {
				if problem.Title != fmt.Sprint("Problem ", i) {
					t.Errorf("[%d] Unexpected problem title: %s", i, problem.Title)
				}
			}
		}
	})
}
