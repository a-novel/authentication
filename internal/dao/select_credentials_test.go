package dao_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	pgctx "github.com/a-novel-kit/context/pgbun"

	"github.com/a-novel/authentication/internal/dao"
)

func TestSelectCredentials(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		fixtures []*dao.CredentialsEntity

		id uuid.UUID

		expect    *dao.CredentialsEntity
		expectErr error
	}{
		{
			name: "Success",

			fixtures: []*dao.CredentialsEntity{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					Email:     "user@provider.com",
					Password:  "password-2-hashed",
					CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
					UpdatedAt: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
				},
			},

			id: uuid.MustParse("00000000-0000-0000-0000-000000000002"),

			expect: &dao.CredentialsEntity{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
				Email:     "user@provider.com",
				Password:  "password-2-hashed",
				CreatedAt: time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				UpdatedAt: time.Date(2021, 1, 2, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "Error/NotFound",

			id: uuid.MustParse("00000000-0000-0000-0000-000000000002"),

			expectErr: dao.ErrCredentialsNotFound,
		},
	}

	repository := dao.NewSelectCredentialsRepository()

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			tx, commit, err := pgctx.NewContextTX(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
			require.NoError(t, err)

			t.Cleanup(func() {
				_ = commit(false)
			})

			db, err := pgctx.Context(tx)
			require.NoError(t, err)

			if len(testCase.fixtures) > 0 {
				_, err = db.NewInsert().Model(&testCase.fixtures).Exec(tx)
				require.NoError(t, err)
			}

			key, err := repository.SelectCredentials(tx, testCase.id)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, key)
		})
	}
}
