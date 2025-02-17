package dao_test

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/stretchr/testify/require"

	pgctx "github.com/a-novel-kit/context/pgbun"

	"github.com/a-novel/authentication/internal/dao"
)

func TestInsertShortCode(t *testing.T) {
	now := time.Now().UTC().Round(time.Second)
	hourAgo := time.Now().Add(-time.Hour).UTC().Round(time.Second)
	hourLater := time.Now().Add(time.Hour).UTC().Round(time.Second)

	testCases := []struct {
		name string

		fixtures []*dao.ShortCodeEntity

		insertData dao.InsertShortCodeData

		expect    *dao.ShortCodeEntity
		expectErr error
	}{
		{
			name: "Success",

			insertData: dao.InsertShortCodeData{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				Data:      []byte("test-data"),
				Now:       now,
				ExpiresAt: hourLater,
			},

			expect: &dao.ShortCodeEntity{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				Data:      []byte("test-data"),
				CreatedAt: now,
				ExpiresAt: hourLater,
			},
		},
		{
			name: "Success/NoData",

			insertData: dao.InsertShortCodeData{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				Now:       now,
				ExpiresAt: hourLater,
			},

			expect: &dao.ShortCodeEntity{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				CreatedAt: now,
				ExpiresAt: hourLater,
			},
		},
		{
			name: "SameTargetDifferentUsage",

			insertData: dao.InsertShortCodeData{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				Data:      []byte("test-data"),
				Now:       now,
				ExpiresAt: hourLater,
			},

			fixtures: []*dao.ShortCodeEntity{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					Code:      "test-code",
					Usage:     "test-usage-2",
					Target:    "test-target",
					Data:      []byte("test-data-2"),
					CreatedAt: hourAgo,
					ExpiresAt: hourLater,
				},
			},

			expect: &dao.ShortCodeEntity{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				Data:      []byte("test-data"),
				CreatedAt: now,
				ExpiresAt: hourLater,
			},
		},
		{
			name: "SameUsageDifferentTarget",

			insertData: dao.InsertShortCodeData{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				Data:      []byte("test-data"),
				Now:       now,
				ExpiresAt: hourLater,
			},

			fixtures: []*dao.ShortCodeEntity{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					Code:      "test-code",
					Usage:     "test-usage",
					Target:    "test-target-2",
					Data:      []byte("test-data-2"),
					CreatedAt: hourAgo,
					ExpiresAt: hourLater,
				},
			},

			expect: &dao.ShortCodeEntity{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				Data:      []byte("test-data"),
				CreatedAt: now,
				ExpiresAt: hourLater,
			},
		},
		{
			name: "Error/AlreadyExists",

			insertData: dao.InsertShortCodeData{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				Data:      []byte("test-data"),
				Now:       now,
				ExpiresAt: hourLater,
			},

			fixtures: []*dao.ShortCodeEntity{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					Code:      "test-code-2",
					Usage:     "test-usage",
					Target:    "test-target",
					Data:      []byte("test-data-2"),
					CreatedAt: hourAgo,
					ExpiresAt: hourLater,
				},
			},

			expectErr: dao.ErrShortCodeAlreadyExists,
		},
		{
			name: "AlreadyExists/Override",

			insertData: dao.InsertShortCodeData{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				Data:      []byte("test-data"),
				Now:       now,
				ExpiresAt: hourLater,
				Override:  true,
			},

			fixtures: []*dao.ShortCodeEntity{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					Code:      "test-code",
					Usage:     "test-usage",
					Target:    "test-target",
					Data:      []byte("test-data-2"),
					CreatedAt: hourAgo,
					ExpiresAt: hourLater,
				},
			},

			expect: &dao.ShortCodeEntity{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				Data:      []byte("test-data"),
				CreatedAt: now,
				ExpiresAt: hourLater,
			},
		},
		{
			name: "AlreadyExists/AlreadyExpired",

			insertData: dao.InsertShortCodeData{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				Data:      []byte("test-data"),
				Now:       now,
				ExpiresAt: hourLater,
			},

			fixtures: []*dao.ShortCodeEntity{
				{
					ID:        uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					Code:      "test-code",
					Usage:     "test-usage",
					Target:    "test-target",
					Data:      []byte("test-data-2"),
					CreatedAt: hourAgo,
					ExpiresAt: hourAgo,
				},
			},

			expect: &dao.ShortCodeEntity{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				Data:      []byte("test-data"),
				CreatedAt: now,
				ExpiresAt: hourLater,
			},
		},
		{
			name: "AlreadyExistsAlreadyDeleted",

			insertData: dao.InsertShortCodeData{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				Data:      []byte("test-data"),
				Now:       now,
				ExpiresAt: hourLater,
			},

			fixtures: []*dao.ShortCodeEntity{
				{
					ID:             uuid.MustParse("00000000-0000-0000-0000-000000000002"),
					Code:           "test-code",
					Usage:          "test-usage",
					Target:         "test-target",
					Data:           []byte("test-data-2"),
					CreatedAt:      hourAgo,
					ExpiresAt:      hourLater,
					DeletedAt:      &hourAgo,
					DeletedComment: lo.ToPtr("test-comment"),
				},
			},

			expect: &dao.ShortCodeEntity{
				ID:        uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Code:      "test-code",
				Usage:     "test-usage",
				Target:    "test-target",
				Data:      []byte("test-data"),
				CreatedAt: now,
				ExpiresAt: hourLater,
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repository := dao.NewInsertShortCodeRepository()

			tx, commit, err := pgctx.NewContextTX(ctx, nil)
			require.NoError(t, err)

			defer func() { _ = commit(false) }()

			db, err := pgctx.Context(tx)
			require.NoError(t, err)

			for _, fixture := range testCase.fixtures {
				_, err = db.NewInsert().Model(fixture).Exec(tx)
				require.NoError(t, err)
			}

			credentials, err := repository.InsertShortCode(tx, testCase.insertData)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, credentials)
		})
	}
}
