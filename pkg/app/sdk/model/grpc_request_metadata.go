package model

import (
	"context"
)

type AuthToken struct {
	projectID        string
	projectKey       string
	chainAccountAddr string
	enableTLS        bool
	domain           string
	chainType        string
}

func NewAuthToken(projectID, projectKey, chainAccountAddr string) AuthToken {
	return AuthToken{projectID, projectKey, chainAccountAddr, true, "bsngate.com", "irisnet"}
}

func (a *AuthToken) GetRequestMetadata(context.Context, ...string) (
	map[string]string, error,
) {
	return map[string]string{"projectIdHeader": a.projectID, "projectKeyHeader": a.projectKey, "chainAccountAddressHeader": a.chainAccountAddr, "x-api-chain-type": a.chainType}, nil
}

func (a *AuthToken) RequireTransportSecurity() bool {
	return a.enableTLS
}

func (a *AuthToken) SetRequireTransportSecurity(enabled bool) {
	a.enableTLS = enabled
}

func (a *AuthToken) GetProjectKey() string {
	return a.projectKey
}

func (a *AuthToken) GetProjectID() string {
	return a.projectID
}

func (a *AuthToken) GetChainAccountAddr() string {
	return a.chainAccountAddr
}

func (a *AuthToken) GetEnableTLS() bool {
	return a.enableTLS
}

func (a *AuthToken) SetDomain(domain string) {
	a.domain = domain
}

func (a *AuthToken) GetDomain() string {
	return a.domain
}

func (a *AuthToken) SetChainType(chainType string) {
	a.chainType = chainType
}

func (a *AuthToken) GetChainType() string {
	return a.chainType
}
