package main

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	pgctx "github.com/a-novel-kit/context/pgbun"

	"github.com/a-novel/authentication/api/codegen"
	"github.com/a-novel/authentication/internal/dao"
	"github.com/a-novel/authentication/models"
)

func TestListUsers(t *testing.T) {
	t.Parallel()

	client, securityClient, err := getServerClient()
	require.NoError(t, err)

	// YOLO
	ctx, err := pgctx.NewContext(t.Context(), nil)
	require.NoError(t, err)

	fixtures := []*dao.CredentialsEntity{
		{
			ID:        uuid.New(),
			Email:     getRandomString() + "@email.com",
			Role:      models.CredentialsRoleUser,
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        uuid.New(),
			Email:     getRandomString() + "@email.com",
			Role:      models.CredentialsRoleAdmin,
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:        uuid.New(),
			Email:     getRandomString() + "@email.com",
			Role:      models.CredentialsRoleUser,
			CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
			UpdatedAt: time.Date(2021, 1, 3, 0, 0, 0, 0, time.UTC),
		},
	}

	tx, err := pgctx.Context(ctx)
	require.NoError(t, err)

	_, err = tx.NewInsert().Model(&fixtures).Exec(ctx)
	require.NoError(t, err)

	token := authAnon(t, client)
	securityClient.SetToken(token)

	// Verify the fixtures are present in the list of returned users.
	t.Log("ListUsers")
	{
		rawRes, err := client.ListUsers(t.Context(), codegen.ListUsersParams{})
		require.NoError(t, err)

		res, ok := rawRes.(*codegen.ListUsersOKApplicationJSON)
		require.True(t, ok)

		targetedUsers := lo.Filter(*res, func(item codegen.User, _ int) bool {
			return lo.Contains(
				lo.Map(fixtures, func(item *dao.CredentialsEntity, _ int) codegen.UserID {
					return codegen.UserID(item.ID)
				}),
				item.GetID(),
			)
		})

		require.Len(t, targetedUsers, len(fixtures))

		for _, user := range targetedUsers {
			require.Contains(t, fixtures, &dao.CredentialsEntity{
				ID:        uuid.UUID(user.GetID()),
				Email:     string(user.GetEmail()),
				Role:      models.CredentialsRole(user.GetRole()),
				CreatedAt: user.GetCreatedAt(),
				UpdatedAt: user.GetUpdatedAt(),
			})
		}
	}

	t.Log("FilterByRole")
	{
		fixtures = []*dao.CredentialsEntity{fixtures[0], fixtures[2]}

		rawRes, err := client.ListUsers(t.Context(), codegen.ListUsersParams{
			Roles: []codegen.CredentialsRole{codegen.CredentialsRoleUser},
		})
		require.NoError(t, err)

		res, ok := rawRes.(*codegen.ListUsersOKApplicationJSON)
		require.True(t, ok)

		targetedUsers := lo.Filter(*res, func(item codegen.User, _ int) bool {
			return lo.Contains(
				lo.Map(fixtures, func(item *dao.CredentialsEntity, _ int) codegen.UserID {
					return codegen.UserID(item.ID)
				}),
				item.GetID(),
			)
		})

		require.Len(t, targetedUsers, len(fixtures))

		for _, user := range targetedUsers {
			require.Contains(t, fixtures, &dao.CredentialsEntity{
				ID:        uuid.UUID(user.GetID()),
				Email:     string(user.GetEmail()),
				Role:      models.CredentialsRole(user.GetRole()),
				CreatedAt: user.GetCreatedAt(),
				UpdatedAt: user.GetUpdatedAt(),
			})
		}
	}
}
