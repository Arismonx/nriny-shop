package playerHandler

import (
	"github.com/Arismonx/nriny-shop/config"
	"github.com/Arismonx/nriny-shop/modules/player/playerUsecase"
)

type (
	PlayerQueueHandler interface{}

	playerQueueHandler struct {
		cfg           *config.Config
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerQueueHandler(
	cfg *config.Config,
	playerUsecase playerUsecase.PlayerUsecaseService,
) PlayerQueueHandler {
	return &playerQueueHandler{cfg, playerUsecase}
}
