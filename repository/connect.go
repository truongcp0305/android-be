package repository

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connn() *mongo.Database {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	u := "mongodb+srv://nguyentruongcp35:123456@cluster0.zrvxwix.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	opts := options.Client().ApplyURI(u).SetServerAPIOptions(serverAPI)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}
	database := client.Database("product")
	return database
}
