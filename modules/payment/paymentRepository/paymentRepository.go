package paymentRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type (
	PaymentRepositoryService interface{}

	paymentRepository struct {
		db *mongo.Client
	}
)

func NewPaymentRepository(
	db *mongo.Client,
) PaymentRepositoryService {
	return &paymentRepository{db}
}

func (r *paymentRepository) paymentConnDB(pctx context.Context) *mongo.Database {
	return r.db.Database("payment-db")
}
