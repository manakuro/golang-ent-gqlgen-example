package graph

import (
	"golang-ent-gqlgen-example/ent"
	"golang-ent-gqlgen-example/graph/generated"

	"github.com/99designs/gqlgen/graphql"
)

type Resolver struct{ client *ent.Client }

func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return generated.NewExecutableSchema(generated.Config{
		Resolvers: &Resolver{client},
	})
}
