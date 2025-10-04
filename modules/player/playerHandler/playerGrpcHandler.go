package playerHandler

import (
	"github.com/Arismonx/nriny-shop/modules/player/playerUsecase"
)

type (
	payerGrpcHandler struct {
		playerUsecase playerUsecase.PlayerUsecaseService
	}
)

func NewPlayerGrpcHandler(
	playerUsecase playerUsecase.PlayerUsecaseService,
) *payerGrpcHandler {
	return &payerGrpcHandler{playerUsecase}
}
