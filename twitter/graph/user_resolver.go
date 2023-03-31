package graph

import (
	"context"

	"github.com/syedwshah/twitter"
)

func mapUser(u twitter.User) *User {
	return &User{
		ID:        u.ID,
		Email:     u.Email,
		Username:  u.Username,
		CreatedAt: u.CreatedAt,
	}
}

func (q *queryResolver) Me(ctx context.Context) (*User, error) {
	panic("implement me")
}