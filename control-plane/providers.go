package main

import "log"

// ---------------- Workforce IAM ------------------------------------------------
type OktaProvider struct{ orgURL, apiToken string }
func (p *OktaProvider) Name() string     { return "Okta" }
func (p *OktaProvider) Category() string  { return "Workforce IAM" }
func (p *OktaProvider) Revoke(blastRadius string) error {
	// TODO: revoke Okta sessions & tokens
	log.Printf("[Okta] Revoking %s access", blastRadius)
	return nil
}
func NewOktaProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &OktaProvider{}, nil
}

type AzureADProvider struct{}
func (p *AzureADProvider) Name() string    { return "Azure AD" }
func (p *AzureADProvider) Category() string { return "Workforce IAM" }
func (p *AzureADProvider) Revoke(blastRadius string) error {
	// Microsoft Graph: revokeSignInSessions, disable accounts
	return nil
}
func NewAzureADProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &AzureADProvider{}, nil
}

type PingProvider struct{}
func (p *PingProvider) Name() string      { return "Ping Identity" }
func (p *PingProvider) Category() string   { return "Workforce IAM" }
func (p *PingProvider) Revoke(blastRadius string) error { return nil }
func NewPingProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &PingProvider{}, nil
}

type OneLoginProvider struct{}
func (p *OneLoginProvider) Name() string    { return "OneLogin" }
func (p *OneLoginProvider) Category() string { return "Workforce IAM" }
func (p *OneLoginProvider) Revoke(blastRadius string) error { return nil }
func NewOneLoginProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &OneLoginProvider{}, nil
}

// ---------------- CIAM --------------------------------------------------------
type Auth0Provider struct{}
func (p *Auth0Provider) Name() string      { return "Auth0" }
func (p *Auth0Provider) Category() string   { return "CIAM" }
func (p *Auth0Provider) Revoke(blastRadius string) error { return nil }
func NewAuth0Provider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &Auth0Provider{}, nil
}

type CognitoProvider struct{}
func (p *CognitoProvider) Name() string    { return "AWS Cognito" }
func (p *CognitoProvider) Category() string { return "CIAM" }
func (p *CognitoProvider) Revoke(blastRadius string) error { return nil }
func NewCognitoProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &CognitoProvider{}, nil
}

// ---------------- PAM ---------------------------------------------------------
type CyberArkProvider struct{}
func (p *CyberArkProvider) Name() string    { return "CyberArk" }
func (p *CyberArkProvider) Category() string { return "PAM" }
func (p *CyberArkProvider) Revoke(blastRadius string) error { return nil }
func NewCyberArkProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &CyberArkProvider{}, nil
}

type DelineaProvider struct{}
func (p *DelineaProvider) Name() string    { return "Delinea" }
func (p *DelineaProvider) Category() string { return "PAM" }
func (p *DelineaProvider) Revoke(blastRadius string) error { return nil }
func NewDelineaProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &DelineaProvider{}, nil
}

type BeyondTrustProvider struct{}
func (p *BeyondTrustProvider) Name() string    { return "BeyondTrust" }
func (p *BeyondTrustProvider) Category() string { return "PAM" }
func (p *BeyondTrustProvider) Revoke(blastRadius string) error { return nil }
func NewBeyondTrustProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &BeyondTrustProvider{}, nil
}

// ---------------- PIM ---------------------------------------------------------
type AzurePIMProvider struct{}
func (p *AzurePIMProvider) Name() string    { return "Azure AD PIM" }
func (p *AzurePIMProvider) Category() string { return "PIM" }
func (p *AzurePIMProvider) Revoke(blastRadius string) error {
	// Deactivate all eligible role assignments
	return nil
}
func NewAzurePIMProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &AzurePIMProvider{}, nil
}

// ---------------- NHI / Secrets -----------------------------------------------
type VaultProvider struct{}
func (p *VaultProvider) Name() string      { return "HashiCorp Vault" }
func (p *VaultProvider) Category() string   { return "NHI / Secrets" }
func (p *VaultProvider) Revoke(blastRadius string) error {
	// Revoke all leases, rotate secrets
	return nil
}
func NewVaultProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &VaultProvider{}, nil
}

