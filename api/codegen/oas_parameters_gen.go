// Code generated by ogen, DO NOT EDIT.

package codegen

import (
	"net/http"

	"github.com/google/uuid"

	"github.com/ogen-go/ogen/conv"
	"github.com/ogen-go/ogen/middleware"
	"github.com/ogen-go/ogen/ogenerrors"
	"github.com/ogen-go/ogen/uri"
)

// EmailExistsParams is parameters of emailExists operation.
type EmailExistsParams struct {
	Email Email
}

func unpackEmailExistsParams(packed middleware.Parameters) (params EmailExistsParams) {
	{
		key := middleware.ParameterKey{
			Name: "email",
			In:   "query",
		}
		params.Email = packed[key].(Email)
	}
	return params
}

func decodeEmailExistsParams(args [0]string, argsEscaped bool, r *http.Request) (params EmailExistsParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: email.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "email",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotEmailVal string
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToString(val)
					if err != nil {
						return err
					}

					paramsDotEmailVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Email = Email(paramsDotEmailVal)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := params.Email.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return err
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "email",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// GetPublicKeyParams is parameters of getPublicKey operation.
type GetPublicKeyParams struct {
	// The unique identifier of the key, conveyed through its KID field.
	Kid KID
}

func unpackGetPublicKeyParams(packed middleware.Parameters) (params GetPublicKeyParams) {
	{
		key := middleware.ParameterKey{
			Name: "kid",
			In:   "query",
		}
		params.Kid = packed[key].(KID)
	}
	return params
}

func decodeGetPublicKeyParams(args [0]string, argsEscaped bool, r *http.Request) (params GetPublicKeyParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: kid.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "kid",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				var paramsDotKidVal uuid.UUID
				if err := func() error {
					val, err := d.DecodeValue()
					if err != nil {
						return err
					}

					c, err := conv.ToUUID(val)
					if err != nil {
						return err
					}

					paramsDotKidVal = c
					return nil
				}(); err != nil {
					return err
				}
				params.Kid = KID(paramsDotKidVal)
				return nil
			}); err != nil {
				return err
			}
		} else {
			return err
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "kid",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// ListPublicKeysParams is parameters of listPublicKeys operation.
type ListPublicKeysParams struct {
	// The intended usage of the keys.
	Usage KeyUsage
}

func unpackListPublicKeysParams(packed middleware.Parameters) (params ListPublicKeysParams) {
	{
		key := middleware.ParameterKey{
			Name: "usage",
			In:   "query",
		}
		params.Usage = packed[key].(KeyUsage)
	}
	return params
}

func decodeListPublicKeysParams(args [0]string, argsEscaped bool, r *http.Request) (params ListPublicKeysParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: usage.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "usage",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.Usage = KeyUsage(c)
				return nil
			}); err != nil {
				return err
			}
			if err := func() error {
				if err := params.Usage.Validate(); err != nil {
					return err
				}
				return nil
			}(); err != nil {
				return err
			}
		} else {
			return err
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "usage",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}

// RefreshSessionParams is parameters of refreshSession operation.
type RefreshSessionParams struct {
	RefreshToken string
	AccessToken  string
}

func unpackRefreshSessionParams(packed middleware.Parameters) (params RefreshSessionParams) {
	{
		key := middleware.ParameterKey{
			Name: "refresh_token",
			In:   "query",
		}
		params.RefreshToken = packed[key].(string)
	}
	{
		key := middleware.ParameterKey{
			Name: "access_token",
			In:   "query",
		}
		params.AccessToken = packed[key].(string)
	}
	return params
}

func decodeRefreshSessionParams(args [0]string, argsEscaped bool, r *http.Request) (params RefreshSessionParams, _ error) {
	q := uri.NewQueryDecoder(r.URL.Query())
	// Decode query: refresh_token.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "refresh_token",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.RefreshToken = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return err
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "refresh_token",
			In:   "query",
			Err:  err,
		}
	}
	// Decode query: access_token.
	if err := func() error {
		cfg := uri.QueryParameterDecodingConfig{
			Name:    "access_token",
			Style:   uri.QueryStyleForm,
			Explode: true,
		}

		if err := q.HasParam(cfg); err == nil {
			if err := q.DecodeParam(cfg, func(d uri.Decoder) error {
				val, err := d.DecodeValue()
				if err != nil {
					return err
				}

				c, err := conv.ToString(val)
				if err != nil {
					return err
				}

				params.AccessToken = c
				return nil
			}); err != nil {
				return err
			}
		} else {
			return err
		}
		return nil
	}(); err != nil {
		return params, &ogenerrors.DecodeParamError{
			Name: "access_token",
			In:   "query",
			Err:  err,
		}
	}
	return params, nil
}
