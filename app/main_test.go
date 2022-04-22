package pioj

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestFileServer(t *testing.T) {
	os.Mkdir("../dist", os.FileMode(0755))
	if file, err := os.OpenFile("../dist/index.html", os.O_RDWR|os.O_CREATE, os.FileMode(0644)); err != nil {
		t.Error(err.Error())
	} else if fi, _ := os.Stat("../dist/index.html"); fi.Size() == 0 {
		file.WriteString("\n")
	}

	mux := http.NewServeMux()
	App{ServeMux: mux, Root: "../dist/", Fallback: "../dist/index.html"}.Handle()

	t.Run("GET/", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)
		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Unexpected status code at %s: %d", req.URL.Path, resp.StatusCode)
		}
	})

	t.Run("GET/404", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/404", nil)
		w := httptest.NewRecorder()
		mux.ServeHTTP(w, req)
		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Unexpected status code at %s: %d", req.URL.Path, resp.StatusCode)
		}
		contentType := resp.Header.Get("Content-Type")
		if contentType != "text/html; charset=utf-8" {
			t.Errorf("Unexpected Content-Type at %s: %s", req.URL.Path, contentType)
		}
	})
}

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
	mux := NewApp(database).ServeMux

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
			return Problem{}, fmt.Errorf("[%d] Unexpected status code at %s: %d", i, path, resp.StatusCode)
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
			if _, err := request(
				"/api/problem/create",
				map[string]string{"title": fmt.Sprint("Problem ", i), "content": "Content"},
				i,
			); err != nil {
				t.Error(err.Error())
			}
		}
	})

	t.Run("create 400", func(t *testing.T) {
		resp := post("/api/problem/create", bytes.NewBufferString("{"))
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Unexpected status code: %d", resp.StatusCode)
		}
	})

	t.Run("get", func(t *testing.T) {
		for i := 1; i <= n; i++ {
			if problem, err := request("/api/problem/get", map[string]int{"id": i}, i); err != nil {
				t.Error(err.Error())
			} else {
				if problem.Title != fmt.Sprint("Problem ", i) {
					t.Errorf("[%d] Unexpected problem title: %s", i, problem.Title)
				}
			}
		}
	})

	t.Run("get 404", func(t *testing.T) {
		resp := postjson("/api/problem/get", map[string]int{"id": n + 1})
		if resp.StatusCode != http.StatusNotFound {
			t.Errorf("Unexpected status code: %d", resp.StatusCode)
		}
	})

	t.Run("get 400", func(t *testing.T) {
		resp := post("/api/problem/get", bytes.NewBufferString("{"))
		if resp.StatusCode != http.StatusBadRequest {
			t.Errorf("Unexpected status code: %d", resp.StatusCode)
		}
	})
}
