package dao_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	pgctx "github.com/a-novel-kit/context/pgbun"

	"github.com/a-novel/authentication/internal/dao"
)

func TestExistsCredentialsEmail(t *testing.T) {
	testCases := []struct {
		name string

		fixtures []*dao.CredentialsEntity

		email string

		expect    bool
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

			email: "user@provider.com",

			expect: true,
		},
		{
			name: "NotFound",

			email: "user@provider.com",

			expect: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			selectKey := dao.NewExistsCredentialsEmailRepository()

			tx, commit, err := pgctx.NewContextTX(ctx, nil)
			require.NoError(t, err)

			defer func() { _ = commit(false) }()

			db, err := pgctx.Context(tx)
			require.NoError(t, err)

			for _, fixture := range testCase.fixtures {
				_, err = db.NewInsert().Model(fixture).Exec(tx)
				require.NoError(t, err)
			}

			key, err := selectKey.ExistsCredentialsEmail(tx, testCase.email)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, key)
		})
	}
}
