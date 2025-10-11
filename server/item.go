package server

import (
	"log"

	"github.com/Arismonx/nriny-shop/modules/item/itemHandler"
	itemPb "github.com/Arismonx/nriny-shop/modules/item/itemPb"
	"github.com/Arismonx/nriny-shop/modules/item/itemRepository"
	"github.com/Arismonx/nriny-shop/modules/item/itemUsecase"
	grpcconnect "github.com/Arismonx/nriny-shop/pkg/grpcConnect"
)

func (s *server) itemService() {
	repo := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repo)
	httpHandler := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)

	// Grpc Client
	go func() {
		grpcServer, lis := grpcconnect.NewGrpcServer(&s.cfg.Jwt, s.cfg.Grpc.GrpcItemUrl)

		itemPb.RegisterItemGrpcServiceServer(grpcServer, grpcHandler)

		log.Printf("Item gRPC server listening on %s", s.cfg.Grpc.GrpcItemUrl)
		grpcServer.Serve(lis)

	}()

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	// health check
	item.GET("", s.healthCheckService)
}
