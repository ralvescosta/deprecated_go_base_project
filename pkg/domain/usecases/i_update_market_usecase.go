package usecases

import (
	"context"

	valueObjects "markets/pkg/domain/value_objects"
)

type IUpdateMarketUseCase interface {
	Execute(ctx context.Context, registerCode string, market valueObjects.MarketValueObjects) (valueObjects.MarketValueObjects, error)
}
