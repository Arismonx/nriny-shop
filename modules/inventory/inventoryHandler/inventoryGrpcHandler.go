package inventoryHandler

import (
	"context"

	inventoryPb "github.com/Arismonx/nriny-shop/modules/inventory/inventoryPb"
	"github.com/Arismonx/nriny-shop/modules/inventory/inventoryUsecase"
)

type (
	inventoryGrpcHandler struct {
		inventoryPb.UnimplementedInventoryGrpcServiceServer
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryGrpcHandler(
	inventoryUsecase inventoryUsecase.InventoryUsecaseService,
) *inventoryGrpcHandler {
	return &inventoryGrpcHandler{
		inventoryUsecase: inventoryUsecase,
	}
}

func (g *inventoryGrpcHandler) IsAvailableToSell(ctx context.Context, req *inventoryPb.IsAvaibleToSellReq) (*inventoryPb.IsAvaibleToSellRes, error) {
	return nil, nil
}
