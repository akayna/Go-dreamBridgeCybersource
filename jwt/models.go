package jwt

import (
	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
	jwt "github.com/dgrijalva/jwt-go"
)

// Claims - Struct with the JWT Claims information
type Claims struct {
	CardinalCredentials *commons.CardinalCredential
	ReferenceID         string
	Payload             interface{}
}

type jwtCyberSourceCustomClains struct {
	jwt.StandardClaims
	OrgUnitID        string      `json:"OrgUnitId"`
	Payload          interface{} `json:"Payload"`
	ReferenceID      string      `json:"ReferenceId"`
	ObjectifyPayload string      `json:"ObjectifyPayload"`
}
