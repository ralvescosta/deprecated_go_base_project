package usecases

import (
	"context"
	"fmt"

	"github.com/ralvescosta/base/pkg/app/errors"
	"github.com/ralvescosta/base/pkg/app/interfaces"
	"github.com/ralvescosta/base/pkg/domain/usecases"
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
)

type updateMarketUseCase struct {
	repo interfaces.IMarketRepository
}

func (pst updateMarketUseCase) Execute(ctx context.Context, registerCode string, market valueObjects.MarketValueObjects) (valueObjects.MarketValueObjects, error) {
	result, err := pst.repo.Find(ctx, valueObjects.MarketValueObjects{Registro: registerCode})
	if err != nil {
		return valueObjects.MarketValueObjects{}, err
	}

	if len(result) == 0 {
		return valueObjects.MarketValueObjects{}, errors.NewNotFoundError(fmt.Sprintf("Market with the RegisterCode: %s was not found", registerCode))
	}

	updated, err := pst.repo.Update(ctx, registerCode, market)
	if err != nil {
		return valueObjects.MarketValueObjects{}, err
	}

	return updated, nil
}

func NewUpdateMarketUseCase(repo interfaces.IMarketRepository) usecases.IUpdateMarketUseCase {
	return updateMarketUseCase{repo}
}
