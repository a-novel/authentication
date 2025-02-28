package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/a-novel/authentication/api/codegen"
)

func TestAnonAuthAPI(t *testing.T) { //nolint:tparallel
	t.Parallel()

	client, securityClient, err := getServerClient()
	require.NoError(t, err)

	securityClient.SetToken("")

	t.Run("CheckSession/Unauthenticated", func(t *testing.T) { //nolint:paralleltest
		_, err = client.CheckSession(t.Context())

		var ogenError *codegen.UnexpectedErrorStatusCode

		require.ErrorAs(t, err, &ogenError)
		require.Equal(t, http.StatusUnauthorized, ogenError.StatusCode)
	})

	t.Run("Authenticate/Anon", func(t *testing.T) { //nolint:paralleltest
		rawRes, err := client.CreateAnonSession(t.Context())
		require.NoError(t, err)

		res, ok := rawRes.(*codegen.Token)
		require.True(t, ok)
		require.NotEmpty(t, res.GetAccessToken())

		securityClient.SetToken(res.GetAccessToken())
	})

	t.Run("CheckSession/Authenticated", func(t *testing.T) { //nolint:paralleltest
		_, err = client.CheckSession(t.Context())
		require.NoError(t, err)
	})
}
