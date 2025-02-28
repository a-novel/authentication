package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/a-novel-kit/configurator/utilstest"

	"github.com/a-novel/authentication/api/codegen"
)

// STORY: The user can create an account, and login with it.

// Email log has the field "message" which contains "send transactional email".
var emailLogCaptureRE = regexp.MustCompile(`"message":"send transactional email"`)

func captureEmailLog(t *testing.T, email string) utilstest.LogCaptureFN {
	t.Helper()

	return func(log string) bool {
		if !emailLogCaptureRE.MatchString(log) {
			return false
		}

		var content struct {
			To []string `json:"to"`
		}

		if err := json.Unmarshal([]byte(log), &content); err != nil {
			t.Errorf("Failed to unmarshal email log (%q): %v", log, err)

			return false
		}

		for _, to := range content.To {
			if to == email {
				return true
			}
		}

		return false
	}
}

func extractShortCode(log string) (string, error) {
	var out struct {
		DynamicTemplateData struct {
			ShortCode string `json:"shortCode"`
		} `json:"dynamicTemplateData"`
	}

	if err := json.Unmarshal([]byte(log), &out); err != nil {
		return "", fmt.Errorf("unmarshal email log: %w", err)
	}

	return out.DynamicTemplateData.ShortCode, nil
}

type userData struct {
	email    string
	password string
	user     string
	token    string
}

func createUser(t *testing.T, client *codegen.Client) *userData {
	t.Helper()

	user := getRandomString()
	email := user + "@example.com"
	password := getRandomString()

	var shortCode, token string

	t.Run("RequestRegistration", func(t *testing.T) {
		capturer := utilstest.WaitForLog(logs, captureEmailLog(t, email), 10*time.Second)

		_, err := client.RequestRegistration(t.Context(), &codegen.RequestRegistrationForm{
			Email: codegen.Email(email),
		})
		require.NoError(t, err)

		log, err := capturer()
		require.NoError(t, err)

		// Check if the email contains the invitation code.
		shortCode, err = extractShortCode(log)
		require.NoError(t, err)
		require.NotEmpty(t, shortCode)
	})

	t.Run("CreateUser", func(t *testing.T) {
		rawRes, err := client.Register(t.Context(), &codegen.RegisterForm{
			Email:     codegen.Email(email),
			Password:  codegen.Password(password),
			ShortCode: codegen.ShortCode(shortCode),
		})
		require.NoError(t, err)

		res, ok := rawRes.(*codegen.Token)
		require.True(t, ok)
		require.NotEmpty(t, res.GetAccessToken())

		token = res.GetAccessToken()
	})

	return &userData{
		email:    email,
		password: password,
		user:     user,
		token:    token,
	}
}

func TestUserLifecycle(t *testing.T) {
	t.Parallel()

	client, securityClient, err := getServerClient()
	require.NoError(t, err)

	token := authAnon(t, client)
	securityClient.SetToken(token)

	user := createUser(t, client)
	securityClient.SetToken(user.token)

	claims := checkSession(t, client)
	userID := claims.GetUserID()

	t.Run("Login/WrongPassword", func(t *testing.T) {
		res, err := client.CreateSession(t.Context(), &codegen.LoginForm{
			Email:    codegen.Email(user.email),
			Password: "fakepassword",
		})
		require.NoError(t, err)

		require.IsType(t, &codegen.ForbiddenError{}, res)
	})

	t.Run("Login", func(t *testing.T) {
		res, err := client.CreateSession(t.Context(), &codegen.LoginForm{
			Email:    codegen.Email(user.email),
			Password: codegen.Password(user.password),
		})
		require.NoError(t, err)

		token, ok := res.(*codegen.Token)
		require.True(t, ok)

		require.NotEqual(t, token.GetAccessToken(), user.token)
		user.token = token.GetAccessToken()
	})

	securityClient.SetToken(user.token)

	claims = checkSession(t, client)
	require.Equal(t, userID, claims.GetUserID())
}
