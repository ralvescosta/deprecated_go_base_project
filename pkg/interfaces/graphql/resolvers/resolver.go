package resolvers

import (
	"github.com/ralvescosta/base/pkg/domain/usecases"
	"github.com/ralvescosta/base/pkg/interfaces/graphql/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
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
	return &Resolver{
		createMarkerUseCase,
		getMarketByQueryUseCase,
		updateMarketUseCase,
		deleteMarketUseCase,
	}
}
