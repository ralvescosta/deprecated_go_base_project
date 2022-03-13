package interfaces

import (
	"context"

	valueObjects "markets/pkg/domain/value_objects"
)

type IMarketRepository interface {
	Create(ctx context.Context, market valueObjects.MarketValueObjects) (valueObjects.MarketValueObjects, error)
	Find(ctx context.Context, market valueObjects.MarketValueObjects) ([]valueObjects.MarketValueObjects, error)
	Delete(ctx context.Context, registerCode string) error
	Update(ctx context.Context, registerCode string, market valueObjects.MarketValueObjects) (valueObjects.MarketValueObjects, error)
}
