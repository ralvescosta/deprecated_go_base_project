package usecases

import (
	"context"
	"markets/pkg/app/interfaces"
	"markets/pkg/domain/usecases"
	valueObjects "markets/pkg/domain/value_objects"
)

type getMarketByQueryUseCase struct {
	repo interfaces.IMarketRepository
}

func (pst getMarketByQueryUseCase) Execute(ctx context.Context, market valueObjects.MarketValueObjects) ([]valueObjects.MarketValueObjects, error) {
	return pst.repo.Find(ctx, market)
}

func NewGetMarketByQueryUseCase(repo interfaces.IMarketRepository) usecases.IGetMarketByQueryUseCase {
	return getMarketByQueryUseCase{repo}
}
