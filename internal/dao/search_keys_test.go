package dao_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	pgctx "github.com/a-novel-kit/context/pgbun"

	"github.com/a-novel/authentication/internal/dao"
	"github.com/a-novel/authentication/models"
)

func TestSearchKeys(t *testing.T) {
	hourAgo := time.Now().Add(-time.Hour).UTC().Round(time.Second)
	hourLater := time.Now().Add(time.Hour).UTC().Round(time.Second)

	testCases := []struct {
		name string

		usage    models.KeyUsage
		fixtures []*dao.KeyEntity

		expect    []*dao.KeyEntity
		expectErr error
	}{
		{
			name: "FilterUsage",

			usage: models.KeyUsageAuth,

			fixtures: []*dao.KeyEntity{
				{
					ID:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					PrivateKey: "cHJpdmF0ZS1rZXktMQ",
					PublicKey:  lo.ToPtr("cHVibGljLWtleS0x"),
					Usage:      models.KeyUsageRefresh,
					CreatedAt:  hourAgo,
					ExpiresAt:  hourLater,
				},
				{
					ID:         uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					PrivateKey: "cHJpdmF0ZS1rZXktMg",
					PublicKey:  lo.ToPtr("cHVibGljLWtleS0y"),
					Usage:      models.KeyUsageAuth,
					CreatedAt:  hourAgo,
					ExpiresAt:  hourLater,
				},
			},

			expect: []*dao.KeyEntity{
				{
					ID:         uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					PrivateKey: "cHJpdmF0ZS1rZXktMg",
					PublicKey:  lo.ToPtr("cHVibGljLWtleS0y"),
					Usage:      models.KeyUsageAuth,
					CreatedAt:  hourAgo,
					ExpiresAt:  hourLater,
				},
			},
		},
		{
			name: "OrderCreation",

			usage: models.KeyUsageAuth,

			fixtures: []*dao.KeyEntity{
				{
					ID:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					PrivateKey: "cHJpdmF0ZS1rZXktMQ",
					PublicKey:  lo.ToPtr("cHVibGljLWtleS0x"),
					Usage:      models.KeyUsageAuth,
					CreatedAt:  hourAgo, // Second
					ExpiresAt:  hourLater,
				},
				{
					ID:         uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					PrivateKey: "cHJpdmF0ZS1rZXktMg",
					PublicKey:  lo.ToPtr("cHVibGljLWtleS0y"),
					Usage:      models.KeyUsageAuth,
					CreatedAt:  hourAgo.Add(time.Minute), // First
					ExpiresAt:  hourLater,
				},
				{
					ID:         uuid.MustParse("00000000-0000-0000-0000-000000000003"),
					PrivateKey: "cHJpdmF0ZS1rZXktMw",
					Usage:      models.KeyUsageAuth,
					CreatedAt:  hourAgo.Add(-time.Minute), // Third
					ExpiresAt:  hourLater,
				},
			},

			expect: []*dao.KeyEntity{
				{
					ID:         uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					PrivateKey: "cHJpdmF0ZS1rZXktMg",
					PublicKey:  lo.ToPtr("cHVibGljLWtleS0y"),
					Usage:      models.KeyUsageAuth,
					CreatedAt:  hourAgo.Add(time.Minute), // First
					ExpiresAt:  hourLater,
				},
				{
					ID:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					PrivateKey: "cHJpdmF0ZS1rZXktMQ",
					PublicKey:  lo.ToPtr("cHVibGljLWtleS0x"),
					Usage:      models.KeyUsageAuth,
					CreatedAt:  hourAgo, // Second
					ExpiresAt:  hourLater,
				},
				{
					ID:         uuid.MustParse("00000000-0000-0000-0000-000000000003"),
					PrivateKey: "cHJpdmF0ZS1rZXktMw",
					Usage:      models.KeyUsageAuth,
					CreatedAt:  hourAgo.Add(-time.Minute), // Third
					ExpiresAt:  hourLater,
				},
			},
		},
		{
			name: "IgnoreExpiredAndDeleted",

			usage: models.KeyUsageAuth,

			fixtures: []*dao.KeyEntity{
				{
					ID:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
					PrivateKey: "cHJpdmF0ZS1rZXktMQ",
					PublicKey:  lo.ToPtr("cHVibGljLWtleS0x"),
					Usage:      models.KeyUsageAuth,
					CreatedAt:  hourAgo,
					ExpiresAt:  hourAgo,
				},
				{
					ID:             uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					PrivateKey:     "cHJpdmF0ZS1rZXktMg",
					PublicKey:      lo.ToPtr("cHVibGljLWtleS0y"),
					Usage:          models.KeyUsageAuth,
					CreatedAt:      hourAgo,
					DeletedAt:      lo.ToPtr(hourAgo),
					DeletedComment: lo.ToPtr("foo"),
					ExpiresAt:      hourLater,
				},
			},

			expect: []*dao.KeyEntity(nil),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			selectKey := dao.NewSearchKeysRepository()

			tx, commit, err := pgctx.NewContextTX(ctx, nil)
			require.NoError(t, err)

			defer func() { _ = commit(false) }()

			db, err := pgctx.Context(tx)
			require.NoError(t, err)

			for _, fixture := range testCase.fixtures {
				_, err = db.NewInsert().Model(fixture).Exec(tx)
				require.NoError(t, err)
			}

			key, err := selectKey.SearchKeys(tx, testCase.usage)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, key)
		})
	}
}
