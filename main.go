package main

import (
	"context"
	"net/http"
	"time"

	pioj "github.com/piterator-org/pioj/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	database := client.Database("pioj")

	http.ListenAndServe(":8080", pioj.NewApp(database).ServeMux)
}
