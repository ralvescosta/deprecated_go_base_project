package repositories

import (
	"context"
	valueObjects "markets/pkg/domain/value_objects"

	"github.com/stretchr/testify/mock"
)

type MarketRepositorySpy struct {
	mock.Mock
}

func (pst MarketRepositorySpy) Create(ctx context.Context, market valueObjects.MarketValueObjects) (valueObjects.MarketValueObjects, error) {
	args := pst.Called(ctx, market)

	return args.Get(0).(valueObjects.MarketValueObjects), args.Error(1)
}

func (pst MarketRepositorySpy) Find(ctx context.Context, market valueObjects.MarketValueObjects) ([]valueObjects.MarketValueObjects, error) {
	args := pst.Called(ctx, market)

	return args.Get(0).([]valueObjects.MarketValueObjects), args.Error(1)
}

func NewMarketRepositorySpy() *MarketRepositorySpy {
	return new(MarketRepositorySpy)
}
