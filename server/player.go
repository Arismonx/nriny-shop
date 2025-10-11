package server

import (
	"log"

	"github.com/Arismonx/nriny-shop/modules/player/playerHandler"
	playerPb "github.com/Arismonx/nriny-shop/modules/player/playerPb"
	"github.com/Arismonx/nriny-shop/modules/player/playerRepository"
	"github.com/Arismonx/nriny-shop/modules/player/playerUsecase"
	grpcconnect "github.com/Arismonx/nriny-shop/pkg/grpcConnect"
)

func (s *server) playerService() {
	repo := playerRepository.NewPlayerRepository(s.db)
	usecase := playerUsecase.NewPlayerUsecase(repo)
	httpHandler := playerHandler.NewPlayerHttpHandler(s.cfg, usecase)
	grpcHandler := playerHandler.NewPlayerGrpcHandler(usecase)
	queueHandler := playerHandler.NewPlayerQueueHandler(s.cfg, usecase)

	// Grpc Client
	go func() {
		grpcServer, lis := grpcconnect.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.GrpcPlayerUrl)

		playerPb.RegisterPlayerGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Player gRPC server listening on %s", s.cfg.Grpc.GrpcPlayerUrl)
		grpcServer.Serve(lis)

	}()

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	player := s.app.Group("/player_v1")

	// health check
	player.GET("", s.healthCheckService)
}
