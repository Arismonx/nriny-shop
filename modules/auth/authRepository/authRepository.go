package authRepository

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

// Repository Pattern
type (
	// Polymorphism
	AuthRepositoryService interface {
		// AuthConnDB(pctx context.Context) *mongo.Database
	}

	// Class Concreat
	authRepository struct {
		db *mongo.Client
	}
)

// Constructor
func NewAuthRepository(db *mongo.Client) AuthRepositoryService {
	return &authRepository{db}
}

// Connect Database Auth
func (r *authRepository) authConnDB(pctx context.Context) *mongo.Database {
	return r.db.Database("auth-db")
}
