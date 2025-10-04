package paymentHandler

import (
	"github.com/Arismonx/nriny-shop/config"
	"github.com/Arismonx/nriny-shop/modules/payment/paymentUsecase"
)

type (
	PaymentHttpHandlerService interface{}

	paymentHttpHandler struct {
		cfg            *config.Config
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentHttpHandler(
	cfg *config.Config,
	paymentUsecase paymentUsecase.PaymentUsecaseService,
) PaymentHttpHandlerService {
	return &paymentHttpHandler{cfg, paymentUsecase}
}
