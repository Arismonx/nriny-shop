package server

import (
	"github.com/Arismonx/nriny-shop/modules/item/itemHandler"
	"github.com/Arismonx/nriny-shop/modules/item/itemRepository"
	"github.com/Arismonx/nriny-shop/modules/item/itemUsecase"
)

func (s *server) itemService() {
	repo := itemRepository.NewItemRepository(s.db)
	usecase := itemUsecase.NewItemUsecase(repo)
	httpHandler := itemHandler.NewItemHttpHandler(s.cfg, usecase)
	grpcHandler := itemHandler.NewItemGrpcHandler(usecase)

	_ = httpHandler
	_ = grpcHandler

	item := s.app.Group("/item_v1")

	// health check
	item.GET("", s.healthCheckService)
}
