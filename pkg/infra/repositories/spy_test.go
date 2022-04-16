package repositories

import (
	"context"
	"testing"

	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
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

func Test_Delete(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewMarketRepositorySpy()

		ctx := context.Background()
		sut.On("Delete", ctx, "register").Return(nil)

		sut.Delete(ctx, "register")

		sut.AssertExpectations(t)
	})
}

func Test_Update(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewMarketRepositorySpy()

		ctx := context.Background()
		market := valueObjects.MarketValueObjects{}
		sut.On("Update", ctx, "register", market).Return(market, nil)

		sut.Update(ctx, "register", market)

		sut.AssertExpectations(t)
	})
}
