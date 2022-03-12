package usecases

import (
	"context"
	"fmt"

	"markets/pkg/app/errors"
	"markets/pkg/app/interfaces"
	"markets/pkg/domain/usecases"
	valueObjects "markets/pkg/domain/value_objects"
)

type deleteMarketUseCase struct {
	repo interfaces.IMarketRepository
}

func (pst deleteMarketUseCase) Execute(ctx context.Context, registerCode string) error {
	result, err := pst.repo.Find(ctx, valueObjects.MarketValueObjects{Registro: registerCode})
	if err != nil {
		return err
	}

	if len(result) == 0 {
		return errors.NewNotFoundError(fmt.Sprintf("Market with the RegisterCode: %s was not found", registerCode))
	}

	return pst.repo.Delete(ctx, registerCode)
}

func NewDeleteMarketUseCase(repo interfaces.IMarketRepository) usecases.IDeleteMarketUseCase {
	return deleteMarketUseCase{repo}
}
