package usecases

import (
	"context"
	"testing"

	"github.com/ralvescosta/base/pkg/app/errors"
	"github.com/ralvescosta/base/pkg/domain/usecases"
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
	"github.com/ralvescosta/base/pkg/infra/repositories"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateMarket_Execute(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeUpdateMarketSutRtn()

		ctx := context.Background()
		sut.repo.On("Find", ctx, valueObjects.MarketValueObjects{Registro: "registro"}).Return([]valueObjects.MarketValueObjects{{}}, nil)
		sut.repo.On("Update", ctx, "registro", sut.marketMocked).Return(sut.marketMocked, nil)

		result, err := sut.useCase.Execute(ctx, "registro", sut.marketMocked)

		assert.NoError(t, err)
		assert.Equal(t, sut.marketMocked, result)
	})

	t.Run("should return erro if some error occur during the update", func(t *testing.T) {
		sut := makeUpdateMarketSutRtn()

		ctx := context.Background()

		sut.repo.On("Find", ctx, valueObjects.MarketValueObjects{Registro: "registro"}).Return([]valueObjects.MarketValueObjects{{}}, nil)
		sut.repo.On("Update", ctx, "registro", sut.marketMocked).Return(valueObjects.MarketValueObjects{}, errors.NewInternalError("some error"))

		_, err := sut.useCase.Execute(ctx, "registro", sut.marketMocked)

		assert.Error(t, err)
		sut.repo.AssertExpectations(t)
	})

	t.Run("should return error if some error occur during the find", func(t *testing.T) {
		sut := makeUpdateMarketSutRtn()

		ctx := context.Background()

		sut.repo.On("Find", ctx, valueObjects.MarketValueObjects{Registro: "registro"}).Return([]valueObjects.MarketValueObjects(nil), errors.NewInternalError("some error"))

		_, err := sut.useCase.Execute(ctx, "registro", sut.marketMocked)

		assert.Error(t, err)
		sut.repo.AssertExpectations(t)
	})

	t.Run("should return notFoundError if the market was not found", func(t *testing.T) {
		sut := makeUpdateMarketSutRtn()

		ctx := context.Background()

		sut.repo.On("Find", ctx, valueObjects.MarketValueObjects{Registro: "registro"}).Return([]valueObjects.MarketValueObjects(nil), nil)

		_, err := sut.useCase.Execute(ctx, "registro", sut.marketMocked)

		assert.Error(t, err)
		assert.IsType(t, errors.NotFoundError{}, err)
		sut.repo.AssertExpectations(t)
	})
}

type updateMarketSutRtn struct {
	repo         *repositories.MarketRepositorySpy
	useCase      usecases.IUpdateMarketUseCase
	marketMocked valueObjects.MarketValueObjects
}

func makeUpdateMarketSutRtn() updateMarketSutRtn {
	repo := repositories.NewMarketRepositorySpy()
	useCase := NewUpdateMarketUseCase(repo)

	marketMocked := valueObjects.MarketValueObjects{}
	return updateMarketSutRtn{repo, useCase, marketMocked}
}
