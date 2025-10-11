package server

import (
	"log"

	"github.com/Arismonx/nriny-shop/modules/inventory/inventoryHandler"
	inventoryPb "github.com/Arismonx/nriny-shop/modules/inventory/inventoryPb"
	"github.com/Arismonx/nriny-shop/modules/inventory/inventoryRepository"
	"github.com/Arismonx/nriny-shop/modules/inventory/inventoryUsecase"
	grpcconnect "github.com/Arismonx/nriny-shop/pkg/grpcConnect"
)

func (s *server) inventoryService() {
	repo := inventoryRepository.NewInventoryRepository(s.db)
	usecase := inventoryUsecase.NewInventoryUsecase(repo)
	httpHandler := inventoryHandler.NewInventoryHttpHandler(s.cfg, usecase)
	grpcHandler := inventoryHandler.NewInventoryGrpcHandler(usecase)
	queueHandler := inventoryHandler.NewInventoryQueueHandler(s.cfg, usecase)

	// Grpc Client
	go func() {
		grpcServer, lis := grpcconnect.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.GrpcInventoryUrl)

		inventoryPb.RegisterInventoryGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Inventory gRPC server listening on %s", s.cfg.Grpc.GrpcInventoryUrl)
		grpcServer.Serve(lis)

	}()

	_ = httpHandler
	_ = grpcHandler
	_ = queueHandler

	inventory := s.app.Group("/inventory_v1")

	// health check
	inventory.GET("", s.healthCheckService)
}
