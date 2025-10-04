package paymentHandler

import (
	"github.com/Arismonx/nriny-shop/modules/payment/paymentUsecase"
)

type (
	paymentGrpcHandler struct {
		paymentUsecase paymentUsecase.PaymentUsecaseService
	}
)

func NewPaymentGrpcHandler(
	paymentUsecase paymentUsecase.PaymentUsecaseService,
) *paymentGrpcHandler {
	return &paymentGrpcHandler{paymentUsecase}
}
