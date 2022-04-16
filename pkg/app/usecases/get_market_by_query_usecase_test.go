package usecases

import (
	"context"
	"testing"

	"github.com/ralvescosta/base/pkg/domain/usecases"
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
	"github.com/ralvescosta/base/pkg/infra/repositories"

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
