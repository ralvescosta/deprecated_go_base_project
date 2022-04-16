package usecases

import (
	"context"

	"github.com/ralvescosta/base/pkg/app/interfaces"
	"github.com/ralvescosta/base/pkg/domain/usecases"
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
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
