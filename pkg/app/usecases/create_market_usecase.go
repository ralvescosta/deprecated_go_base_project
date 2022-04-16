package usecases

import (
	"context"

	"github.com/ralvescosta/base/pkg/app/interfaces"
	"github.com/ralvescosta/base/pkg/domain/usecases"
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
)

type createMarketUseCase struct {
	repo interfaces.IMarketRepository
}

func (pst createMarketUseCase) Execute(ctx context.Context, market valueObjects.MarketValueObjects) (valueObjects.MarketValueObjects, bool, error) {
	marketCreated, err := pst.repo.Find(ctx, valueObjects.MarketValueObjects{Registro: market.Registro})
	if err != nil {
		return valueObjects.MarketValueObjects{}, false, err
	}

	if marketCreated != nil || len(marketCreated) > 0 {
		return marketCreated[0], true, nil
	}

	result, err := pst.repo.Create(ctx, market)
	if err != nil {
		return valueObjects.MarketValueObjects{}, false, err
	}

	return result, false, nil
}

func NewCreateMarketUseCase(repo interfaces.IMarketRepository) usecases.ICreateMarketUseCase {
	return createMarketUseCase{repo}
}
