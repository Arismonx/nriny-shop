package inventory

import (
	"github.com/Arismonx/nriny-shop/modules/item"
	"github.com/Arismonx/nriny-shop/modules/model"
)

type (
	InventoryUpdateReq struct {
		PlayerId string `json:"player_id" validate:"required,max=64"`
		ItemId   string `json:"item_id" validate:"required,max=64"`
	}

	ItemInInventory struct {
		InventoryId string `json:"inventory_id"`
		*item.ItemShowCase
	}

	PlyerInventory struct {
		PlayerId string `json:"player_id"`
		*model.PaginateRes
	}
)
