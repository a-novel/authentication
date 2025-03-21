// Code generated by ogen, DO NOT EDIT.

package codegen

import (
	"fmt"

	"github.com/go-faster/errors"

	"github.com/ogen-go/ogen/validate"
)

func (s Alg) Validate() error {
	switch s {
	case "HS256":
		return nil
	case "HS384":
		return nil
	case "HS512":
		return nil
	case "RS256":
		return nil
	case "RS384":
		return nil
	case "RS512":
		return nil
	case "ES256":
		return nil
	case "ES384":
		return nil
	case "ES512":
		return nil
	case "PS256":
		return nil
	case "PS384":
		return nil
	case "PS512":
		return nil
	case "EdDSA":
		return nil
	case "RSA-OAEP":
		return nil
	case "RSA-OAEP-256":
		return nil
	case "A128KW":
		return nil
	case "A192KW":
		return nil
	case "A256KW":
		return nil
	case "dir":
		return nil
	case "ECDH-ES":
		return nil
	case "ECDH-ES+A128KW":
		return nil
	case "ECDH-ES+A192KW":
		return nil
	case "ECDH-ES+A256KW":
		return nil
	case "A128GCMKW":
		return nil
	case "A192GCMKW":
		return nil
	case "A256GCMKW":
		return nil
	case "PBES2-HS256+A128KW":
		return nil
	case "PBES2-HS384+A192KW":
		return nil
	case "PBES2-HS512+A256KW":
		return nil
	case "A128CBC-HS256":
		return nil
	case "A192CBC-HS384":
		return nil
	case "A256CBC-HS512":
		return nil
	case "A128GCM":
		return nil
	case "A192GCM":
		return nil
	case "A256GCM":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *Claims) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if s.Roles == nil {
			return errors.New("nil is invalid value")
		}
		if err := (validate.Array{
			MinLength:    1,
			MinLengthSet: true,
			MaxLength:    0,
			MaxLengthSet: false,
		}).ValidateLength(len(s.Roles)); err != nil {
			return errors.Wrap(err, "array")
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "roles",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s CredentialsRole) Validate() error {
	switch s {
	case "user":
		return nil
	case "admin":
		return nil
	case "super_admin":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *Dependency) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Status.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "status",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s DependencyStatus) Validate() error {
	switch s {
	case "up":
		return nil
	case "down":
		return nil
	case "unknown":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s Email) Validate() error {
	alias := (string)(s)
	if err := (validate.String{
		MinLength:    3,
		MinLengthSet: true,
		MaxLength:    64,
		MaxLengthSet: true,
		Email:        true,
		Hostname:     false,
		Regex:        nil,
	}).Validate(string(alias)); err != nil {
		return errors.Wrap(err, "string")
	}
	return nil
}

func (s *Health) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Postgres.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "postgres",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Sendgrid.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "sendgrid",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *JWK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Kty.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "kty",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Use.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "use",
			Error: err,
		})
	}
	if err := func() error {
		if s.KeyOps == nil {
			return errors.New("nil is invalid value")
		}
		var failures []validate.FieldError
		for i, elem := range s.KeyOps {
			if err := func() error {
				if err := elem.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				failures = append(failures, validate.FieldError{
					Name:  fmt.Sprintf("[%d]", i),
					Error: err,
				})
			}
		}
		if len(failures) > 0 {
			return &validate.Error{Fields: failures}
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "key_ops",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Alg.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "alg",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s KTY) Validate() error {
	switch s {
	case "oct":
		return nil
	case "RSA":
		return nil
	case "EC":
		return nil
	case "OKP":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s KeyOp) Validate() error {
	switch s {
	case "sign":
		return nil
	case "verify":
		return nil
	case "encrypt":
		return nil
	case "decrypt":
		return nil
	case "wrapKey":
		return nil
	case "unwrapKey":
		return nil
	case "deriveKey":
		return nil
	case "deriveBits":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s KeyUsage) Validate() error {
	switch s {
	case "auth":
		return nil
	case "refresh":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s ListPublicKeysOKApplicationJSON) Validate() error {
	alias := ([]JWK)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range alias {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s ListUsersOKApplicationJSON) Validate() error {
	alias := ([]User)(s)
	if alias == nil {
		return errors.New("nil is invalid value")
	}
	var failures []validate.FieldError
	for i, elem := range alias {
		if err := func() error {
			if err := elem.Validate(); err != nil {
				return err
			}
			return nil
		}(); err != nil {
			failures = append(failures, validate.FieldError{
				Name:  fmt.Sprintf("[%d]", i),
				Error: err,
			})
		}
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *LoginForm) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Email.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Password.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "password",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s Password) Validate() error {
	alias := (string)(s)
	if err := (validate.String{
		MinLength:    2,
		MinLengthSet: true,
		MaxLength:    512,
		MaxLengthSet: true,
		Email:        false,
		Hostname:     false,
		Regex:        nil,
	}).Validate(string(alias)); err != nil {
		return errors.Wrap(err, "string")
	}
	return nil
}

func (s *RegisterForm) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Email.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Password.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "password",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.ShortCode.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "shortCode",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *RequestEmailUpdateForm) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Email.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *RequestPasswordResetForm) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Email.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *RequestRegistrationForm) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Email.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *ResetPasswordForm) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Password.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "password",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.ShortCode.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "shortCode",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s ShortCode) Validate() error {
	alias := (string)(s)
	if err := (validate.String{
		MinLength:    1,
		MinLengthSet: true,
		MaxLength:    32,
		MaxLengthSet: true,
		Email:        false,
		Hostname:     false,
		Regex:        nil,
	}).Validate(string(alias)); err != nil {
		return errors.Wrap(err, "string")
	}
	return nil
}

func (s *UpdateEmailForm) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.ShortCode.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "shortCode",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *UpdateEmailOK) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Email.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *UpdatePasswordForm) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Password.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "password",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.CurrentPassword.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "currentPassword",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s *UpdateRoleForm) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Role.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "role",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}

func (s Use) Validate() error {
	switch s {
	case "sig":
		return nil
	case "enc":
		return nil
	default:
		return errors.Errorf("invalid value: %v", s)
	}
}

func (s *User) Validate() error {
	if s == nil {
		return validate.ErrNilPointer
	}

	var failures []validate.FieldError
	if err := func() error {
		if err := s.Email.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "email",
			Error: err,
		})
	}
	if err := func() error {
		if err := s.Role.Validate(); err != nil {
			return err
		}
		return nil
	}(); err != nil {
		failures = append(failures, validate.FieldError{
			Name:  "role",
			Error: err,
		})
	}
	if len(failures) > 0 {
		return &validate.Error{Fields: failures}
	}
	return nil
}
