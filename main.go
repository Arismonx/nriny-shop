package main

import (
	"context"
	"log"
	"os"

	"github.com/Arismonx/nriny-shop/config"
)

func main() {
	ctx := context.Background()
	_ = ctx

	cfg := config.LoadConfig(func() string {
		if len(os.Args) < 2 {
			log.Fatal("Error: .env path is required")
		}

		return os.Args[1]
	}())

	log.Println(cfg)
}
