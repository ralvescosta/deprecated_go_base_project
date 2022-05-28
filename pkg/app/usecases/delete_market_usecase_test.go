package usecases

import (
	"context"

	"testing"

	"github.com/ralvescosta/base/pkg/app/errors"
	"github.com/ralvescosta/base/pkg/domain/usecases"
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
	"github.com/ralvescosta/base/pkg/infra/repositories"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DeleteMarketUseCaseTestSuite struct {
	suite.Suite
}

func TestDeleteMarketUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(DeleteMarketUseCaseTestSuite))
}

func TestFunc() {}

func Test_DeleteMarket_Execute(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeDeleteMarketSut()

		ctx := context.Background()

		sut.repo.On("Find", ctx, valueObjects.MarketValueObjects{Registro: "registro"}).Return([]valueObjects.MarketValueObjects{{}}, nil)
		sut.repo.On("Delete", ctx, "registro").Return(nil)

		err := sut.useCase.Execute(ctx, "registro")

		assert.NoError(t, err)
		sut.repo.AssertExpectations(t)
	})

	t.Run("should return notFoundError if the market was not found", func(t *testing.T) {
		sut := makeDeleteMarketSut()

		ctx := context.Background()

		sut.repo.On("Find", ctx, valueObjects.MarketValueObjects{Registro: "registro"}).Return([]valueObjects.MarketValueObjects(nil), nil)

		err := sut.useCase.Execute(ctx, "registro")

		assert.Error(t, err)
		assert.IsType(t, errors.NotFoundError{}, err)
		sut.repo.AssertExpectations(t)
	})

	t.Run("should return error if some error occur during the find", func(t *testing.T) {
		sut := makeDeleteMarketSut()

		ctx := context.Background()

		sut.repo.On("Find", ctx, valueObjects.MarketValueObjects{Registro: "registro"}).Return([]valueObjects.MarketValueObjects(nil), errors.NewInternalError("some error"))

		err := sut.useCase.Execute(ctx, "registro")

		assert.Error(t, err)
		assert.IsType(t, errors.InternalError{}, err)
		sut.repo.AssertExpectations(t)
	})
}

type deleteMarketSutRtn struct {
	repo    *repositories.MarketRepositorySpy
	useCase usecases.IDeleteMarketUseCase
}

func makeDeleteMarketSut() deleteMarketSutRtn {
	repo := repositories.NewMarketRepositorySpy()

	useCase := NewDeleteMarketUseCase(repo)
	return deleteMarketSutRtn{repo, useCase}
}
