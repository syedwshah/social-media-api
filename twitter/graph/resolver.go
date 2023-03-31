package graph

//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
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
