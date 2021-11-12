package db

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"news/helper"
	"time"
)

func InitMongo() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(helper.GetEnv("MONGO_URI", "mongodb://207.148.76.65:27017")))
	if err != nil {
		log.Println(err)
		panic(err)
	}

	if err := client.Ping(ctx, nil); err != nil {
		log.Println(err)
		panic(err)
	}

	return client.Database("news")
}
