package usecases

import (
	"context"

	valueObjects "markets/pkg/domain/value_objects"
)

type IGetMarketByQueryUseCase interface {
	Execute(ctx context.Context, market valueObjects.MarketValueObjects) ([]valueObjects.MarketValueObjects, error)
}
