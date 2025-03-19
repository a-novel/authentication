package api_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/a-novel/authentication/api"
	"github.com/a-novel/authentication/api/codegen"
	apimocks "github.com/a-novel/authentication/api/mocks"
	"github.com/a-novel/authentication/internal/services"
)

func TestRequestPasswordReset(t *testing.T) {
	t.Parallel()

	errFoo := errors.New("foo")

	type requestPasswordResetData struct {
		err error
	}

	testCases := []struct {
		name string

		form *codegen.RequestPasswordResetForm

		requestPasswordResetData *requestPasswordResetData

		expect    codegen.RequestPasswordResetRes
		expectErr error
	}{
		{
			name: "Success",

			form: &codegen.RequestPasswordResetForm{
				Email: "user@provider.com",
			},

			requestPasswordResetData: &requestPasswordResetData{},

			expect: &codegen.RequestPasswordResetNoContent{},
		},
		{
			name: "RequestPasswordResetError",

			form: &codegen.RequestPasswordResetForm{
				Email: "user@provider.com",
			},

			requestPasswordResetData: &requestPasswordResetData{
				err: errFoo,
			},

			expectErr: errFoo,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			t.Parallel()

			source := apimocks.NewMockRequestPasswordResetService(t)

			if testCase.requestPasswordResetData != nil {
				source.EXPECT().
					RequestPasswordReset(t.Context(), services.RequestPasswordResetRequest{
						Email: string(testCase.form.GetEmail()),
					}).
					Return(nil, testCase.requestPasswordResetData.err)
			}

			handler := api.API{RequestPasswordResetService: source}

			res, err := handler.RequestPasswordReset(t.Context(), testCase.form)
			require.ErrorIs(t, err, testCase.expectErr)
			require.Equal(t, testCase.expect, res)

			source.AssertExpectations(t)
		})
	}
}
