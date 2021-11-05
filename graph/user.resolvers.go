package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"golang-ent-gqlgen-example/ent"
	"golang-ent-gqlgen-example/ent/user"
	"golang-ent-gqlgen-example/graph/generated"
	"time"
)

func (r *queryResolver) User(ctx context.Context) (*ent.User, error) {
	return r.client.User.Query().Where(user.IDEQ(1)).Only(ctx)
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *ent.User) (string, error) {
	return obj.CreatedAt.Format(time.RFC3339), nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
