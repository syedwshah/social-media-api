package graph

import (
	"context"
	"net/http"

	"github.com/99designs/gqlgen/graphql"
	"github.com/syedwshah/twitter"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	AuthService  twitter.AuthService
	TweetService twitter.TweetService
}

type queryResolver struct {
	*Resolver
}

func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct {
	*Resolver
}

func (m *Resolver) Mutation() MutationResolver {
	return &mutationResolver{m}
}

// send error 400 in response
func buildBadRequestError(ctx context.Context, err error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]interface{}{
			"code": http.StatusBadRequest,
		},
	}
}

// send error 401 in response
func buildUnauthenticatedError(ctx context.Context, err error) error {
	return &gqlerror.Error{
		Message: err.Error(),
		Path:    graphql.GetPath(ctx),
		Extensions: map[string]interface{}{
			"code": http.StatusUnauthorized,
		},
	}
}
