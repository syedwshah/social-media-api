//go:build integration
// +build integration

package domain

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/syedwshah/twitter"
	"github.com/syedwshah/twitter/test_helpers"
)

func TestIntegrationAuthServices_Register(t *testing.T) {
	validInput := twitter.RegisterInput{
		Username:        "bob",
		Email:           "bob@example.com",
		Password:        "password",
		ConfirmPassword: "password",
	}

	t.Run("can register a user", func(t *testing.T) {
		ctx := context.Background()

		defer test_helpers.TeardownDB(ctx, t, db)

		res, err := authService.Register(ctx, validInput)
		require.NoError(t, err)

		require.NotEmpty(t, res.User.ID)
		require.Equal(t, validInput.Email, res.User.Email)
		require.Equal(t, validInput.Username, res.User.Username)
		require.NotEqual(t, validInput.Password, res.User.Password) //pass should not be plain text anymore
	})

	t.Run("existing username", func(t *testing.T) {
		ctx := context.Background()

		defer test_helpers.TeardownDB(ctx, t, db)

		_, err := authService.Register(ctx, validInput)
		require.NoError(t, err)

		_, err = authService.Register(ctx, twitter.RegisterInput{
			Username:        validInput.Username,
			Email:           "bob2@example.com",
			Password:        "password",
			ConfirmPassword: "password",
		})
		require.ErrorIs(t, err, twitter.ErrUsernameTaken)

	})

	t.Run("existing email", func(t *testing.T) {
		ctx := context.Background()

		defer test_helpers.TeardownDB(ctx, t, db)

		_, err := authService.Register(ctx, validInput)
		require.NoError(t, err)

		_, err = authService.Register(ctx, twitter.RegisterInput{
			Username:        "bob2",
			Email:           validInput.Email,
			Password:        "password",
			ConfirmPassword: "password",
		})
		require.ErrorIs(t, err, twitter.ErrEmailTaken)

	})
}
