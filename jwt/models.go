package jwt

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/rafaelcunha/Go-CyberSource/cybersourcecommons"
)

// Claims - Struct with the JWT Claims information
type Claims struct {
	CardinalCredentials *cybersourcecommons.CardinalCredential
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
