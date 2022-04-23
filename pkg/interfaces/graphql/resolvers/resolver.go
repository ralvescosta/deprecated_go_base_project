package resolvers

import (
	"github.com/ralvescosta/base/pkg/domain/usecases"
	"github.com/ralvescosta/base/pkg/interfaces/graphql/graph/generated"
	"github.com/ralvescosta/base/pkg/interfaces/graphql/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	marketCreatedNotifier   chan *model.Market
	createMarkerUseCase     usecases.ICreateMarketUseCase
	getMarketByQueryUseCase usecases.IGetMarketByQueryUseCase
	updateMarketUseCase     usecases.IUpdateMarketUseCase
	deleteMarketUseCase     usecases.IDeleteMarketUseCase
}

func NewResolver(
	createMarkerUseCase usecases.ICreateMarketUseCase,
	getMarketByQueryUseCase usecases.IGetMarketByQueryUseCase,
	updateMarketUseCase usecases.IUpdateMarketUseCase,
	deleteMarketUseCase usecases.IDeleteMarketUseCase,
) generated.ResolverRoot {

	marketCreatedNotifier := make(chan *model.Market)

	return &Resolver{
		marketCreatedNotifier,
		createMarkerUseCase,
		getMarketByQueryUseCase,
		updateMarketUseCase,
		deleteMarketUseCase,
	}
}
