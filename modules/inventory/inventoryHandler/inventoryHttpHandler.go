package inventoryHandler

import (
	"github.com/Arismonx/nriny-shop/config"
	"github.com/Arismonx/nriny-shop/modules/inventory/inventoryUsecase"
)

type (
	InventoryHttpHandlerService interface{}

	inventoryHttpHandler struct {
		cfg                 *config.Config
		inventoryRepository inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryHttpHandler(
	cfg *config.Config,
	inventoryRepository inventoryUsecase.InventoryUsecaseService,
) InventoryHttpHandlerService {
	return &inventoryHttpHandler{cfg, inventoryRepository}
}
