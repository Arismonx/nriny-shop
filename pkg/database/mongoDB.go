package database

import (
	"context"
	"log"
	"time"

	"github.com/Arismonx/nriny-shop/config"
	// "go.mongodb.org/mongo-driver/v2/mongo"
	// "go.mongodb.org/mongo-driver/v2/mongo/options"
	// "go.mongodb.org/mongo-driver/v2/mongo/readpref"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectingMongoDB(pctx context.Context, cfg *config.Config) *mongo.Client {
	// client, err := mongo.Connect(options.Client().ApplyURI(cfg.Db.DB_Url))

	client, err := mongo.Connect(pctx, options.Client().ApplyURI(cfg.Db.DB_Url))

	if err != nil {
		log.Fatal("Error: connecting database :", err)
	}

	ctx, cancel := context.WithTimeout(pctx, 2*time.Second)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal("Error: cannot pinging! :", err)
	}

	return client
}
