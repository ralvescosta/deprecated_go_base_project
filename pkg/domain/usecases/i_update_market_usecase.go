package usecases

import (
	"context"

	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
)

type IUpdateMarketUseCase interface {
	Execute(ctx context.Context, registerCode string, market valueObjects.MarketValueObjects) (valueObjects.MarketValueObjects, error)
}
