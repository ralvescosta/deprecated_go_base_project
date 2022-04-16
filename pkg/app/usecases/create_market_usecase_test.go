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

func Test_CreateMarket_Execute(t *testing.T) {
	t.Run("should execute correctly when market does not exist yet", func(t *testing.T) {
		sut := makeCreateMarketSut()

		ctx := context.Background()

		sut.repo.On(
			"Find",
			ctx,
			valueObjects.MarketValueObjects{Registro: sut.marketMocked.Registro},
		).Return([]valueObjects.MarketValueObjects(nil), nil)
		sut.repo.On("Create", ctx, sut.marketMocked).Return(sut.marketMocked, nil)

		_, alreadyCreated, err := sut.useCase.Execute(ctx, sut.marketMocked)

		assert.NoError(t, err)
		assert.False(t, alreadyCreated)
		sut.repo.AssertExpectations(t)
	})

	t.Run("should return erro if some error occur during the insert", func(t *testing.T) {
		sut := makeCreateMarketSut()

		ctx := context.Background()

		sut.repo.On(
			"Find",
			ctx,
			valueObjects.MarketValueObjects{Registro: sut.marketMocked.Registro},
		).Return([]valueObjects.MarketValueObjects(nil), nil)
		sut.repo.On("Create", ctx, sut.marketMocked).Return(valueObjects.MarketValueObjects{}, errors.NewInternalError("some error"))

		_, alreadyCreated, err := sut.useCase.Execute(ctx, sut.marketMocked)

		assert.Error(t, err)
		assert.False(t, alreadyCreated)
		sut.repo.AssertExpectations(t)
	})

	t.Run("should execute correctly when market already exist", func(t *testing.T) {
		sut := makeCreateMarketSut()

		ctx := context.Background()

		sut.repo.On(
			"Find",
			ctx,
			valueObjects.MarketValueObjects{Registro: sut.marketMocked.Registro},
		).Return([]valueObjects.MarketValueObjects{{}}, nil)

		_, alreadyCreated, err := sut.useCase.Execute(ctx, sut.marketMocked)

		assert.NoError(t, err)
		assert.True(t, alreadyCreated)
		sut.repo.AssertExpectations(t)
	})

	t.Run("should return error if some error occur during the find", func(t *testing.T) {
		sut := makeCreateMarketSut()

		ctx := context.Background()

		sut.repo.On(
			"Find",
			ctx,
			valueObjects.MarketValueObjects{Registro: sut.marketMocked.Registro},
		).Return([]valueObjects.MarketValueObjects(nil), errors.NewInternalError("some error"))

		_, alreadyCreated, err := sut.useCase.Execute(ctx, sut.marketMocked)

		assert.Error(t, err)
		assert.False(t, alreadyCreated)
		sut.repo.AssertExpectations(t)
	})
}

type createMarketSutRtn struct {
	repo         *repositories.MarketRepositorySpy
	useCase      usecases.ICreateMarketUseCase
	marketMocked valueObjects.MarketValueObjects
}

func makeCreateMarketSut() createMarketSutRtn {
	repo := repositories.NewMarketRepositorySpy()

	useCase := NewCreateMarketUseCase(repo)

	marketMocked := valueObjects.MarketValueObjects{}

	return createMarketSutRtn{repo, useCase, marketMocked}
}
