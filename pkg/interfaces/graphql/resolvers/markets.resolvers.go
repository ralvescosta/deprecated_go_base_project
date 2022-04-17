package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ralvescosta/base/pkg/interfaces/graphql/graph/generated"
	"github.com/ralvescosta/base/pkg/interfaces/graphql/graph/model"
)

func (r *mutationResolver) CreateMarket(ctx context.Context, create model.CreateMarket) (*model.Market, error) {
	return &model.Market{}, nil
}

func (r *mutationResolver) UpdateMarket(ctx context.Context, update model.MarketToUpdate) (*model.Market, error) {
	return &model.Market{}, nil
}

func (r *mutationResolver) DeleteMarket(ctx context.Context, registerCode string) (bool, error) {
	return true, nil
}

func (r *queryResolver) GetMarkets(ctx context.Context, query model.MarketFilters) ([]*model.Market, error) {
	result, err := r.getMarketByQueryUseCase.Execute(ctx, model.MarketFilterToValueObject(query))
	if err != nil {
		return nil, err
	}

	return model.ValueObjectSliceToMarketSlice(result), nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
