package usecases

import (
	"context"
	valueObjects "markets/pkg/domain/value_objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateKeySpy_Execute(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewCreateMarketUseCaseSpy()

		ctx := context.Background()
		market := valueObjects.MarketValueObjects{}

		sut.On("Execute", ctx, market).Return(market, false, nil)

		result, alreadyCreated, err := sut.Execute(ctx, market)

		assert.NoError(t, err)
		assert.False(t, alreadyCreated)
		assert.Equal(t, market, result)
		sut.AssertExpectations(t)
	})
}

func Test_DeleteMarketSpy_Execute(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewDeleteMarketUseCaseSpy()

		ctx := context.Background()

		sut.On("Execute", ctx, "registro").Return(nil)

		err := sut.Execute(ctx, "registro")

		assert.NoError(t, err)
		sut.AssertExpectations(t)
	})
}

func Test_GetMarketByQuerySpy_Execute(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewGetMarketByQueryUseCaseSpy()

		ctx := context.Background()
		market := valueObjects.MarketValueObjects{}

		sut.On("Execute", ctx, market).Return([]valueObjects.MarketValueObjects{{}}, nil)

		result, err := sut.Execute(ctx, market)

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		sut.AssertExpectations(t)
	})
}

func Test_UpdateMarketSpy_Execute(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := NewUpdateMarketUseCaseSpy()

		ctx := context.Background()
		market := valueObjects.MarketValueObjects{}

		sut.On("Execute", ctx, "registro", market).Return(market, nil)

		result, err := sut.Execute(ctx, "registro", market)

		assert.NoError(t, err)
		assert.Equal(t, market, result)
		sut.AssertExpectations(t)
	})
}
