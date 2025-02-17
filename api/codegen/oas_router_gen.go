// Code generated by ogen, DO NOT EDIT.

package codegen

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/ogen-go/ogen/uri"
)

func (s *Server) cutPrefix(path string) (string, bool) {
	prefix := s.cfg.Prefix
	if prefix == "" {
		return path, true
	}
	if !strings.HasPrefix(path, prefix) {
		// Prefix doesn't match.
		return "", false
	}
	// Cut prefix from the path.
	return strings.TrimPrefix(path, prefix), true
}

// ServeHTTP serves http request as defined by OpenAPI v3 specification,
// calling handler that matches the path or returning not found error.
func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	elem := r.URL.Path
	elemIsEscaped := false
	if rawPath := r.URL.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
			elemIsEscaped = strings.ContainsRune(elem, '%')
		}
	}

	elem, ok := s.cutPrefix(elem)
	if !ok || len(elem) == 0 {
		s.notFound(w, r)
		return
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "credentials"
				origElem := elem
				if l := len("credentials"); len(elem) >= l && elem[0:l] == "credentials" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch r.Method {
					case "PUT":
						s.handleRegisterRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "PUT")
					}

					return
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'e': // Prefix: "email"
						origElem := elem
						if l := len("email"); len(elem) >= l && elem[0:l] == "email" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleEmailExistsRequest([0]string{}, elemIsEscaped, w, r)
							case "PATCH":
								s.handleUpdateEmailRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET,PATCH")
							}

							return
						}

						elem = origElem
					case 'p': // Prefix: "password"
						origElem := elem
						if l := len("password"); len(elem) >= l && elem[0:l] == "password" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch r.Method {
							case "PATCH":
								s.handleUpdatePasswordRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "PATCH")
							}

							return
						}
						switch elem[0] {
						case '/': // Prefix: "/reset"
							origElem := elem
							if l := len("/reset"); len(elem) >= l && elem[0:l] == "/reset" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "PATCH":
									s.handleResetPasswordRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "PATCH")
								}

								return
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'h': // Prefix: "healthcheck"
				origElem := elem
				if l := len("healthcheck"); len(elem) >= l && elem[0:l] == "healthcheck" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch r.Method {
					case "GET":
						s.handleHealthcheckRequest([0]string{}, elemIsEscaped, w, r)
					default:
						s.notAllowed(w, r, "GET")
					}

					return
				}

				elem = origElem
			case 'p': // Prefix: "p"
				origElem := elem
				if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'i': // Prefix: "ing"
					origElem := elem
					if l := len("ing"); len(elem) >= l && elem[0:l] == "ing" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch r.Method {
						case "GET":
							s.handlePingRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}

					elem = origElem
				case 'u': // Prefix: "ublic-keys"
					origElem := elem
					if l := len("ublic-keys"); len(elem) >= l && elem[0:l] == "ublic-keys" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleGetPublicKeyRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/list"
						origElem := elem
						if l := len("/list"); len(elem) >= l && elem[0:l] == "/list" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "GET":
								s.handleListPublicKeysRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "GET")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 's': // Prefix: "s"
				origElem := elem
				if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'e': // Prefix: "ession"
					origElem := elem
					if l := len("ession"); len(elem) >= l && elem[0:l] == "ession" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch r.Method {
						case "GET":
							s.handleCheckSessionRequest([0]string{}, elemIsEscaped, w, r)
						case "PUT":
							s.handleCreateSessionRequest([0]string{}, elemIsEscaped, w, r)
						default:
							s.notAllowed(w, r, "GET,PUT")
						}

						return
					}
					switch elem[0] {
					case '/': // Prefix: "/anon"
						origElem := elem
						if l := len("/anon"); len(elem) >= l && elem[0:l] == "/anon" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "PUT":
								s.handleCreateAnonSessionRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "PUT")
							}

							return
						}

						elem = origElem
					}

					elem = origElem
				case 'h': // Prefix: "hort-code/"
					origElem := elem
					if l := len("hort-code/"); len(elem) >= l && elem[0:l] == "hort-code/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'r': // Prefix: "register"
						origElem := elem
						if l := len("register"); len(elem) >= l && elem[0:l] == "register" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch r.Method {
							case "PUT":
								s.handleRequestRegistrationRequest([0]string{}, elemIsEscaped, w, r)
							default:
								s.notAllowed(w, r, "PUT")
							}

							return
						}

						elem = origElem
					case 'u': // Prefix: "update-"
						origElem := elem
						if l := len("update-"); len(elem) >= l && elem[0:l] == "update-" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'e': // Prefix: "email"
							origElem := elem
							if l := len("email"); len(elem) >= l && elem[0:l] == "email" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "PUT":
									s.handleRequestEmailUpdateRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "PUT")
								}

								return
							}

							elem = origElem
						case 'p': // Prefix: "password"
							origElem := elem
							if l := len("password"); len(elem) >= l && elem[0:l] == "password" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch r.Method {
								case "PUT":
									s.handleRequestPasswordResetRequest([0]string{}, elemIsEscaped, w, r)
								default:
									s.notAllowed(w, r, "PUT")
								}

								return
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	s.notFound(w, r)
}

