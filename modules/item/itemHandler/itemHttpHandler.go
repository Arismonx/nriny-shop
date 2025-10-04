package itemHandler

import (
	"github.com/Arismonx/nriny-shop/config"
	"github.com/Arismonx/nriny-shop/modules/item/itemUsecase"
)

type (
	ItemHttpHandlerService interface{}

	itemHttpHandler struct {
		cfg         *config.Config
		itemUsecase itemUsecase.ItemUsecaseService
	}
)

func NewItemHttpHandler(
	cfg *config.Config,
	itemUsecase itemUsecase.ItemUsecaseService,
) ItemHttpHandlerService {
	return &itemHttpHandler{cfg, itemUsecase}
}
