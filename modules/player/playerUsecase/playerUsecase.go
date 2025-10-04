package playerUsecase

import "github.com/Arismonx/nriny-shop/modules/player/playerRepository"

type (
	PlayerUsecaseService interface{}

	playerUsecase struct {
		playerRepository playerRepository.PlayerRepositoryService
	}
)

func NewPlayerUsecase(
	playerRepository playerRepository.PlayerRepositoryService,
) PlayerUsecaseService {
	return &playerUsecase{playerRepository}
}