// Route is route object.
type Route struct {
	name        string
	summary     string
	operationID string
	pathPattern string
	count       int
	args        [0]string
}

// Name returns ogen operation name.
//
// It is guaranteed to be unique and not empty.
func (r Route) Name() string {
	return r.name
}

// Summary returns OpenAPI summary.
func (r Route) Summary() string {
	return r.summary
}

// OperationID returns OpenAPI operationId.
func (r Route) OperationID() string {
	return r.operationID
}

// PathPattern returns OpenAPI path.
func (r Route) PathPattern() string {
	return r.pathPattern
}

// Args returns parsed arguments.
func (r Route) Args() []string {
	return r.args[:r.count]
}

// FindRoute finds Route for given method and path.
//
// Note: this method does not unescape path or handle reserved characters in path properly. Use FindPath instead.
func (s *Server) FindRoute(method, path string) (Route, bool) {
	return s.FindPath(method, &url.URL{Path: path})
}

// FindPath finds Route for given method and URL.
func (s *Server) FindPath(method string, u *url.URL) (r Route, _ bool) {
	var (
		elem = u.Path
		args = r.args
	)
	if rawPath := u.RawPath; rawPath != "" {
		if normalized, ok := uri.NormalizeEscapedPath(rawPath); ok {
			elem = normalized
		}
		defer func() {
			for i, arg := range r.args[:r.count] {
				if unescaped, err := url.PathUnescape(arg); err == nil {
					r.args[i] = unescaped
				}
			}
		}()
	}

	elem, ok := s.cutPrefix(elem)
	if !ok {
		return r, false
	}

	// Static code generated router with unwrapped path search.
	switch {
	default:
		if len(elem) == 0 {
			break
		}
		switch elem[0] {
		case '/': // Prefix: "/"
			origElem := elem
			if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
				elem = elem[l:]
			} else {
				break
			}

			if len(elem) == 0 {
				break
			}
			switch elem[0] {
			case 'c': // Prefix: "credentials"
				origElem := elem
				if l := len("credentials"); len(elem) >= l && elem[0:l] == "credentials" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					switch method {
					case "PUT":
						r.name = RegisterOperation
						r.summary = "Create a new user."
						r.operationID = "register"
						r.pathPattern = "/credentials"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}
				switch elem[0] {
				case '/': // Prefix: "/"
					origElem := elem
					if l := len("/"); len(elem) >= l && elem[0:l] == "/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'e': // Prefix: "email"
						origElem := elem
						if l := len("email"); len(elem) >= l && elem[0:l] == "email" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = EmailExistsOperation
								r.summary = "Check the availability of an email."
								r.operationID = "emailExists"
								r.pathPattern = "/credentials/email"
								r.args = args
								r.count = 0
								return r, true
							case "PATCH":
								r.name = UpdateEmailOperation
								r.summary = "Update the email of an user."
								r.operationID = "updateEmail"
								r.pathPattern = "/credentials/email"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 'p': // Prefix: "password"
						origElem := elem
						if l := len("password"); len(elem) >= l && elem[0:l] == "password" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							switch method {
							case "PATCH":
								r.name = UpdatePasswordOperation
								r.summary = "Update the password of an user."
								r.operationID = "updatePassword"
								r.pathPattern = "/credentials/password"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}
						switch elem[0] {
						case '/': // Prefix: "/reset"
							origElem := elem
							if l := len("/reset"); len(elem) >= l && elem[0:l] == "/reset" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "PATCH":
									r.name = ResetPasswordOperation
									r.summary = "Reset the password of an user."
									r.operationID = "resetPassword"
									r.pathPattern = "/credentials/password/reset"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 'h': // Prefix: "healthcheck"
				origElem := elem
				if l := len("healthcheck"); len(elem) >= l && elem[0:l] == "healthcheck" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					// Leaf node.
					switch method {
					case "GET":
						r.name = HealthcheckOperation
						r.summary = "Check the health of the service."
						r.operationID = "healthcheck"
						r.pathPattern = "/healthcheck"
						r.args = args
						r.count = 0
						return r, true
					default:
						return
					}
				}

				elem = origElem
			case 'p': // Prefix: "p"
				origElem := elem
				if l := len("p"); len(elem) >= l && elem[0:l] == "p" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'i': // Prefix: "ing"
					origElem := elem
					if l := len("ing"); len(elem) >= l && elem[0:l] == "ing" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						// Leaf node.
						switch method {
						case "GET":
							r.name = PingOperation
							r.summary = "Check the status of the service."
							r.operationID = "ping"
							r.pathPattern = "/ping"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}

					elem = origElem
				case 'u': // Prefix: "ublic-keys"
					origElem := elem
					if l := len("ublic-keys"); len(elem) >= l && elem[0:l] == "ublic-keys" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = GetPublicKeyOperation
							r.summary = "Get the public keys used for JSON Web Algorithms."
							r.operationID = "getPublicKey"
							r.pathPattern = "/public-keys"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/list"
						origElem := elem
						if l := len("/list"); len(elem) >= l && elem[0:l] == "/list" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "GET":
								r.name = ListPublicKeysOperation
								r.summary = "List all public keys used for JSON Web Algorithms."
								r.operationID = "listPublicKeys"
								r.pathPattern = "/public-keys/list"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			case 's': // Prefix: "s"
				origElem := elem
				if l := len("s"); len(elem) >= l && elem[0:l] == "s" {
					elem = elem[l:]
				} else {
					break
				}

				if len(elem) == 0 {
					break
				}
				switch elem[0] {
				case 'e': // Prefix: "ession"
					origElem := elem
					if l := len("ession"); len(elem) >= l && elem[0:l] == "ession" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						switch method {
						case "GET":
							r.name = CheckSessionOperation
							r.summary = "Check the status of a session."
							r.operationID = "checkSession"
							r.pathPattern = "/session"
							r.args = args
							r.count = 0
							return r, true
						case "PUT":
							r.name = CreateSessionOperation
							r.summary = "Create a new session."
							r.operationID = "createSession"
							r.pathPattern = "/session"
							r.args = args
							r.count = 0
							return r, true
						default:
							return
						}
					}
					switch elem[0] {
					case '/': // Prefix: "/anon"
						origElem := elem
						if l := len("/anon"); len(elem) >= l && elem[0:l] == "/anon" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "PUT":
								r.name = CreateAnonSessionOperation
								r.summary = "Create a new anonymous session."
								r.operationID = "createAnonSession"
								r.pathPattern = "/session/anon"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					}

					elem = origElem
				case 'h': // Prefix: "hort-code/"
					origElem := elem
					if l := len("hort-code/"); len(elem) >= l && elem[0:l] == "hort-code/" {
						elem = elem[l:]
					} else {
						break
					}

					if len(elem) == 0 {
						break
					}
					switch elem[0] {
					case 'r': // Prefix: "register"
						origElem := elem
						if l := len("register"); len(elem) >= l && elem[0:l] == "register" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							// Leaf node.
							switch method {
							case "PUT":
								r.name = RequestRegistrationOperation
								r.summary = "Set a new short code for user registration."
								r.operationID = "requestRegistration"
								r.pathPattern = "/short-code/register"
								r.args = args
								r.count = 0
								return r, true
							default:
								return
							}
						}

						elem = origElem
					case 'u': // Prefix: "update-"
						origElem := elem
						if l := len("update-"); len(elem) >= l && elem[0:l] == "update-" {
							elem = elem[l:]
						} else {
							break
						}

						if len(elem) == 0 {
							break
						}
						switch elem[0] {
						case 'e': // Prefix: "email"
							origElem := elem
							if l := len("email"); len(elem) >= l && elem[0:l] == "email" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "PUT":
									r.name = RequestEmailUpdateOperation
									r.summary = "Set a new short code for user email change."
									r.operationID = "requestEmailUpdate"
									r.pathPattern = "/short-code/update-email"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						case 'p': // Prefix: "password"
							origElem := elem
							if l := len("password"); len(elem) >= l && elem[0:l] == "password" {
								elem = elem[l:]
							} else {
								break
							}

							if len(elem) == 0 {
								// Leaf node.
								switch method {
								case "PUT":
									r.name = RequestPasswordResetOperation
									r.summary = "Set a new short code for user password change."
									r.operationID = "requestPasswordReset"
									r.pathPattern = "/short-code/update-password"
									r.args = args
									r.count = 0
									return r, true
								default:
									return
								}
							}

							elem = origElem
						}

						elem = origElem
					}

					elem = origElem
				}

				elem = origElem
			}

			elem = origElem
		}
	}
	return r, false
}
