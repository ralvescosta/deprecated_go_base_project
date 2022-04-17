package resolvers

import (
	"github.com/ralvescosta/base/pkg/domain/usecases"
	"github.com/ralvescosta/base/pkg/interfaces/graphql/graph/generated"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	getMarketByQueryUseCase usecases.IGetMarketByQueryUseCase
}

func NewResolver(getMarketByQueryUseCase usecases.IGetMarketByQueryUseCase) generated.ResolverRoot {
	return &Resolver{getMarketByQueryUseCase}
}
