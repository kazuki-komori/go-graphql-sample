package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/kazuki-komori/go-gql-sample/graphql/generated"
)

func (r *queryResolver) Hello(ctx context.Context, text *string) (*string, error) {
	res := *text + "word"
	return &res, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
