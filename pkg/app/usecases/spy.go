package usecases

import (
	"context"

	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"

	"github.com/stretchr/testify/mock"
)

//
type CreateMarketUseCaseSpy struct {
	mock.Mock
}

func (pst CreateMarketUseCaseSpy) Execute(ctx context.Context, market valueObjects.MarketValueObjects) (valueObjects.MarketValueObjects, bool, error) {
	args := pst.Called(ctx, market)

	return args.Get(0).(valueObjects.MarketValueObjects), args.Bool(1), args.Error(2)
}
func NewCreateMarketUseCaseSpy() *CreateMarketUseCaseSpy {
	return new(CreateMarketUseCaseSpy)
}

//
type DeleteMarketUseCaseSpy struct {
	mock.Mock
}

func (pst DeleteMarketUseCaseSpy) Execute(ctx context.Context, registerCode string) error {
	args := pst.Called(ctx, registerCode)

	return args.Error(0)
}

func NewDeleteMarketUseCaseSpy() *DeleteMarketUseCaseSpy {
	return new(DeleteMarketUseCaseSpy)
}

//
type GetMarketByQueryUseCaseSpy struct {
	mock.Mock
}

func (pst GetMarketByQueryUseCaseSpy) Execute(ctx context.Context, market valueObjects.MarketValueObjects) ([]valueObjects.MarketValueObjects, error) {
	args := pst.Called(ctx, market)

	return args.Get(0).([]valueObjects.MarketValueObjects), args.Error(1)
}

func NewGetMarketByQueryUseCaseSpy() *GetMarketByQueryUseCaseSpy {
	return new(GetMarketByQueryUseCaseSpy)
}

//
type UpdateMarketUseCaseSpy struct {
	mock.Mock
}

func (pst UpdateMarketUseCaseSpy) Execute(ctx context.Context, registerCode string, market valueObjects.MarketValueObjects) (valueObjects.MarketValueObjects, error) {
	args := pst.Called(ctx, registerCode, market)

	return args.Get(0).(valueObjects.MarketValueObjects), args.Error(1)
}

func NewUpdateMarketUseCaseSpy() *UpdateMarketUseCaseSpy {
	return new(UpdateMarketUseCaseSpy)
}
