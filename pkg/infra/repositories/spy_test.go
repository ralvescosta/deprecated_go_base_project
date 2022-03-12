package repositories

import (
	"context"
	valueObjects "markets/pkg/domain/value_objects"
	"testing"
)

func Test_Create(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewMarketRepositorySpy()

		market := valueObjects.MarketValueObjects{}
		ctx := context.Background()
		sut.On("Create", ctx, market).Return(market, nil)

		sut.Create(ctx, market)

		sut.AssertExpectations(t)
	})
}

func Test_Find(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewMarketRepositorySpy()

		market := valueObjects.MarketValueObjects{}
		ctx := context.Background()
		sut.On("Find", ctx, market).Return([]valueObjects.MarketValueObjects{}, nil)

		sut.Find(ctx, market)

		sut.AssertExpectations(t)
	})
}
