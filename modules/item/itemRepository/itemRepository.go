package itemRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	ItemRepositoryService interface{}

	itemRepository struct {
		db *mongo.Client
	}
)

func NewItemRepository(db *mongo.Client) ItemRepositoryService {
	return itemRepository{db}
}

func (r *itemRepository) itemConnDB(pctx context.Context) *mongo.Database {
	return r.db.Database("item-db")
}
