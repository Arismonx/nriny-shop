package main

import (
	"context"
	"log"
	"os"

	"github.com/Arismonx/nriny-shop/config"
	"github.com/Arismonx/nriny-shop/pkg/database"
)

func main() {
	ctx := context.Background()

	//init config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}

		return os.Args[1]
	}())

	//database connecting
	db := database.ConnectingMongoDB(ctx, &cfg)
	log.Println(db)

	defer func() {
		if err := db.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}
