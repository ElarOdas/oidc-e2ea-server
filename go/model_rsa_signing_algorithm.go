/*
 * OIDC ICT Userinfo Endpoint
 *
 * Endpoint for OpenID Connect's ID Certifcation Token endpoint for userinfo.
 *
 * API version: 0.4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package ict

type RsaSigningAlgorithm string

// List of RsaSigningAlgorithms
const (
	RS256 RsaSigningAlgorithm = "RS256"
	RS384 RsaSigningAlgorithm = "RS384"
	RS512 RsaSigningAlgorithm = "RS512"
)
