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

type EcCurve string

// List of EcCurves
const (
	P256 EcCurve = "P-256"
	P384 EcCurve = "P-384"
	P521 EcCurve = "P-521"
)

func EcCurveFromName(name string) (EcCurve, bool) {
	switch name {
	case "P-256":
		return P256, true
	case "P-384":
		return P384, true
	case "P-521":
		return P521, true
	default:
		return "", false
	}
}

func EcCurveFromJson(json map[string]interface{}, attributeName string) (EcCurve, error) {
	// Read curve name from json object
	crvString, err := StringFromJson(json, attributeName)
	if err != nil {
		return "", err
	}

	// Convert attribute value to EcCurve
	crv, ok := EcCurveFromName(crvString)
	if !ok {
		return crv, errors.New("elliptic curve name '" + string(crv) + "' not supported")
	}

	// Return result
	return crv, nil
}
