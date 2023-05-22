package goweb

type SecurityRuleChecking func(Response, Request) bool

const AuthorizationAcceptAnyRole = "ANY_ROLE"
