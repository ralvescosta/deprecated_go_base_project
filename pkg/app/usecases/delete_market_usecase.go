package usecases

import (
	"context"
	"fmt"

	"github.com/ralvescosta/base/pkg/app/errors"
	"github.com/ralvescosta/base/pkg/app/interfaces"
	"github.com/ralvescosta/base/pkg/domain/usecases"
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
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
