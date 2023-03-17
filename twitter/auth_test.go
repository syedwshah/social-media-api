package twitter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRegisterInput_Sanitize(t *testing.T) {
	input := RegisterInput{
		Username:        " Bob ",
		Email:           " BOB@gamil.com ",
		Password:        "password",
		ConfirmPassword: "password",
	}

	want := RegisterInput{
		Username:        "Bob",
		Email:           "bob@gamil.com",
		Password:        "password",
		ConfirmPassword: "password",
	}

	input.Sanitize()

	require.Equal(t, want, input)
}

func TestRegisterInput_Validate(t *testing.T) {
	testCases := []struct {
		name  string
		input RegisterInput
		err   error
	}{
		{
			name: "valid",
			input: RegisterInput{
				Username:        "Bob",
				Email:           "bob@gamil.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: nil,
		},
		{
			name: "invalid email",
			input: RegisterInput{
				Username:        "Bob",
				Email:           "bob",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		},
		{
			name: "username too short",
			input: RegisterInput{
				Username:        "B",
				Email:           "bob@gamil.com",
				Password:        "password",
				ConfirmPassword: "password",
			},
			err: ErrValidation,
		},
		{
			name: "password too short",
			input: RegisterInput{
				Username:        "Bob",
				Email:           "bob@gamil.com",
				Password:        "pass",
				ConfirmPassword: "pass",
			},
			err: ErrValidation,
		},
		{
			name: "password does not match confirm password",
			input: RegisterInput{
				Username:        "Bob",
				Email:           "bob@gamil.com",
				Password:        "password",
				ConfirmPassword: "wrongpassword",
			},
			err: ErrValidation,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()

			if tc.err != nil {
				//error
				require.ErrorIs(t, err, tc.err)
			} else {
				//no error
				require.NoError(t, err)
			}
		})
	}
}

func TestLoginInput_Sanitize(t *testing.T) {
	input := LoginInput{
		Email:    " BOB@gamil.com ",
		Password: "password",
	}

	want := LoginInput{
		Email:    "bob@gamil.com",
		Password: "password",
	}

	input.Sanitize()

	require.Equal(t, want, input)
}

func TestLoginInput_Validate(t *testing.T) {
	testCases := []struct {
		name  string
		input LoginInput
		err   error
	}{
		{
			name: "valid",
			input: LoginInput{
				Email:    "bob@gamil.com",
				Password: "password",
			},
			err: nil,
		},
		{
			name: "invalid email",
			input: LoginInput{
				Email:    "bob",
				Password: "password",
			},
			err: ErrValidation,
		},
		{
			name: "empty password",
			input: LoginInput{
				Email:    "bob@gamil.com",
				Password: "",
			},
			err: ErrValidation,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := tc.input.Validate()

			if tc.err != nil {
				//error
				require.ErrorIs(t, err, tc.err)
			} else {
				//no error
				require.NoError(t, err)
			}
		})
	}
}
