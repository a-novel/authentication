package lib_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/a-novel/authentication/internal/lib"
)

func TestScrypt(t *testing.T) {
	t.Parallel()

	password := "password"

	encrypted, err := lib.GenerateScrypt(password, lib.ScryptParamsDefault)
	require.NoError(t, err)
	require.NotEmpty(t, encrypted)

	testCases := []struct {
		name string

		password  string
		encrypted string

		expectErr error
	}{
		{
			name: "OK",

			password:  password,
			encrypted: encrypted,
		},
		{
			name: "WrongPassword",

			password:  "wrongpassword",
			encrypted: encrypted,

			expectErr: lib.ErrInvalidPassword,
		},
		{
			name: "Malformed/NotEnoughParts",

			password:  "password",
			encrypted: "malformed$",

			expectErr: lib.ErrInvalidHash,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			err = lib.CompareScrypt(testCase.password, testCase.encrypted)
			require.ErrorIs(t, err, testCase.expectErr)
		})
	}
}
