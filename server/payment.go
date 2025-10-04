package server

import (
	"github.com/Arismonx/nriny-shop/modules/payment/paymentHandler"
	"github.com/Arismonx/nriny-shop/modules/payment/paymentRepository"
	"github.com/Arismonx/nriny-shop/modules/payment/paymentUsecase"
)

func (s *server) paymentService() {
	repo := paymentRepository.NewPaymentRepository(s.db)
	usecase := paymentUsecase.NewPaymentUsecase(repo)
	httpHandler := paymentHandler.NewPaymentHttpHandler(s.cfg, usecase)
	grpcHandler := paymentHandler.NewPaymentGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	payment := s.app.Group("/payment_v1")

	// health check
	_ = payment
}
