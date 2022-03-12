package usecases

import "context"

type IDeleteMarketUseCase interface {
	Execute(ctx context.Context, registerCode string) error
}
