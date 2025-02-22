package services

import (
	"errors"
	"fmt"

	"github.com/samber/lo"
	"golang.org/x/crypto/ed25519"

	"github.com/a-novel-kit/context"
	"github.com/a-novel-kit/jwt"
	"github.com/a-novel-kit/jwt/jwk"
	"github.com/a-novel-kit/jwt/jwp"
	"github.com/a-novel-kit/jwt/jws"

	"github.com/a-novel/authentication/config"
	"github.com/a-novel/authentication/models"
)

var (
	ErrMismatchRefreshClaims                = errors.New("refresh token and access token do not match")
	ErrTokenIssuedWithDifferentRefreshToken = errors.New("access token was not issued with the provided refresh token")

	ErrConsumeRefreshTokenService = errors.New("ConsumeRefreshTokenService.ConsumeRefreshToken")
)

func NewErrConsumeRefreshTokenService(err error) error {
	return errors.Join(err, ErrConsumeRefreshTokenService)
}

type ConsumeRefreshTokenSource interface {
	IssueToken(ctx context.Context, request IssueTokenRequest) (string, error)
}

type ConsumeRefreshTokenRequest struct {
	// The last valid access token must be provided as an extra security measure. It may be expired, but must still
	// be signed by an active key. Information in the access token is checked against the refresh token claims, and
	// used to populate the new access token claims.
	AccessToken  string
	RefreshToken string
}

type ConsumeRefreshTokenService struct {
	source                ConsumeRefreshTokenSource
	accessTokenRecipient  *jwt.Recipient
	refreshTokenRecipient *jwt.Recipient
}

func (service *ConsumeRefreshTokenService) ConsumeRefreshToken(
	ctx context.Context, request ConsumeRefreshTokenRequest,
) (string, error) {
	if request.AccessToken == "" {
		return "", NewErrConsumeRefreshTokenService(fmt.Errorf("%w: access token is empty", models.ErrUnauthorized))
	}

	if request.RefreshToken == "" {
		return "", NewErrConsumeRefreshTokenService(fmt.Errorf("%w: refresh token is empty", models.ErrUnauthorized))
	}

	var accessTokenClaims models.AccessTokenClaims
	if err := service.accessTokenRecipient.Consume(ctx, request.AccessToken, &accessTokenClaims); err != nil {
		if errors.Is(err, jws.ErrInvalidSignature) {
			return "", NewErrConsumeRefreshTokenService(fmt.Errorf("consume access token: %w", models.ErrUnauthorized))
		}

		return "", NewErrConsumeRefreshTokenService(fmt.Errorf("consume access token: %w", err))
	}

	var refreshTokenClaims models.RefreshTokenClaims
	if err := service.refreshTokenRecipient.Consume(ctx, request.RefreshToken, &refreshTokenClaims); err != nil {
		if errors.Is(err, jws.ErrInvalidSignature) {
			return "", NewErrConsumeRefreshTokenService(fmt.Errorf("consume refresh token: %w", models.ErrUnauthorized))
		}

		return "", NewErrConsumeRefreshTokenService(fmt.Errorf("consume refresh token: %w", err))
	}

	if lo.FromPtr(accessTokenClaims.UserID) != refreshTokenClaims.UserID {
		return "", NewErrConsumeRefreshTokenService(fmt.Errorf(
			"%w (accessToken.userID: %s, refreshToken.userID: %s)",
			ErrMismatchRefreshClaims,
			lo.FromPtr(accessTokenClaims.UserID),
			refreshTokenClaims.UserID,
		))
	}

	if accessTokenClaims.RefreshTokenID != nil && *accessTokenClaims.RefreshTokenID != refreshTokenClaims.Jti {
		return "", NewErrConsumeRefreshTokenService(ErrTokenIssuedWithDifferentRefreshToken)
	}

	accessToken, err := service.source.IssueToken(ctx, IssueTokenRequest{
		UserID:         accessTokenClaims.UserID,
		Roles:          accessTokenClaims.Roles,
		RefreshTokenID: &refreshTokenClaims.Jti,
	})
	if err != nil {
		return "", NewErrConsumeRefreshTokenService(fmt.Errorf("issue accessToken: %w", err))
	}

	return accessToken, nil
}

func NewConsumeRefreshTokenService(
	source ConsumeRefreshTokenSource,
	accessTokenKeysSource *jwk.Source[ed25519.PublicKey],
	refreshTokenKeysSource *jwk.Source[ed25519.PublicKey],
) *ConsumeRefreshTokenService {
	accessTokenVerifier := jws.NewSourcedED25519Verifier(accessTokenKeysSource)
	refreshTokenVerifier := jws.NewSourcedED25519Verifier(refreshTokenKeysSource)

	accessTokenDeserializer := jwp.NewClaimsChecker(&jwp.ClaimsCheckerConfig{
		Checks: []jwp.ClaimsCheck{
			jwp.NewClaimsCheckTarget(jwt.TargetConfig{
				Issuer:   config.Tokens.Usages[models.KeyUsageAuth].Issuer,
				Audience: config.Tokens.Usages[models.KeyUsageAuth].Audience,
				Subject:  config.Tokens.Usages[models.KeyUsageAuth].Subject,
			}),
			// Ignore timestamps checks. We just need to ensure this service issued the token, not that it's
			// still valid.
		},
	})
	refreshTokenDeserializer := jwp.NewClaimsChecker(&jwp.ClaimsCheckerConfig{
		Checks: []jwp.ClaimsCheck{
			jwp.NewClaimsCheckTarget(jwt.TargetConfig{
				Issuer:   config.Tokens.Usages[models.KeyUsageRefresh].Issuer,
				Audience: config.Tokens.Usages[models.KeyUsageRefresh].Audience,
				Subject:  config.Tokens.Usages[models.KeyUsageRefresh].Subject,
			}),
			jwp.NewClaimsCheckTimestamp(config.Tokens.Usages[models.KeyUsageRefresh].Leeway, true),
		},
	})

	return &ConsumeRefreshTokenService{
		source: source,
		accessTokenRecipient: jwt.NewRecipient(jwt.RecipientConfig{
			Plugins:      []jwt.RecipientPlugin{accessTokenVerifier},
			Deserializer: accessTokenDeserializer.Unmarshal,
		}),
		refreshTokenRecipient: jwt.NewRecipient(jwt.RecipientConfig{
			Plugins:      []jwt.RecipientPlugin{refreshTokenVerifier},
			Deserializer: refreshTokenDeserializer.Unmarshal,
		}),
	}
}
