package main

import (
	"context"
	"log"
	"os"

	"github.com/Arismonx/nriny-shop/config"
	"github.com/Arismonx/nriny-shop/pkg/database/migration"
)

func main() {
	ctx := context.Background()
	_ = ctx

	//init config
	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}

		return os.Args[1]
	}())

	switch cfg.App.AppName {
	case "auth":
		migration.AuthMigrate(ctx, &cfg)
	case "item":
	case "inventory":
	case "payment":
	case "player":
	}
}
