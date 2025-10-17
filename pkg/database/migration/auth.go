package migration

import (
	"context"
	"log"

	"github.com/Arismonx/nriny-shop/config"
	"github.com/Arismonx/nriny-shop/modules/auth"
	"github.com/Arismonx/nriny-shop/pkg/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func authDbConn(pctx context.Context, cfg *config.Config) *mongo.Database {
	return database.ConnectingMongoDB(pctx, cfg).Database("auth_db")
}

func AuthMigrate(pctx context.Context, cfg *config.Config) {
	db := authDbConn(pctx, cfg)
	defer func() {
		if err := db.Client().Disconnect(pctx); err != nil {
			log.Printf("Error disconnecting from MongoDB: %v", err)
		}
	}()

	log.Println("Starting auth migration...")

	// collect auth
	collection_auth := db.Collection("auth")
	indexs_auth, err := collection_auth.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"player_id", 1}}},
		{Keys: bson.D{{"refresh_token", 1}}},
	})

	if err != nil {
		log.Fatalf("Failed to create indexes for 'auth' collection: %v", err)
	}
	log.Printf("Created %d indexes for 'auth' collection.", len(indexs_auth))

	//collect role
	collection_role := db.Collection("roles")
	indexs_role, _ := collection_role.Indexes().CreateMany(pctx, []mongo.IndexModel{
		{Keys: bson.D{{"_id", 1}}},
		{Keys: bson.D{{"code", 1}}},
	})

	if err != nil {
		log.Fatalf("Failed to create indexes for 'role' collection: %v", err)
	}
	log.Printf("Created %d indexes for 'role' collection.", len(indexs_role))

	// role data
	document := func() []any {
		roles := []*auth.Role{
			{
				Title: "user",
				Code:  0,
			},
			{
				Title: "admin",
				Code:  1,
			},
		}
		docs := make([]any, 0)

		for _, r := range roles {
			docs = append(docs, r)
		}
		return docs
	}()

	result, err := collection_role.InsertMany(pctx, document, nil)

	if err != nil {
		panic(err)
	}

	log.Println("Migrate collect complete:", result)

}
