package usecases

import (
	"context"
	"markets/pkg/domain/usecases"
	valueObjects "markets/pkg/domain/value_objects"
	"markets/pkg/infra/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetMarketByQuery_Execute(t *testing.T) {
	t.Run("should execute correctly", func(t *testing.T) {
		sut := makeGetMarketByQuerySut()

		ctx := context.Background()

		sut.repo.On("Find", ctx, sut.marketMocked).Return([]valueObjects.MarketValueObjects{{}}, nil)

		result, err := sut.useCase.Execute(ctx, sut.marketMocked)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})
}

type getMarketByQuerySutRtn struct {
	repo         *repositories.MarketRepositorySpy
	useCase      usecases.IGetMarketByQueryUseCase
	marketMocked valueObjects.MarketValueObjects
}

func makeGetMarketByQuerySut() getMarketByQuerySutRtn {
	repo := repositories.NewMarketRepositorySpy()

	useCase := NewGetMarketByQueryUseCase(repo)

	marketMocked := valueObjects.MarketValueObjects{}
	return getMarketByQuerySutRtn{repo, useCase, marketMocked}
}
