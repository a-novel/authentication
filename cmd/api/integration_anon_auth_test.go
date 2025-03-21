package main

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/a-novel/authentication/api/apiclient/testapiclient"
	"github.com/a-novel/authentication/api/codegen"
)

// STORY: The user can create an anonymous session, and this session is valid.

func authAnon(t *testing.T, client *codegen.Client) string {
	t.Helper()

	var token string

	t.Log("Authenticate/Anon")
	{
		rawRes, err := client.CreateAnonSession(t.Context())
		require.NoError(t, err)

		res, ok := rawRes.(*codegen.Token)
		require.True(t, ok)
		require.NotEmpty(t, res.GetAccessToken())

		token = res.GetAccessToken()
	}

	return token
}

func checkSession(t *testing.T, client *codegen.Client) *codegen.Claims {
	rawRes, err := client.CheckSession(t.Context())
	require.NoError(t, err)

	res, ok := rawRes.(*codegen.Claims)
	require.True(t, ok)

	return res
}

func TestAnonAuthAPI(t *testing.T) {
	client, securityClient, err := testapiclient.GetServerClient()
	require.NoError(t, err)

	t.Log("CheckSession/Unauthenticated")
	{
		_, err = client.CheckSession(t.Context())

		var ogenError *codegen.UnexpectedErrorStatusCode

		require.ErrorAs(t, err, &ogenError)
		require.Equal(t, http.StatusUnauthorized, ogenError.StatusCode)
	}

	token := authAnon(t, client)
	securityClient.SetToken(token)

	_ = checkSession(t, client)
}
