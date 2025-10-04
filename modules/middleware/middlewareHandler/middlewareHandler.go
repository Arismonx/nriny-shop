package middlewareHandler

import (
	"github.com/Arismonx/nriny-shop/config"
	"github.com/Arismonx/nriny-shop/modules/middleware/middlewareUsecase"
)

type (
	MiddlewareHandlerService interface{}

	middlewareHandler struct {
		cfg               *config.Config
		MiddlewareUsecase middlewareUsecase.MiddlewareUsecaseService
	}
)

func NewMiddlewareHandler(
	cfg *config.Config,
	MiddlewareUsecase middlewareUsecase.MiddlewareUsecaseService,
) MiddlewareHandlerService {
	return &middlewareHandler{cfg, MiddlewareUsecase}
}
