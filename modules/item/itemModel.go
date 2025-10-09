package item

import "github.com/Arismonx/nriny-shop/modules/model"

type (
	CreateItemReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validare:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=256"`
		Damage   int     `json:"damage" validate:"required"`
	}

	ItemShowCase struct {
		ItemId   string  `json:"item_id"`
		Title    string  `json:"title"`
		Price    float64 `json:"price"`
		ImageUrl string  `json:"image_url"`
		Damage   int     `json:"damage"`
	}

	ItemSearchReq struct {
		Title string `json:"title"`
		model.PaginateReq
	}

	ItemUpdateReq struct {
		Title    string  `json:"title" validate:"required,max=64"`
		Price    float64 `json:"price" validate:"required"`
		ImageUrl string  `json:"image_url" validate:"required,max=256"`
		Damage   int     `json:"damage" validate:"required"`
	}

	EnableOrDisableItemReq struct {
		UsageStatus bool `json:"status"`
	}
)
