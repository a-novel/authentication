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

func TestInsertKey(t *testing.T) {
	testCases := []struct {
		name       string
		insertData dao.InsertKeyData
	}{
		{
			name: "WithPublicKey",

			insertData: dao.InsertKeyData{
				ID:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				PrivateKey: "cHJpdmF0ZS1rZXktMQ",
				PublicKey:  lo.ToPtr("cHVibGljLWtleS0x"),
				Usage:      models.KeyUsageAuth,
				Now:        time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				Expiration: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "WithoutPublicKey",

			insertData: dao.InsertKeyData{
				ID:         uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				PrivateKey: "cHJpdmF0ZS1rZXktMQ",
				Usage:      models.KeyUsageAuth,
				Now:        time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC),
				Expiration: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			repository := dao.NewInsertKeyRepository()

			tx, commit, err := pgctx.NewContextTX(ctx, nil)
			require.NoError(t, err)

			defer func() { _ = commit(false) }()

			key, err := repository.InsertKey(tx, testCase.insertData)
			require.NoError(t, err)
			require.NotNil(t, key)
		})
	}
}
