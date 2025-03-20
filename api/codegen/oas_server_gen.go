// Code generated by ogen, DO NOT EDIT.

package codegen

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CheckSession implements checkSession operation.
	//
	// Takes an empty request with authorization headers, and check the validity of those headers. If the
	// headers can
	// be used to access any protected resource, the session is considered valid and the decoded claims
	// are returned
	// as a success response. Otherwise, an error will be sent, explaining why the session is invalid.
	//
	// GET /session
	CheckSession(ctx context.Context) (CheckSessionRes, error)
	// CreateAnonSession implements createAnonSession operation.
	//
	// Create a new anonymous session. An anonymous session is delivered without constraint, and grants
	// basic access
	// to apis with low protection.
	//
	// PUT /session/anon
	CreateAnonSession(ctx context.Context) (CreateAnonSessionRes, error)
	// CreateRefreshToken implements createRefreshToken operation.
	//
	// Issue a new refresh token. The access token used for this request must not be anonymous, and must
	// come from
	// direct login (not a refresh token).
	//
	// PUT /session/refresh
	CreateRefreshToken(ctx context.Context) (CreateRefreshTokenRes, error)
	// CreateSession implements createSession operation.
	//
	// Create a new session, using a set of credentials. The provided credentials will be used to
	// validate the
	// identity of the caller. Once the credentials have been verified, a token is issued. The access
	// rights it grants
	// may depend on the profile of the user.
	//
	// PUT /session
	CreateSession(ctx context.Context, req *LoginForm) (CreateSessionRes, error)
	// EmailExists implements emailExists operation.
	//
	// Returns an empty, successful response if the email is already associated to an user. Otherwise, it
	// fails
	// with a not found (404) error status.
	//
	// GET /credentials/email
	EmailExists(ctx context.Context, params EmailExistsParams) (EmailExistsRes, error)
	// GetPublicKey implements getPublicKey operation.
	//
	// Get a public key from its usage.
	//
	// GET /public-keys
	GetPublicKey(ctx context.Context, params GetPublicKeyParams) (GetPublicKeyRes, error)
	// Healthcheck implements healthcheck operation.
	//
	// Returns a detailed report of the health of the service, including every dependency.
	//
	// GET /healthcheck
	Healthcheck(ctx context.Context) (HealthcheckRes, error)
	// ListPublicKeys implements listPublicKeys operation.
	//
	// Get all public keys from the service that match a given usage.
	//
	// GET /public-keys/list
	ListPublicKeys(ctx context.Context, params ListPublicKeysParams) (ListPublicKeysRes, error)
	// ListUsers implements listUsers operation.
	//
	// List users in the database.
	//
	// GET /users
	ListUsers(ctx context.Context, params ListUsersParams) (ListUsersRes, error)
	// Ping implements ping operation.
	//
	// Check the status of the service. If the service is running, a successful response is returned.
	//
	// GET /ping
	Ping(ctx context.Context) (PingRes, error)
	// RefreshSession implements refreshSession operation.
	//
	// Takes a refresh token, and use it to issue a new access token.
	//
	// PATCH /session/refresh
	RefreshSession(ctx context.Context, params RefreshSessionParams) (RefreshSessionRes, error)
	// Register implements register operation.
	//
	// Create a new user. The form must contain a short code, that was sent through a registration link
	// at the user
	// desired email.
	// On success, a valid access token is returned, that can be used to access higher-privilege routes.
	//
	// PUT /credentials
	Register(ctx context.Context, req *RegisterForm) (RegisterRes, error)
	// RequestEmailUpdate implements requestEmailUpdate operation.
	//
	// Create a new short code for updating the email of an user. This short code is sent to the new
	// address.
	// If the user clicks on it, it should take it to a page that will forward the short code back to
	// this API.
	// Once done, the email associated to the user will be updated automatically.
	// This route requires to be called by an authenticated user. Anonymous sessions cannot trigger an
	// email
	// update request.
	// NOTE that this request does not verify the availability of an email. You may check this beforehand,
	//  using
	// the `Email Exists` endpoint.
	// If multiple email update links are requested for the same email, only the last one will be valid.
	//
	// PUT /short-code/update-email
	RequestEmailUpdate(ctx context.Context, req *RequestEmailUpdateForm) (RequestEmailUpdateRes, error)
	// RequestPasswordReset implements requestPasswordReset operation.
	//
	// Create a new short code for updating the password of an user. This short code is sent to the new
	// address.
	// If the user clicks on it, it should take it to a page that will forward the short code back to
	// this API.
	// Once done, the password of the user is updated.
	// This route does not require authentication (although it requires at least an anonymous session).
	// This
	// allow users who forgot their password to reset it.
	// If multiple password update links are requested for the same email, only the last one will be
	// valid.
	//
	// PUT /short-code/update-password
	RequestPasswordReset(ctx context.Context, req *RequestPasswordResetForm) (RequestPasswordResetRes, error)
	// RequestRegistration implements requestRegistration operation.
	//
	// To prevent spam in our user database, registration must be done through a link sent by e-mail, so
	// we can
	// ensure this address is valid.
	// When a user registers, the short code it received must be sent along with the registration payload.
	//  The email
	// of the payload MUST match the email the short code was sent to, and is used to retrieve the short
	// code.
	// NOTE that this request does not verify the availability of an email. You may check this beforehand,
	//  using
	// the `Email Exists` endpoint.
	// If multiple registration links are requested for the same email, only the last one will be valid.
	//
	// PUT /short-code/register
	RequestRegistration(ctx context.Context, req *RequestRegistrationForm) (RequestRegistrationRes, error)
	// ResetPassword implements resetPassword operation.
	//
	// Reset the password of an user. This route allows an unauthenticated session to update the password
	// of a user.
	// To prevent security issues, this route requires a short code that was sent to the email of the
	// user that
	// requested the password reset.
	//
	// PATCH /credentials/password/reset
	ResetPassword(ctx context.Context, req *ResetPasswordForm) (ResetPasswordRes, error)
	// UpdateEmail implements updateEmail operation.
	//
	// Update the email of an user. This route requires a valid short code, that was sent to the new
	// email.
	// If the short code is valid, the email of the user is updated with the email address the short code
	// was
	// sent to.
	//
	// PATCH /credentials/email
	UpdateEmail(ctx context.Context, req *UpdateEmailForm) (UpdateEmailRes, error)
	// UpdatePassword implements updatePassword operation.
	//
	// Update the password of an user. This route requires the original password of the user, to double
	// check
	// the identity of the caller.
	//
	// PATCH /credentials/password
	UpdatePassword(ctx context.Context, req *UpdatePasswordForm) (UpdatePasswordRes, error)
	// UpdateRole implements updateRole operation.
	//
	// Update the role of an user. This route requires the original password of the user, to double check
	// the identity of the caller.
	//
	// PATCH /credentials/role
	UpdateRole(ctx context.Context, req *UpdateRoleForm) (UpdateRoleRes, error)
	// NewError creates *UnexpectedErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *UnexpectedErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
