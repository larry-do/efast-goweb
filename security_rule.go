package goweb

import "github.com/pkg/errors"

type SecurityRuleChecking func(Request) error

const AuthorizationAcceptAnyRole = "ANY_ROLE"

var (
	ErrorForbidden = errors.New("I know you, man. But you can not do this here.")
	ErrorUnauthorized = errors.New("I don't know who you are, man.")
)
