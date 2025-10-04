package playerRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	PlayerRepositoryService interface{}

	playerRepository struct {
		db *mongo.Client
	}
)

func NewPlayerRepository(
	db *mongo.Client,
) PlayerRepositoryService {
	return &playerRepository{db}
}

func (r *playerRepository) playerConnDB(pctx context.Context) *mongo.Database {
	return r.db.Database("player-db")
}
