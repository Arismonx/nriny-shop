package authHandler

import (
	"github.com/Arismonx/nriny-shop/config"
	"github.com/Arismonx/nriny-shop/modules/auth/authUsecase"
)

type (
	AuthHttpHandlerService interface{}

	authHttpHandler struct {
		cfg         *config.Config
		authUsecase authUsecase.AuthUsecaseService
	}
)

func NewAuthHttpHandler(
	cfg *config.Config,
	authUsecase authUsecase.AuthUsecaseService,
) AuthHttpHandlerService {
	return &authHttpHandler{cfg, authUsecase}
}
