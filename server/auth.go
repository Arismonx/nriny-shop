package server

import (
	"log"

	"github.com/Arismonx/nriny-shop/modules/auth/authHandler"
	"github.com/Arismonx/nriny-shop/modules/auth/authRepository"
	"github.com/Arismonx/nriny-shop/modules/auth/authUsecase"
	grpcconnect "github.com/Arismonx/nriny-shop/pkg/grpcConnect"

	authPb "github.com/Arismonx/nriny-shop/modules/auth/authPb"
)

func (s *server) authService() {
	repo := authRepository.NewAuthRepository(s.db)
	usecase := authUsecase.NewAuthUsecase(repo)
	httpHandler := authHandler.NewAuthHttpHandler(s.cfg, usecase)
	grpcHandler := authHandler.NewAuthGrpcHandler(usecase)

	// Grpc Client
	go func() {
		grpcServer, lis := grpcconnect.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.GrpcAuthUrl)

		authPb.RegisterAuthGrpcServiceServer(grpcServer, grpcHandler)
		
		log.Printf("Auth gRPC server listening on %s", s.cfg.Grpc.GrpcAuthUrl)
		grpcServer.Serve(lis)

	}()

	_ = httpHandler
	_ = grpcHandler

	auth := s.app.Group("/auth_v1")

	// health check
	auth.GET("", s.healthCheckService)
}
