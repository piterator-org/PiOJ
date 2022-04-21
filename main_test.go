package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
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

	post := func(path string, body io.Reader) *http.Response {
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, httptest.NewRequest(http.MethodPost, path, body))
		return w.Result()
	}

	postjson := func(path string, data any) *http.Response {
		body, _ := json.Marshal(data)
		return post(path, bytes.NewBuffer(body))
	}

	request := func(path string, data any, i int) (Problem, error) {
		resp := postjson(path, data)
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
			if _, err := request("/problem/create", map[string]string{"title": fmt.Sprint("Problem ", i), "content": "Content"}, i); err != nil {
				t.Error(err.Error())
			}
		}
	})

	t.Run("create 400", func(t *testing.T) {
		resp := post("/problem/create", bytes.NewBufferString("{"))
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Unexpected status code: %d", resp.StatusCode)
		}
	})

	t.Run("get", func(t *testing.T) {
		for i := 1; i <= n; i++ {
			if problem, err := request("/problem/get", map[string]int{"id": i}, i); err != nil {
				t.Error(err.Error())
			} else {
				if problem.Title != fmt.Sprint("Problem ", i) {
					t.Errorf("[%d] Unexpected problem title: %s", i, problem.Title)
				}
			}
		}
	})

	t.Run("get 404", func(t *testing.T) {
		resp := postjson("/problem/get", map[string]int{"id": n + 1})
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Unexpected status code: %d", resp.StatusCode)
		}
	})

	t.Run("get 400", func(t *testing.T) {
		resp := post("/problem/get", bytes.NewBufferString("{"))
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Unexpected status code: %d", resp.StatusCode)
		}
	})
}
