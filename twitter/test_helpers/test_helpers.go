package test_helpers

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/syedwshah/twitter"
	"github.com/syedwshah/twitter/faker"
	"github.com/syedwshah/twitter/postgres"
)

func TeardownDB(ctx context.Context, t *testing.T, db *postgres.DB) {
	t.Helper()

	err := db.Truncate(ctx)
	require.NoError(t, err)

}

func CreateUser(ctx context.Context, t *testing.T, userRepo twitter.UserRepo) twitter.User {
	t.Helper()

	user, err := userRepo.Create(ctx, twitter.User{
		Username: faker.Username(),
		Email:    faker.Email(),
		Password: faker.Password,
	})
	require.NoError(t, err)

	return user
}

func LoginUser(ctx context.Context, t *testing.T, user twitter.User) context.Context {
	t.Helper()

	return twitter.PutUserIDIntoContext(ctx, user.ID)
}
