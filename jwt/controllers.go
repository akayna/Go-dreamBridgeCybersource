package jwt

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/akayna/Go-dreamBridgeUtils/timeutils"
	jwt "github.com/dgrijalva/jwt-go"
)

// Constants

// DefaultExpirationTTLMilis - Default JWT token TTL in miliseconds
const DefaultExpirationTTLMilis int64 = 3600000

// Functions

// GetJWTString - Return the JWT token as string
func (jwtClaims Claims) GetJWTString() string {

	claims := jwtCyberSourceCustomClains{
		jwt.StandardClaims{
			IssuedAt: time.Now().Unix(),
			//ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(30)).Unix(),
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(4)).Unix(),
			Issuer:    jwtClaims.CardinalCredentials.APIIdentifier,
			Id:        timeutils.GetTimeID(),
		},
		jwtClaims.CardinalCredentials.OrgUnitID,
		jwtClaims.Payload,
		jwtClaims.ReferenceID,
		"true",
	}

	fmt.Println("JWT Clains:")

	b, err := json.Marshal(claims)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fmt.Println(string(b))

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(jwtClaims.CardinalCredentials.APIKeyID))

	if err != nil {
		log.Fatalln(err)
		return ""
	}

	return tokenString
}

// ValidateJWTSignature - Validate the jwt signature
func ValidateJWTSignature(jwtTokenString, APIKeyID string) (bool, *jwt.Token) {

	log.Println("Token:")
	log.Println(jwtTokenString)

	log.Println("Key:")
	log.Println(APIKeyID)

	// Parse the token.  Load the key from command line option
	token, err := jwt.Parse(jwtTokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(APIKeyID), nil
	})

	if err != nil {
		log.Println("cybersourcejwt - validateJWTSignature - Error validating jwt signature. Error:")
		log.Println(err)
		return false, nil
	}

	return true, token
}

// ValidateReadJWT - Validate and read the jwt
func ValidateReadJWT(jwtTokenString, APIKeyID string) jwt.MapClaims {

	log.Println("Token:")
	log.Println(jwtTokenString)

	//log.Println("Key:")
	//log.Println(APIKeyID)

	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(jwtTokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(APIKeyID), nil
	})
	if err != nil {
		log.Println("cybersourcejwt - ValidateReadJWT - Error validating jwt. Error:")
		log.Println(err)
		return nil
	}

	// do something with decoded claims
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}

	return claims
}
