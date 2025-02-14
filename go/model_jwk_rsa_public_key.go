/*
 * OIDC ICT Userinfo Endpoint
 *
 * Endpoint for OpenID Connect's ID Certifcation Token endpoint for userinfo.
 *
 * API version: 0.4.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package ict

import (
	"errors"
)

type JwkRsaPublicKey struct {
	KeyType  KeyType `json:"kty"`
	Modulus  string  `json:"n"`
	Exponent string  `json:"e"`
}

func RsaJwkFromJson(json map[string]interface{}) (JwkRsaPublicKey, error) {
	// Parse key type
	keyType, err := KeyTypeFromJson(json, "kty")
	if err != nil {
		return JwkRsaPublicKey{}, errors.New("failed to parse key type: " + err.Error())
	}
	if keyType != RSA {
		return JwkRsaPublicKey{}, errors.New("failed to parse key type: expected attribute 'kty' to be 'RSA' but found '" + string(keyType) + "'")
	}

	// Parse exponent
	exponent, err := StringFromJson(json, "e")
	if err != nil {
		return JwkRsaPublicKey{}, errors.New("failed to parse exponent: " + err.Error())
	}

	// Parse modulus
	modulus, err := StringFromJson(json, "n")
	if err != nil {
		return JwkRsaPublicKey{}, errors.New("failed to parse modulus: " + err.Error())
	}

	// Return as struct
	return JwkRsaPublicKey{
		KeyType:  keyType,
		Modulus:  modulus,
		Exponent: exponent,
	}, nil
}
