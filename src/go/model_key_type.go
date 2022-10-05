/*
 * OIDC RIDT Userinfo Endpoint
 *
 * Endpoint for OpenID Connect's Remote ID Token endpoint for userinfo.
 *
 * API version: 0.2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package ridt

type KeyType string

// List of KeyTypes
const (
	EC  KeyType = "EC"
	RSA KeyType = "RSA"
)

func KeyTypeFromString(value string) (KeyType, bool) {
	switch value {
	case "EC":
		return EC, true
	case "RSA":
		return RSA, true
	default:
		return "", false
	}
}
