package api

import (
	"errors"
	"fmt"

	"github.com/a-novel-kit/context"

	"github.com/a-novel/authentication/api/codegen"
	"github.com/a-novel/authentication/internal/lib"
	"github.com/a-novel/authentication/internal/services"
)

func (api *API) UpdatePassword(
	ctx context.Context, req *codegen.UpdatePasswordForm) (codegen.UpdatePasswordRes, error,
) {
	userID, err := RequireUserID(ctx)
	if err != nil {
		return nil, fmt.Errorf("update password: %w", err)
	}

	err = api.UpdatePasswordService.UpdatePassword(ctx, services.UpdatePasswordRequest{
		Password:        string(req.Password),
		CurrentPassword: string(req.CurrentPassword),
		UserID:          userID,
	})

	switch {
	case errors.Is(err, lib.ErrInvalidPassword):
		return &codegen.ForbiddenError{Error: "invalid user password"}, nil
	case err != nil:
		return nil, fmt.Errorf("update password: %w", err)
	}

	return &codegen.UpdatePasswordNoContent{}, nil
}
