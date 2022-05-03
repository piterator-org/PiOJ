package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/go-redis/redis/v8"
	pioj "github.com/piterator-org/pioj/app"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DatabaseConfiguration struct {
	MongoDB struct {
		URI string
	}
	Redis *redis.Options
}

type UserConfiguration struct {
	pioj.Configuration
	Database DatabaseConfiguration
}

func main() {
	var config UserConfiguration

	if len(os.Args) > 1 {
		if _, err := toml.DecodeFile(os.Args[1], &config); err != nil {
			panic(err)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.Database.MongoDB.URI))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
	database := client.Database("pioj")

	rdb := redis.NewClient(config.Database.Redis)

	http.ListenAndServe(":8080", pioj.NewApp(config.Configuration, database, rdb).ServeMux)
}
