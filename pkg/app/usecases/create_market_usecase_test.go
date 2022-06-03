package usecases

import (
	"context"

	"testing"

	"github.com/ralvescosta/base/pkg/app/errors"
	"github.com/ralvescosta/base/pkg/domain/usecases"
	valueObjects "github.com/ralvescosta/base/pkg/domain/value_objects"
	"github.com/ralvescosta/base/pkg/infra/repositories"

	"github.com/stretchr/testify/suite"
)

type CreateMarketUseCaseTestSuite struct {
	suite.Suite
}

func TestCreateMarketUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(CreateMarketUseCaseTestSuite))
}

func (s *CreateMarketUseCaseTestSuite) TestCreateMarket() {
	sut := makeCreateMarketSut()

	ctx := context.Background()

	sut.repo.On(
		"Find",
		ctx,
		valueObjects.MarketValueObjects{Registro: sut.marketMocked.Registro},
	).Return([]valueObjects.MarketValueObjects(nil), nil)
	sut.repo.On("Create", ctx, sut.marketMocked).Return(sut.marketMocked, nil)

	_, alreadyCreated, err := sut.useCase.Execute(ctx, sut.marketMocked)

	s.NoError(err)
	s.False(alreadyCreated)
	sut.repo.AssertExpectations(s.T())
}

func (s *CreateMarketUseCaseTestSuite) TestCreateMarketInsertErr() {
	sut := makeCreateMarketSut()

	ctx := context.Background()

	sut.repo.On(
		"Find",
		ctx,
		valueObjects.MarketValueObjects{Registro: sut.marketMocked.Registro},
	).Return([]valueObjects.MarketValueObjects(nil), nil)
	sut.repo.On("Create", ctx, sut.marketMocked).Return(valueObjects.MarketValueObjects{}, errors.NewInternalError("some error"))

	_, alreadyCreated, err := sut.useCase.Execute(ctx, sut.marketMocked)

	s.Error(err)
	s.False(alreadyCreated)
	sut.repo.AssertExpectations(s.T())
}

func (s *CreateMarketUseCaseTestSuite) TestCreateMarketConflictErr() {
	sut := makeCreateMarketSut()

	ctx := context.Background()

	sut.repo.On(
		"Find",
		ctx,
		valueObjects.MarketValueObjects{Registro: sut.marketMocked.Registro},
	).Return([]valueObjects.MarketValueObjects{{}}, nil)

	_, alreadyCreated, err := sut.useCase.Execute(ctx, sut.marketMocked)

	s.NoError(err)
	s.True(alreadyCreated)
	sut.repo.AssertExpectations(s.T())
}

func (s *CreateMarketUseCaseTestSuite) TestCreateMarketFindErr() {
	sut := makeCreateMarketSut()

	ctx := context.Background()

	sut.repo.On(
		"Find",
		ctx,
		valueObjects.MarketValueObjects{Registro: sut.marketMocked.Registro},
	).Return([]valueObjects.MarketValueObjects(nil), errors.NewInternalError("some error"))

	_, alreadyCreated, err := sut.useCase.Execute(ctx, sut.marketMocked)

	s.Error(err)
	s.False(alreadyCreated)
	sut.repo.AssertExpectations(s.T())
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
