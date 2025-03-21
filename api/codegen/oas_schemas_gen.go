// Code generated by ogen, DO NOT EDIT.

package codegen

import (
	"fmt"
	"io"
	"time"

	"github.com/go-faster/errors"
	"github.com/go-faster/jx"
	"github.com/google/uuid"
)

func (s *UnexpectedErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// The algorithm this key can be used for.
// Ref: #/components/schemas/Alg
type Alg string

const (
	AlgHS256            Alg = "HS256"
	AlgHS384            Alg = "HS384"
	AlgHS512            Alg = "HS512"
	AlgRS256            Alg = "RS256"
	AlgRS384            Alg = "RS384"
	AlgRS512            Alg = "RS512"
	AlgES256            Alg = "ES256"
	AlgES384            Alg = "ES384"
	AlgES512            Alg = "ES512"
	AlgPS256            Alg = "PS256"
	AlgPS384            Alg = "PS384"
	AlgPS512            Alg = "PS512"
	AlgEdDSA            Alg = "EdDSA"
	AlgRSAOAEP          Alg = "RSA-OAEP"
	AlgRSAOAEP256       Alg = "RSA-OAEP-256"
	AlgA128KW           Alg = "A128KW"
	AlgA192KW           Alg = "A192KW"
	AlgA256KW           Alg = "A256KW"
	AlgDir              Alg = "dir"
	AlgECDHES           Alg = "ECDH-ES"
	AlgECDHESA128KW     Alg = "ECDH-ES+A128KW"
	AlgECDHESA192KW     Alg = "ECDH-ES+A192KW"
	AlgECDHESA256KW     Alg = "ECDH-ES+A256KW"
	AlgA128GCMKW        Alg = "A128GCMKW"
	AlgA192GCMKW        Alg = "A192GCMKW"
	AlgA256GCMKW        Alg = "A256GCMKW"
	AlgPBES2HS256A128KW Alg = "PBES2-HS256+A128KW"
	AlgPBES2HS384A192KW Alg = "PBES2-HS384+A192KW"
	AlgPBES2HS512A256KW Alg = "PBES2-HS512+A256KW"
	AlgA128CBCHS256     Alg = "A128CBC-HS256"
	AlgA192CBCHS384     Alg = "A192CBC-HS384"
	AlgA256CBCHS512     Alg = "A256CBC-HS512"
	AlgA128GCM          Alg = "A128GCM"
	AlgA192GCM          Alg = "A192GCM"
	AlgA256GCM          Alg = "A256GCM"
)

// AllValues returns all Alg values.
func (Alg) AllValues() []Alg {
	return []Alg{
		AlgHS256,
		AlgHS384,
		AlgHS512,
		AlgRS256,
		AlgRS384,
		AlgRS512,
		AlgES256,
		AlgES384,
		AlgES512,
		AlgPS256,
		AlgPS384,
		AlgPS512,
		AlgEdDSA,
		AlgRSAOAEP,
		AlgRSAOAEP256,
		AlgA128KW,
		AlgA192KW,
		AlgA256KW,
		AlgDir,
		AlgECDHES,
		AlgECDHESA128KW,
		AlgECDHESA192KW,
		AlgECDHESA256KW,
		AlgA128GCMKW,
		AlgA192GCMKW,
		AlgA256GCMKW,
		AlgPBES2HS256A128KW,
		AlgPBES2HS384A192KW,
		AlgPBES2HS512A256KW,
		AlgA128CBCHS256,
		AlgA192CBCHS384,
		AlgA256CBCHS512,
		AlgA128GCM,
		AlgA192GCM,
		AlgA256GCM,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s Alg) MarshalText() ([]byte, error) {
	switch s {
	case AlgHS256:
		return []byte(s), nil
	case AlgHS384:
		return []byte(s), nil
	case AlgHS512:
		return []byte(s), nil
	case AlgRS256:
		return []byte(s), nil
	case AlgRS384:
		return []byte(s), nil
	case AlgRS512:
		return []byte(s), nil
	case AlgES256:
		return []byte(s), nil
	case AlgES384:
		return []byte(s), nil
	case AlgES512:
		return []byte(s), nil
	case AlgPS256:
		return []byte(s), nil
	case AlgPS384:
		return []byte(s), nil
	case AlgPS512:
		return []byte(s), nil
	case AlgEdDSA:
		return []byte(s), nil
	case AlgRSAOAEP:
		return []byte(s), nil
	case AlgRSAOAEP256:
		return []byte(s), nil
	case AlgA128KW:
		return []byte(s), nil
	case AlgA192KW:
		return []byte(s), nil
	case AlgA256KW:
		return []byte(s), nil
	case AlgDir:
		return []byte(s), nil
	case AlgECDHES:
		return []byte(s), nil
	case AlgECDHESA128KW:
		return []byte(s), nil
	case AlgECDHESA192KW:
		return []byte(s), nil
	case AlgECDHESA256KW:
		return []byte(s), nil
	case AlgA128GCMKW:
		return []byte(s), nil
	case AlgA192GCMKW:
		return []byte(s), nil
	case AlgA256GCMKW:
		return []byte(s), nil
	case AlgPBES2HS256A128KW:
		return []byte(s), nil
	case AlgPBES2HS384A192KW:
		return []byte(s), nil
	case AlgPBES2HS512A256KW:
		return []byte(s), nil
	case AlgA128CBCHS256:
		return []byte(s), nil
	case AlgA192CBCHS384:
		return []byte(s), nil
	case AlgA256CBCHS512:
		return []byte(s), nil
	case AlgA128GCM:
		return []byte(s), nil
	case AlgA192GCM:
		return []byte(s), nil
	case AlgA256GCM:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *Alg) UnmarshalText(data []byte) error {
	switch Alg(data) {
	case AlgHS256:
		*s = AlgHS256
		return nil
	case AlgHS384:
		*s = AlgHS384
		return nil
	case AlgHS512:
		*s = AlgHS512
		return nil
	case AlgRS256:
		*s = AlgRS256
		return nil
	case AlgRS384:
		*s = AlgRS384
		return nil
	case AlgRS512:
		*s = AlgRS512
		return nil
	case AlgES256:
		*s = AlgES256
		return nil
	case AlgES384:
		*s = AlgES384
		return nil
	case AlgES512:
		*s = AlgES512
		return nil
	case AlgPS256:
		*s = AlgPS256
		return nil
	case AlgPS384:
		*s = AlgPS384
		return nil
	case AlgPS512:
		*s = AlgPS512
		return nil
	case AlgEdDSA:
		*s = AlgEdDSA
		return nil
	case AlgRSAOAEP:
		*s = AlgRSAOAEP
		return nil
	case AlgRSAOAEP256:
		*s = AlgRSAOAEP256
		return nil
	case AlgA128KW:
		*s = AlgA128KW
		return nil
	case AlgA192KW:
		*s = AlgA192KW
		return nil
	case AlgA256KW:
		*s = AlgA256KW
		return nil
	case AlgDir:
		*s = AlgDir
		return nil
	case AlgECDHES:
		*s = AlgECDHES
		return nil
	case AlgECDHESA128KW:
		*s = AlgECDHESA128KW
		return nil
	case AlgECDHESA192KW:
		*s = AlgECDHESA192KW
		return nil
	case AlgECDHESA256KW:
		*s = AlgECDHESA256KW
		return nil
	case AlgA128GCMKW:
		*s = AlgA128GCMKW
		return nil
	case AlgA192GCMKW:
		*s = AlgA192GCMKW
		return nil
	case AlgA256GCMKW:
		*s = AlgA256GCMKW
		return nil
	case AlgPBES2HS256A128KW:
		*s = AlgPBES2HS256A128KW
		return nil
	case AlgPBES2HS384A192KW:
		*s = AlgPBES2HS384A192KW
		return nil
	case AlgPBES2HS512A256KW:
		*s = AlgPBES2HS512A256KW
		return nil
	case AlgA128CBCHS256:
		*s = AlgA128CBCHS256
		return nil
	case AlgA192CBCHS384:
		*s = AlgA192CBCHS384
		return nil
	case AlgA256CBCHS512:
		*s = AlgA256CBCHS512
		return nil
	case AlgA128GCM:
		*s = AlgA128GCM
		return nil
	case AlgA192GCM:
		*s = AlgA192GCM
		return nil
	case AlgA256GCM:
		*s = AlgA256GCM
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type BearerAuth struct {
	Token string
}

// GetToken returns the value of Token.
func (s *BearerAuth) GetToken() string {
	return s.Token
}

// SetToken sets the value of Token.
func (s *BearerAuth) SetToken(val string) {
	s.Token = val
}

// Ref: #/components/schemas/Claims
type Claims struct {
	// The unique identifier of the user. Can be null if the session is anonymous.
	UserID OptUUID `json:"userID"`
	// The roles granted by the token.
	Roles []Role `json:"roles"`
	// The unique identifier of the refresh token. Can be null if the session is anonymous.
	RefreshTokenID OptString `json:"refreshTokenID"`
}

// GetUserID returns the value of UserID.
func (s *Claims) GetUserID() OptUUID {
	return s.UserID
}

// GetRoles returns the value of Roles.
func (s *Claims) GetRoles() []Role {
	return s.Roles
}

// GetRefreshTokenID returns the value of RefreshTokenID.
func (s *Claims) GetRefreshTokenID() OptString {
	return s.RefreshTokenID
}

// SetUserID sets the value of UserID.
func (s *Claims) SetUserID(val OptUUID) {
	s.UserID = val
}

// SetRoles sets the value of Roles.
func (s *Claims) SetRoles(val []Role) {
	s.Roles = val
}

// SetRefreshTokenID sets the value of RefreshTokenID.
func (s *Claims) SetRefreshTokenID(val OptString) {
	s.RefreshTokenID = val
}

func (*Claims) checkSessionRes() {}

// Ref: #/components/schemas/ConflictError
type ConflictError struct {
	// The error message.
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *ConflictError) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *ConflictError) SetError(val string) {
	s.Error = val
}

func (*ConflictError) registerRes()    {}
func (*ConflictError) updateEmailRes() {}

// CreateAnonSessionIMATeapot is response for CreateAnonSession operation.
type CreateAnonSessionIMATeapot struct{}

func (*CreateAnonSessionIMATeapot) createAnonSessionRes() {}

// A role specifically assigned to a user.
// Ref: #/components/schemas/CredentialsRole
type CredentialsRole string

const (
	CredentialsRoleUser       CredentialsRole = "user"
	CredentialsRoleAdmin      CredentialsRole = "admin"
	CredentialsRoleSuperAdmin CredentialsRole = "super_admin"
)

// AllValues returns all CredentialsRole values.
func (CredentialsRole) AllValues() []CredentialsRole {
	return []CredentialsRole{
		CredentialsRoleUser,
		CredentialsRoleAdmin,
		CredentialsRoleSuperAdmin,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s CredentialsRole) MarshalText() ([]byte, error) {
	switch s {
	case CredentialsRoleUser:
		return []byte(s), nil
	case CredentialsRoleAdmin:
		return []byte(s), nil
	case CredentialsRoleSuperAdmin:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *CredentialsRole) UnmarshalText(data []byte) error {
	switch CredentialsRole(data) {
	case CredentialsRoleUser:
		*s = CredentialsRoleUser
		return nil
	case CredentialsRoleAdmin:
		*s = CredentialsRoleAdmin
		return nil
	case CredentialsRoleSuperAdmin:
		*s = CredentialsRoleSuperAdmin
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/Dependency
type Dependency struct {
	// The name of the dependency.
	Name            string           `json:"name"`
	Status          DependencyStatus `json:"status"`
	AdditionalProps DependencyAdditional
}

// GetName returns the value of Name.
func (s *Dependency) GetName() string {
	return s.Name
}

// GetStatus returns the value of Status.
func (s *Dependency) GetStatus() DependencyStatus {
	return s.Status
}

// GetAdditionalProps returns the value of AdditionalProps.
func (s *Dependency) GetAdditionalProps() DependencyAdditional {
	return s.AdditionalProps
}

// SetName sets the value of Name.
func (s *Dependency) SetName(val string) {
	s.Name = val
}

// SetStatus sets the value of Status.
func (s *Dependency) SetStatus(val DependencyStatus) {
	s.Status = val
}

// SetAdditionalProps sets the value of AdditionalProps.
func (s *Dependency) SetAdditionalProps(val DependencyAdditional) {
	s.AdditionalProps = val
}

type DependencyAdditional map[string]jx.Raw

func (s *DependencyAdditional) init() DependencyAdditional {
	m := *s
	if m == nil {
		m = map[string]jx.Raw{}
		*s = m
	}
	return m
}

// The status of a dependency.
// Ref: #/components/schemas/DependencyStatus
type DependencyStatus string

const (
	DependencyStatusUp      DependencyStatus = "up"
	DependencyStatusDown    DependencyStatus = "down"
	DependencyStatusUnknown DependencyStatus = "unknown"
)

// AllValues returns all DependencyStatus values.
func (DependencyStatus) AllValues() []DependencyStatus {
	return []DependencyStatus{
		DependencyStatusUp,
		DependencyStatusDown,
		DependencyStatusUnknown,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s DependencyStatus) MarshalText() ([]byte, error) {
	switch s {
	case DependencyStatusUp:
		return []byte(s), nil
	case DependencyStatusDown:
		return []byte(s), nil
	case DependencyStatusUnknown:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *DependencyStatus) UnmarshalText(data []byte) error {
	switch DependencyStatus(data) {
	case DependencyStatusUp:
		*s = DependencyStatusUp
		return nil
	case DependencyStatusDown:
		*s = DependencyStatusDown
		return nil
	case DependencyStatusUnknown:
		*s = DependencyStatusUnknown
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

type Email string

// EmailExistsNoContent is response for EmailExists operation.
type EmailExistsNoContent struct{}

func (*EmailExistsNoContent) emailExistsRes() {}

// Ref: #/components/schemas/ForbiddenError
type ForbiddenError struct {
	// The error message.
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *ForbiddenError) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *ForbiddenError) SetError(val string) {
	s.Error = val
}

func (*ForbiddenError) checkSessionRes()       {}
func (*ForbiddenError) createRefreshTokenRes() {}
func (*ForbiddenError) createSessionRes()      {}
func (*ForbiddenError) refreshSessionRes()     {}
func (*ForbiddenError) registerRes()           {}
func (*ForbiddenError) resetPasswordRes()      {}
func (*ForbiddenError) updateEmailRes()        {}
func (*ForbiddenError) updatePasswordRes()     {}
func (*ForbiddenError) updateRoleRes()         {}

// Ref: #/components/schemas/Health
type Health struct {
	Postgres Dependency `json:"postgres"`
	Sendgrid Dependency `json:"sendgrid"`
}

// GetPostgres returns the value of Postgres.
func (s *Health) GetPostgres() Dependency {
	return s.Postgres
}

// GetSendgrid returns the value of Sendgrid.
func (s *Health) GetSendgrid() Dependency {
	return s.Sendgrid
}

// SetPostgres sets the value of Postgres.
func (s *Health) SetPostgres(val Dependency) {
	s.Postgres = val
}

// SetSendgrid sets the value of Sendgrid.
func (s *Health) SetSendgrid(val Dependency) {
	s.Sendgrid = val
}

func (*Health) healthcheckRes() {}

// HealthcheckIMATeapot is response for Healthcheck operation.
type HealthcheckIMATeapot struct{}

func (*HealthcheckIMATeapot) healthcheckRes() {}

// Ref: #/components/schemas/JWK
type JWK struct {
	Kty             KTY     `json:"kty"`
	Use             Use     `json:"use"`
	KeyOps          []KeyOp `json:"key_ops"`
	Alg             Alg     `json:"alg"`
	Kid             OptKID  `json:"kid"`
	AdditionalProps JWKAdditional
}

// GetKty returns the value of Kty.
func (s *JWK) GetKty() KTY {
	return s.Kty
}

// GetUse returns the value of Use.
func (s *JWK) GetUse() Use {
	return s.Use
}

// GetKeyOps returns the value of KeyOps.
func (s *JWK) GetKeyOps() []KeyOp {
	return s.KeyOps
}

// GetAlg returns the value of Alg.
func (s *JWK) GetAlg() Alg {
	return s.Alg
}

// GetKid returns the value of Kid.
func (s *JWK) GetKid() OptKID {
	return s.Kid
}

// GetAdditionalProps returns the value of AdditionalProps.
func (s *JWK) GetAdditionalProps() JWKAdditional {
	return s.AdditionalProps
}

// SetKty sets the value of Kty.
func (s *JWK) SetKty(val KTY) {
	s.Kty = val
}

// SetUse sets the value of Use.
func (s *JWK) SetUse(val Use) {
	s.Use = val
}

// SetKeyOps sets the value of KeyOps.
func (s *JWK) SetKeyOps(val []KeyOp) {
	s.KeyOps = val
}

// SetAlg sets the value of Alg.
func (s *JWK) SetAlg(val Alg) {
	s.Alg = val
}

// SetKid sets the value of Kid.
func (s *JWK) SetKid(val OptKID) {
	s.Kid = val
}

// SetAdditionalProps sets the value of AdditionalProps.
func (s *JWK) SetAdditionalProps(val JWKAdditional) {
	s.AdditionalProps = val
}

func (*JWK) getPublicKeyRes() {}

type JWKAdditional map[string]jx.Raw

func (s *JWKAdditional) init() JWKAdditional {
	m := *s
	if m == nil {
		m = map[string]jx.Raw{}
		*s = m
	}
	return m
}

type KID uuid.UUID

// The type of the key embedded in the JWK.
// Ref: #/components/schemas/KTY
type KTY string

const (
	KTYOct KTY = "oct"
	KTYRSA KTY = "RSA"
	KTYEC  KTY = "EC"
	KTYOKP KTY = "OKP"
)

// AllValues returns all KTY values.
func (KTY) AllValues() []KTY {
	return []KTY{
		KTYOct,
		KTYRSA,
		KTYEC,
		KTYOKP,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s KTY) MarshalText() ([]byte, error) {
	switch s {
	case KTYOct:
		return []byte(s), nil
	case KTYRSA:
		return []byte(s), nil
	case KTYEC:
		return []byte(s), nil
	case KTYOKP:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *KTY) UnmarshalText(data []byte) error {
	switch KTY(data) {
	case KTYOct:
		*s = KTYOct
		return nil
	case KTYRSA:
		*s = KTYRSA
		return nil
	case KTYEC:
		*s = KTYEC
		return nil
	case KTYOKP:
		*s = KTYOKP
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// The operation that can be performed using the key.
// Ref: #/components/schemas/KeyOp
type KeyOp string

const (
	KeyOpSign       KeyOp = "sign"
	KeyOpVerify     KeyOp = "verify"
	KeyOpEncrypt    KeyOp = "encrypt"
	KeyOpDecrypt    KeyOp = "decrypt"
	KeyOpWrapKey    KeyOp = "wrapKey"
	KeyOpUnwrapKey  KeyOp = "unwrapKey"
	KeyOpDeriveKey  KeyOp = "deriveKey"
	KeyOpDeriveBits KeyOp = "deriveBits"
)

// AllValues returns all KeyOp values.
func (KeyOp) AllValues() []KeyOp {
	return []KeyOp{
		KeyOpSign,
		KeyOpVerify,
		KeyOpEncrypt,
		KeyOpDecrypt,
		KeyOpWrapKey,
		KeyOpUnwrapKey,
		KeyOpDeriveKey,
		KeyOpDeriveBits,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s KeyOp) MarshalText() ([]byte, error) {
	switch s {
	case KeyOpSign:
		return []byte(s), nil
	case KeyOpVerify:
		return []byte(s), nil
	case KeyOpEncrypt:
		return []byte(s), nil
	case KeyOpDecrypt:
		return []byte(s), nil
	case KeyOpWrapKey:
		return []byte(s), nil
	case KeyOpUnwrapKey:
		return []byte(s), nil
	case KeyOpDeriveKey:
		return []byte(s), nil
	case KeyOpDeriveBits:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *KeyOp) UnmarshalText(data []byte) error {
	switch KeyOp(data) {
	case KeyOpSign:
		*s = KeyOpSign
		return nil
	case KeyOpVerify:
		*s = KeyOpVerify
		return nil
	case KeyOpEncrypt:
		*s = KeyOpEncrypt
		return nil
	case KeyOpDecrypt:
		*s = KeyOpDecrypt
		return nil
	case KeyOpWrapKey:
		*s = KeyOpWrapKey
		return nil
	case KeyOpUnwrapKey:
		*s = KeyOpUnwrapKey
		return nil
	case KeyOpDeriveKey:
		*s = KeyOpDeriveKey
		return nil
	case KeyOpDeriveBits:
		*s = KeyOpDeriveBits
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// The intended usage of the key.
// Ref: #/components/schemas/KeyUsage
type KeyUsage string

const (
	KeyUsageAuth    KeyUsage = "auth"
	KeyUsageRefresh KeyUsage = "refresh"
)

// AllValues returns all KeyUsage values.
func (KeyUsage) AllValues() []KeyUsage {
	return []KeyUsage{
		KeyUsageAuth,
		KeyUsageRefresh,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s KeyUsage) MarshalText() ([]byte, error) {
	switch s {
	case KeyUsageAuth:
		return []byte(s), nil
	case KeyUsageRefresh:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *KeyUsage) UnmarshalText(data []byte) error {
	switch KeyUsage(data) {
	case KeyUsageAuth:
		*s = KeyUsageAuth
		return nil
	case KeyUsageRefresh:
		*s = KeyUsageRefresh
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// ListPublicKeysIMATeapot is response for ListPublicKeys operation.
type ListPublicKeysIMATeapot struct{}

func (*ListPublicKeysIMATeapot) listPublicKeysRes() {}

type ListPublicKeysOKApplicationJSON []JWK

func (*ListPublicKeysOKApplicationJSON) listPublicKeysRes() {}

// ListUsersIMATeapot is response for ListUsers operation.
type ListUsersIMATeapot struct{}

func (*ListUsersIMATeapot) listUsersRes() {}

type ListUsersOKApplicationJSON []User

func (*ListUsersOKApplicationJSON) listUsersRes() {}

// Data used to authenticate a user. It usually includes some private information only known to the
// user
// (password, secret question, etc), that is checked against some protected data on the server. If
// this
// information is correct, the user is authenticated and granted a special token.
// Ref: #/components/schemas/LoginForm
type LoginForm struct {
	Email    Email    `json:"email"`
	Password Password `json:"password"`
}

// GetEmail returns the value of Email.
func (s *LoginForm) GetEmail() Email {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *LoginForm) GetPassword() Password {
	return s.Password
}

// SetEmail sets the value of Email.
func (s *LoginForm) SetEmail(val Email) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *LoginForm) SetPassword(val Password) {
	s.Password = val
}

// Ref: #/components/schemas/NotFoundError
type NotFoundError struct {
	// The error message.
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *NotFoundError) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *NotFoundError) SetError(val string) {
	s.Error = val
}

func (*NotFoundError) createSessionRes()        {}
func (*NotFoundError) emailExistsRes()          {}
func (*NotFoundError) getPublicKeyRes()         {}
func (*NotFoundError) requestPasswordResetRes() {}
func (*NotFoundError) updateEmailRes()          {}
func (*NotFoundError) updateRoleRes()           {}

// NewOptInt returns new OptInt with value set to v.
func NewOptInt(v int) OptInt {
	return OptInt{
		Value: v,
		Set:   true,
	}
}

// OptInt is optional int.
type OptInt struct {
	Value int
	Set   bool
}

// IsSet returns true if OptInt was set.
func (o OptInt) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptInt) Reset() {
	var v int
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptInt) SetTo(v int) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptInt) Get() (v int, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptInt) Or(d int) int {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptKID returns new OptKID with value set to v.
func NewOptKID(v KID) OptKID {
	return OptKID{
		Value: v,
		Set:   true,
	}
}

// OptKID is optional KID.
type OptKID struct {
	Value KID
	Set   bool
}

// IsSet returns true if OptKID was set.
func (o OptKID) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptKID) Reset() {
	var v KID
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptKID) SetTo(v KID) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptKID) Get() (v KID, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptKID) Or(d KID) KID {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptString returns new OptString with value set to v.
func NewOptString(v string) OptString {
	return OptString{
		Value: v,
		Set:   true,
	}
}

// OptString is optional string.
type OptString struct {
	Value string
	Set   bool
}

// IsSet returns true if OptString was set.
func (o OptString) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptString) Reset() {
	var v string
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptString) SetTo(v string) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptString) Get() (v string, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptString) Or(d string) string {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

// NewOptUUID returns new OptUUID with value set to v.
func NewOptUUID(v uuid.UUID) OptUUID {
	return OptUUID{
		Value: v,
		Set:   true,
	}
}

// OptUUID is optional uuid.UUID.
type OptUUID struct {
	Value uuid.UUID
	Set   bool
}

// IsSet returns true if OptUUID was set.
func (o OptUUID) IsSet() bool { return o.Set }

// Reset unsets value.
func (o *OptUUID) Reset() {
	var v uuid.UUID
	o.Value = v
	o.Set = false
}

// SetTo sets value to v.
func (o *OptUUID) SetTo(v uuid.UUID) {
	o.Set = true
	o.Value = v
}

// Get returns value and boolean that denotes whether value was set.
func (o OptUUID) Get() (v uuid.UUID, ok bool) {
	if !o.Set {
		return v, false
	}
	return o.Value, true
}

// Or returns value if set, or given parameter if does not.
func (o OptUUID) Or(d uuid.UUID) uuid.UUID {
	if v, ok := o.Get(); ok {
		return v
	}
	return d
}

type Password string

// PingIMATeapot is response for Ping operation.
type PingIMATeapot struct{}

func (*PingIMATeapot) pingRes() {}

type PingOK struct {
	Data io.Reader
}

// Read reads data from the Data reader.
//
// Kept to satisfy the io.Reader interface.
func (s PingOK) Read(p []byte) (n int, err error) {
	if s.Data == nil {
		return 0, io.EOF
	}
	return s.Data.Read(p)
}

func (*PingOK) pingRes() {}

// Ref: #/components/schemas/RefreshToken
type RefreshToken struct {
	// The token used to refresh the session.
	RefreshToken string `json:"refreshToken"`
}

// GetRefreshToken returns the value of RefreshToken.
func (s *RefreshToken) GetRefreshToken() string {
	return s.RefreshToken
}

// SetRefreshToken sets the value of RefreshToken.
func (s *RefreshToken) SetRefreshToken(val string) {
	s.RefreshToken = val
}

func (*RefreshToken) createRefreshTokenRes() {}

// Data used to create a user.
// Ref: #/components/schemas/RegisterForm
type RegisterForm struct {
	// The email of the new user. This email must be available, and also match the one that received
	// the short code / registration link.
	Email     Email     `json:"email"`
	Password  Password  `json:"password"`
	ShortCode ShortCode `json:"shortCode"`
}

// GetEmail returns the value of Email.
func (s *RegisterForm) GetEmail() Email {
	return s.Email
}

// GetPassword returns the value of Password.
func (s *RegisterForm) GetPassword() Password {
	return s.Password
}

// GetShortCode returns the value of ShortCode.
func (s *RegisterForm) GetShortCode() ShortCode {
	return s.ShortCode
}

// SetEmail sets the value of Email.
func (s *RegisterForm) SetEmail(val Email) {
	s.Email = val
}

// SetPassword sets the value of Password.
func (s *RegisterForm) SetPassword(val Password) {
	s.Password = val
}

// SetShortCode sets the value of ShortCode.
func (s *RegisterForm) SetShortCode(val ShortCode) {
	s.ShortCode = val
}

// Create a new email update link.
// Ref: #/components/schemas/RequestEmailUpdateForm
type RequestEmailUpdateForm struct {
	// The bew email of the new user. This email must be available at the time of validation.
	Email Email `json:"email"`
}

// GetEmail returns the value of Email.
func (s *RequestEmailUpdateForm) GetEmail() Email {
	return s.Email
}

// SetEmail sets the value of Email.
func (s *RequestEmailUpdateForm) SetEmail(val Email) {
	s.Email = val
}

// RequestEmailUpdateIMATeapot is response for RequestEmailUpdate operation.
type RequestEmailUpdateIMATeapot struct{}

func (*RequestEmailUpdateIMATeapot) requestEmailUpdateRes() {}

// RequestEmailUpdateNoContent is response for RequestEmailUpdate operation.
type RequestEmailUpdateNoContent struct{}

func (*RequestEmailUpdateNoContent) requestEmailUpdateRes() {}

// Create a new password update link.
// Ref: #/components/schemas/RequestPasswordResetForm
type RequestPasswordResetForm struct {
	// The email of the user. This email must match a user in the database.
	Email Email `json:"email"`
}

// GetEmail returns the value of Email.
func (s *RequestPasswordResetForm) GetEmail() Email {
	return s.Email
}

// SetEmail sets the value of Email.
func (s *RequestPasswordResetForm) SetEmail(val Email) {
	s.Email = val
}

// RequestPasswordResetNoContent is response for RequestPasswordReset operation.
type RequestPasswordResetNoContent struct{}

func (*RequestPasswordResetNoContent) requestPasswordResetRes() {}

// Create a new registration link.
// Ref: #/components/schemas/RequestRegistrationForm
type RequestRegistrationForm struct {
	// The email of the new user. This email must be available at the time of registration.
	Email Email `json:"email"`
}

// GetEmail returns the value of Email.
func (s *RequestRegistrationForm) GetEmail() Email {
	return s.Email
}

// SetEmail sets the value of Email.
func (s *RequestRegistrationForm) SetEmail(val Email) {
	s.Email = val
}

// RequestRegistrationIMATeapot is response for RequestRegistration operation.
type RequestRegistrationIMATeapot struct{}

func (*RequestRegistrationIMATeapot) requestRegistrationRes() {}

// RequestRegistrationNoContent is response for RequestRegistration operation.
type RequestRegistrationNoContent struct{}

func (*RequestRegistrationNoContent) requestRegistrationRes() {}

// Data used to reset the password of a user.
// Ref: #/components/schemas/ResetPasswordForm
type ResetPasswordForm struct {
	// The ID of the user that requested a password reset is sent in the reset link of the reset password
	// email.
	UserID UserID `json:"userID"`
	// The new password of the user.
	Password  Password  `json:"password"`
	ShortCode ShortCode `json:"shortCode"`
}

// GetUserID returns the value of UserID.
func (s *ResetPasswordForm) GetUserID() UserID {
	return s.UserID
}

// GetPassword returns the value of Password.
func (s *ResetPasswordForm) GetPassword() Password {
	return s.Password
}

// GetShortCode returns the value of ShortCode.
func (s *ResetPasswordForm) GetShortCode() ShortCode {
	return s.ShortCode
}

// SetUserID sets the value of UserID.
func (s *ResetPasswordForm) SetUserID(val UserID) {
	s.UserID = val
}

// SetPassword sets the value of Password.
func (s *ResetPasswordForm) SetPassword(val Password) {
	s.Password = val
}

// SetShortCode sets the value of ShortCode.
func (s *ResetPasswordForm) SetShortCode(val ShortCode) {
	s.ShortCode = val
}

// ResetPasswordNoContent is response for ResetPassword operation.
type ResetPasswordNoContent struct{}

func (*ResetPasswordNoContent) resetPasswordRes() {}

type Role string

type ShortCode string

// Ref: #/components/schemas/Token
type Token struct {
	// The token used to authenticate the session. This token can be passed as a header to http requests
	// on
	// protected routes.
	AccessToken string `json:"accessToken"`
}

// GetAccessToken returns the value of AccessToken.
func (s *Token) GetAccessToken() string {
	return s.AccessToken
}

// SetAccessToken sets the value of AccessToken.
func (s *Token) SetAccessToken(val string) {
	s.AccessToken = val
}

func (*Token) createAnonSessionRes() {}
func (*Token) createSessionRes()     {}
func (*Token) refreshSessionRes()    {}
func (*Token) registerRes()          {}

// Ref: #/components/schemas/UnexpectedError
type UnexpectedError struct {
	// The error message.
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *UnexpectedError) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *UnexpectedError) SetError(val string) {
	s.Error = val
}

// UnexpectedErrorStatusCode wraps UnexpectedError with StatusCode.
type UnexpectedErrorStatusCode struct {
	StatusCode int
	Response   UnexpectedError
}

// GetStatusCode returns the value of StatusCode.
func (s *UnexpectedErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *UnexpectedErrorStatusCode) GetResponse() UnexpectedError {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *UnexpectedErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *UnexpectedErrorStatusCode) SetResponse(val UnexpectedError) {
	s.Response = val
}

// Ref: #/components/schemas/UnprocessableEntityError
type UnprocessableEntityError struct {
	// The error message.
	Error string `json:"error"`
}

// GetError returns the value of Error.
func (s *UnprocessableEntityError) GetError() string {
	return s.Error
}

// SetError sets the value of Error.
func (s *UnprocessableEntityError) SetError(val string) {
	s.Error = val
}

func (*UnprocessableEntityError) refreshSessionRes() {}
func (*UnprocessableEntityError) updateRoleRes()     {}

// Data used to update the email of a user.
// Ref: #/components/schemas/UpdateEmailForm
type UpdateEmailForm struct {
	// The id of the user that requested the email update. This ID is usually sent along the short code
	// in the email update link.
	UserID    UserID    `json:"userID"`
	ShortCode ShortCode `json:"shortCode"`
}

// GetUserID returns the value of UserID.
func (s *UpdateEmailForm) GetUserID() UserID {
	return s.UserID
}

// GetShortCode returns the value of ShortCode.
func (s *UpdateEmailForm) GetShortCode() ShortCode {
	return s.ShortCode
}

// SetUserID sets the value of UserID.
func (s *UpdateEmailForm) SetUserID(val UserID) {
	s.UserID = val
}

// SetShortCode sets the value of ShortCode.
func (s *UpdateEmailForm) SetShortCode(val ShortCode) {
	s.ShortCode = val
}

type UpdateEmailOK struct {
	// The new email of the user.
	Email Email `json:"email"`
}

// GetEmail returns the value of Email.
func (s *UpdateEmailOK) GetEmail() Email {
	return s.Email
}

// SetEmail sets the value of Email.
func (s *UpdateEmailOK) SetEmail(val Email) {
	s.Email = val
}

func (*UpdateEmailOK) updateEmailRes() {}

// Data used to update the password of a user.
// Ref: #/components/schemas/UpdatePasswordForm
type UpdatePasswordForm struct {
	// The new password of the user.
	Password Password `json:"password"`
	// The current password of the user, used for further verification of the caller identity.
	CurrentPassword Password `json:"currentPassword"`
}

// GetPassword returns the value of Password.
func (s *UpdatePasswordForm) GetPassword() Password {
	return s.Password
}

// GetCurrentPassword returns the value of CurrentPassword.
func (s *UpdatePasswordForm) GetCurrentPassword() Password {
	return s.CurrentPassword
}

// SetPassword sets the value of Password.
func (s *UpdatePasswordForm) SetPassword(val Password) {
	s.Password = val
}

// SetCurrentPassword sets the value of CurrentPassword.
func (s *UpdatePasswordForm) SetCurrentPassword(val Password) {
	s.CurrentPassword = val
}

// UpdatePasswordNoContent is response for UpdatePassword operation.
type UpdatePasswordNoContent struct{}

func (*UpdatePasswordNoContent) updatePasswordRes() {}

// Data used to update the role of a user. The user requesting the update must follow some
// specific rules.
// - A user cannot upgrade other users to a role higher than its own.
// - A user can only downgrade other users to a role lower than its own.
// For example, the following operations are permitted:
// - ✅ A (super_admin) upgrades B (admin) to super_admin.
// - ✅ A (admin) upgrades B (user) to admin.
// - ✅ A (super_admin) downgrades B (admin) to user.
// But the following operations are not:
// - ❌ A (admin) upgrades B (user) to super_admin.
// - ❌ A (admin) downgrades B (admin) to user.
// Ref: #/components/schemas/UpdateRoleForm
type UpdateRoleForm struct {
	// The id of the user who's role is to be updated.
	UserID UserID `json:"userID"`
	// The new role of the user.
	Role CredentialsRole `json:"role"`
}

// GetUserID returns the value of UserID.
func (s *UpdateRoleForm) GetUserID() UserID {
	return s.UserID
}

// GetRole returns the value of Role.
func (s *UpdateRoleForm) GetRole() CredentialsRole {
	return s.Role
}

// SetUserID sets the value of UserID.
func (s *UpdateRoleForm) SetUserID(val UserID) {
	s.UserID = val
}

// SetRole sets the value of Role.
func (s *UpdateRoleForm) SetRole(val CredentialsRole) {
	s.Role = val
}

// The intended use of the public key.
// Ref: #/components/schemas/Use
type Use string

const (
	UseSig Use = "sig"
	UseEnc Use = "enc"
)

// AllValues returns all Use values.
func (Use) AllValues() []Use {
	return []Use{
		UseSig,
		UseEnc,
	}
}

// MarshalText implements encoding.TextMarshaler.
func (s Use) MarshalText() ([]byte, error) {
	switch s {
	case UseSig:
		return []byte(s), nil
	case UseEnc:
		return []byte(s), nil
	default:
		return nil, errors.Errorf("invalid value: %q", s)
	}
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (s *Use) UnmarshalText(data []byte) error {
	switch Use(data) {
	case UseSig:
		*s = UseSig
		return nil
	case UseEnc:
		*s = UseEnc
		return nil
	default:
		return errors.Errorf("invalid value: %q", data)
	}
}

// Ref: #/components/schemas/User
type User struct {
	// The unique identifier of the user.
	ID UserID `json:"id"`
	// The email of the user.
	Email Email `json:"email"`
	// The role of the user.
	Role CredentialsRole `json:"role"`
	// The date and time the user was created.
	CreatedAt time.Time `json:"createdAt"`
	// The date and time the user was last updated.
	UpdatedAt time.Time `json:"updatedAt"`
}

// GetID returns the value of ID.
func (s *User) GetID() UserID {
	return s.ID
}

// GetEmail returns the value of Email.
func (s *User) GetEmail() Email {
	return s.Email
}

// GetRole returns the value of Role.
func (s *User) GetRole() CredentialsRole {
	return s.Role
}

// GetCreatedAt returns the value of CreatedAt.
func (s *User) GetCreatedAt() time.Time {
	return s.CreatedAt
}

// GetUpdatedAt returns the value of UpdatedAt.
func (s *User) GetUpdatedAt() time.Time {
	return s.UpdatedAt
}

// SetID sets the value of ID.
func (s *User) SetID(val UserID) {
	s.ID = val
}

// SetEmail sets the value of Email.
func (s *User) SetEmail(val Email) {
	s.Email = val
}

// SetRole sets the value of Role.
func (s *User) SetRole(val CredentialsRole) {
	s.Role = val
}

// SetCreatedAt sets the value of CreatedAt.
func (s *User) SetCreatedAt(val time.Time) {
	s.CreatedAt = val
}

// SetUpdatedAt sets the value of UpdatedAt.
func (s *User) SetUpdatedAt(val time.Time) {
	s.UpdatedAt = val
}

func (*User) updateRoleRes() {}

type UserID uuid.UUID
