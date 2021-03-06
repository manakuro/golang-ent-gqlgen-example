// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (a *ArticleQuery) CollectFields(ctx context.Context, satisfies ...string) *ArticleQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		a = a.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return a
}

func (a *ArticleQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *ArticleQuery {
	return a
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) *UserQuery {
	if fc := graphql.GetFieldContext(ctx); fc != nil {
		u = u.collectField(graphql.GetOperationContext(ctx), fc.Field, satisfies...)
	}
	return u
}

func (u *UserQuery) collectField(ctx *graphql.OperationContext, field graphql.CollectedField, satisfies ...string) *UserQuery {
	return u
}
