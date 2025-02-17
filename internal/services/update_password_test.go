package services_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"

	pgctx "github.com/a-novel-kit/context/pgbun"

	"github.com/a-novel/authentication/internal/dao"
	"github.com/a-novel/authentication/internal/lib"
	"github.com/a-novel/authentication/internal/services"
	servicesmocks "github.com/a-novel/authentication/internal/services/mocks"
	"github.com/a-novel/authentication/models"
)

func TestUpdatePassword(t *testing.T) {
	t.Parallel()

	errFoo := errors.New("foo")

	passwordRaw := "password"
	passwordScrypted, err := lib.GenerateScrypt(passwordRaw, lib.ScryptParamsDefault)
	require.NoError(t, err)

	type consumeShortCodeData struct {
		resp *models.ShortCode
		err  error
	}

	type selectCredentialsData struct {
		resp *dao.CredentialsEntity
		err  error
	}

	type updateCredentialsPasswordData struct {
		resp *dao.CredentialsEntity
		err  error
	}

	testCases := []struct {
		name string

		request services.UpdatePasswordRequest

		consumeShortCodeData          *consumeShortCodeData
		selectCredentialsData         *selectCredentialsData
		updateCredentialsPasswordData *updateCredentialsPasswordData

		expectErr error
	}{
		{
			name: "Success/ShorCode",

			request: services.UpdatePasswordRequest{
				UserID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Password:  "new-password",
				ShortCode: "short-code",
			},

			consumeShortCodeData: &consumeShortCodeData{
				resp: &models.ShortCode{},
			},

			updateCredentialsPasswordData: &updateCredentialsPasswordData{
				resp: &dao.CredentialsEntity{},
			},
		},
		{
			name: "Success/CurrentPassword",

			request: services.UpdatePasswordRequest{
				UserID:          uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Password:        "new-password",
				CurrentPassword: passwordRaw,
			},

			selectCredentialsData: &selectCredentialsData{
				resp: &dao.CredentialsEntity{
					Password: passwordScrypted,
				},
			},

			updateCredentialsPasswordData: &updateCredentialsPasswordData{
				resp: &dao.CredentialsEntity{},
			},
		},
		{
			name: "Error/UploadCredentials",

			request: services.UpdatePasswordRequest{
				UserID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Password:  "new-password",
				ShortCode: "short-code",
			},

			consumeShortCodeData: &consumeShortCodeData{
				resp: &models.ShortCode{},
			},

			updateCredentialsPasswordData: &updateCredentialsPasswordData{
				err: errFoo,
			},

			expectErr: errFoo,
		},
		{
			name: "Error/ConsumesShortCode",

			request: services.UpdatePasswordRequest{
				UserID:    uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Password:  "new-password",
				ShortCode: "short-code",
			},

			consumeShortCodeData: &consumeShortCodeData{
				err: errFoo,
			},

			expectErr: errFoo,
		},
		{
			name: "Error/SelectCredentials",

			request: services.UpdatePasswordRequest{
				UserID:          uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Password:        "new-password",
				CurrentPassword: passwordRaw,
			},

			selectCredentialsData: &selectCredentialsData{
				err: errFoo,
			},

			expectErr: errFoo,
		},
		{
			name: "Error/WrongPassword",

			request: services.UpdatePasswordRequest{
				UserID:          uuid.MustParse("00000000-0000-0000-0000-000000000001"),
				Password:        "new-password",
				CurrentPassword: "fake-password",
			},

			selectCredentialsData: &selectCredentialsData{
				resp: &dao.CredentialsEntity{
					Password: passwordScrypted,
				},
			},

			expectErr: lib.ErrInvalidPassword,
		},
		{
			name: "Error/MissingShortCodeAndCurrentPassword",

			request: services.UpdatePasswordRequest{
				UserID: uuid.MustParse("00000000-0000-0000-0000-000000000001"),
			},

			expectErr: services.ErrMissingShortCodeAndCurrentPassword,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			ctx, err := pgctx.NewContext(t.Context(), nil)
			require.NoError(t, err)

			source := servicesmocks.NewMockUpdatePasswordSource(t)

			if testCase.consumeShortCodeData != nil {
				source.
					On("ConsumeShortCode", mock.Anything, services.ConsumeShortCodeRequest{
						Usage:  models.ShortCodeUsageResetPassword,
						Target: testCase.request.UserID.String(),
						Code:   testCase.request.ShortCode,
					}).
					Return(testCase.consumeShortCodeData.resp, testCase.consumeShortCodeData.err)
			}

			if testCase.selectCredentialsData != nil {
				source.
					On("SelectCredentials", mock.Anything, testCase.request.UserID).
					Return(testCase.selectCredentialsData.resp, testCase.selectCredentialsData.err)
			}

			if testCase.updateCredentialsPasswordData != nil {
				source.
					On(
						"UpdateCredentialsPassword",
						mock.Anything, testCase.request.UserID,
						mock.AnythingOfType("dao.UpdateCredentialsPasswordData"),
					).
					Return(testCase.updateCredentialsPasswordData.resp, testCase.updateCredentialsPasswordData.err)
			}

			service := services.NewUpdatePasswordService(source)

			err = service.UpdatePassword(ctx, testCase.request)
			require.ErrorIs(t, err, testCase.expectErr)

			source.AssertExpectations(t)
		})
	}
}
