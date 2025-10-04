package authHandler

import (
	"github.com/Arismonx/nriny-shop/modules/auth/authUsecase"
)

type (
	authGrpcHandler struct {
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthGrpcHandler(
	authUsecase authUsecase.AuthUsecaseService,
) *authGrpcHandler {
	return &authGrpcHandler{authUsecase}
}