type AWSSecretsProvider struct{}
func (p *AWSSecretsProvider) Name() string    { return "AWS Secrets Manager" }
func (p *AWSSecretsProvider) Category() string { return "NHI / Secrets" }
func (p *AWSSecretsProvider) Revoke(blastRadius string) error { return nil }
func NewAWSSecretsProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &AWSSecretsProvider{}, nil
}

type AkeylessProvider struct{}
func (p *AkeylessProvider) Name() string     { return "Akeyless" }
func (p *AkeylessProvider) Category() string  { return "NHI / Secrets" }
func (p *AkeylessProvider) Revoke(blastRadius string) error { return nil }
func NewAkeylessProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &AkeylessProvider{}, nil
}

// ---------------- API Gateways ------------------------------------------------
type KongProvider struct{}
func (p *KongProvider) Name() string       { return "Kong" }
func (p *KongProvider) Category() string    { return "API Gateway" }
func (p *KongProvider) Revoke(blastRadius string) error {
	// Block all routes / consumers
	return nil
}
func NewKongProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &KongProvider{}, nil
}

type AWSAPIGatewayProvider struct{}
func (p *AWSAPIGatewayProvider) Name() string    { return "AWS API Gateway" }
func (p *AWSAPIGatewayProvider) Category() string { return "API Gateway" }
func (p *AWSAPIGatewayProvider) Revoke(blastRadius string) error { return nil }
func NewAWSAPIGatewayProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &AWSAPIGatewayProvider{}, nil
}

// ---------------- Cloud IAM ---------------------------------------------------
type AWSIAMProvider struct{}
func (p *AWSIAMProvider) Name() string      { return "AWS IAM" }
func (p *AWSIAMProvider) Category() string   { return "Cloud IAM" }
func (p *AWSIAMProvider) Revoke(blastRadius string) error {
	// Attach deny-all SCPs / inline policies
	return nil
}
func NewAWSIAMProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &AWSIAMProvider{}, nil
}

type GCPIAMProvider struct{}
func (p *GCPIAMProvider) Name() string      { return "GCP IAM" }
func (p *GCPIAMProvider) Category() string   { return "Cloud IAM" }
func (p *GCPIAMProvider) Revoke(blastRadius string) error { return nil }
func NewGCPIAMProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &GCPIAMProvider{}, nil
}

// ---------------- Zero Trust --------------------------------------------------
type ZscalerProvider struct{}
func (p *ZscalerProvider) Name() string     { return "Zscaler" }
func (p *ZscalerProvider) Category() string  { return "Zero Trust" }
func (p *ZscalerProvider) Revoke(blastRadius string) error { return nil }
func NewZscalerProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &ZscalerProvider{}, nil
}

type CloudflareZTProvider struct{}
func (p *CloudflareZTProvider) Name() string    { return "Cloudflare Zero Trust" }
func (p *CloudflareZTProvider) Category() string { return "Zero Trust" }
func (p *CloudflareZTProvider) Revoke(blastRadius string) error { return nil }
func NewCloudflareZTProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &CloudflareZTProvider{}, nil
}

// ---------------- Endpoint / EDR ----------------------------------------------
type CrowdStrikeProvider struct{}
func (p *CrowdStrikeProvider) Name() string    { return "CrowdStrike" }
func (p *CrowdStrikeProvider) Category() string { return "Endpoint" }
func (p *CrowdStrikeProvider) Revoke(blastRadius string) error {
	// Network contain hosts
	return nil
}
func NewCrowdStrikeProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &CrowdStrikeProvider{}, nil
}

// ---------------- Session Stores ----------------------------------------------
type RedisSessionProvider struct{}
func (p *RedisSessionProvider) Name() string    { return "Redis Sessions" }
func (p *RedisSessionProvider) Category() string { return "Session Store" }
func (p *RedisSessionProvider) Revoke(blastRadius string) error {
	// FLUSHALL or delete session keys
	return nil
}
func NewRedisSessionProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &RedisSessionProvider{}, nil
}

// ---------------- Service Mesh ------------------------------------------------
type IstioProvider struct{}
func (p *IstioProvider) Name() string       { return "Istio" }
func (p *IstioProvider) Category() string    { return "Service Mesh" }
func (p *IstioProvider) Revoke(blastRadius string) error {
	// Deny-all AuthorizationPolicy
	return nil
}
func NewIstioProvider(cfg map[string]interface{}) (RevocationProvider, error) {
	return &IstioProvider{}, nil
}
