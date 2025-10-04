package inventoryHandler

import (
	"github.com/Arismonx/nriny-shop/config"
	"github.com/Arismonx/nriny-shop/modules/inventory/inventoryUsecase"
)

type (
	InventoryQueueHandler interface{}

	inventoryQueueHandler struct {
		cfg              *config.Config
		inventoryUsecase inventoryUsecase.InventoryUsecaseService
	}
)

func NewInventoryQueueHandler(
	cfg *config.Config,
	inventoryUsecase inventoryUsecase.InventoryUsecaseService,
) InventoryQueueHandler {
	return &inventoryQueueHandler{cfg, inventoryUsecase}
}
