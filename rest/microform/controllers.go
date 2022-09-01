package microform

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"

	"github.com/akayna/Go-dreamBridgeCybersource/rest"
	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
	"github.com/akayna/Go-dreamBridgeUtils/jwtutils"
	"github.com/dgrijalva/jwt-go"
)

var generateMicroformContext = "/microform/v2/sessions"
var getJWTPublicKey = "/flex/v2/public-keys"

var lastJwkStr string

// GenerateMicroformContext - Calls the Cybersource API to generate de Microform context
func GenerateMicroformContext(credentials *commons.CyberSourceCredential, targetOrigins []string) (string, string, error) {
	// Create the payload
	var generateKeyData generateContextRequest

	generateKeyData.Targetorigins = targetOrigins

	payload, err := json.Marshal(generateKeyData)
	if err != nil {
		log.Println("microform - GenerateMicroformContext - Error converting struct to json string.")
		return "", "Error converting struct generateKeyData to json.", err
	}

	// Make the POST to generate the Microform Context
	generateContextResp, err := rest.RestFullPOST(credentials, generateMicroformContext, string(payload))

	if err != nil {
		log.Println("microform - GenerateMicroformContext - Error executing POST request.")
		return "", "Error executing POST request.", err
	}

	// Confirms a good HTTP response
	if generateContextResp.StatusCode > 299 || generateContextResp.StatusCode < 200 {

		return "", generateContextResp.Body, nil
	}

	// Validate the capture context
	contextJWT := generateContextResp.Body
	//log.Println("JWT: ", contextJWT)

	token, err := jwt.Parse(contextJWT, getPublicKey)

	if err != nil {
		log.Println("microform - GenerateMicroformContext - Error parsing JWT.")
		return "", "Error parsing JWT.", err
	}

	// Get the JWK and save it to validade the generated token in the frontend
	claims := token.Claims.(jwt.MapClaims)
	ctx := claims["flx"].(map[string]interface{})
	jwk := ctx["jwk"].(map[string]interface{})

	jwkStr, err := json.Marshal(jwk)

	if err != nil {
		log.Println("microform - GenerateMicroformContext - Error parsing JWK.")
		return "", "Error parsing JWK.", err
	}

	lastJwkStr = string(jwkStr)

	return contextJWT, "", nil
}

func getPublicKey(token *jwt.Token) (interface{}, error) {

	kid := token.Header["kid"]

	if kid == nil {
		err := errors.New("error reading context JWT kid")
		log.Println("microform - getPublicKey - Error reading Context JWT kid.")
		return nil, err
	}

	// Get the public key for validation
	getJWTPublicKeyResp, err := rest.RestFullGETNoCerdentials(getJWTPublicKey + "/" + kid.(string))

	if err != nil {
		log.Println("microform - getPublicKey - Error executing GET request to get the JWT public key.")
		return nil, err
	}

	if err != nil {
		log.Println("microform - getPublicKey - Error executing GET request to get the JWT public key.")
		return nil, err
	}

	// Confirms a good HTTP response
	if getJWTPublicKeyResp.StatusCode > 299 || getJWTPublicKeyResp.StatusCode < 200 {
		log.Println("microform - getPublicKey - Request status code: ", getJWTPublicKeyResp.StatusCode)
		err := errors.New("request status code " + strconv.Itoa(getJWTPublicKeyResp.StatusCode))
		return nil, err
	}

	//log.Println("Public Key: ", getJWTPublicKeyResp.Body)

	// Converts the json jwk to rsa key
	var jwk jwtutils.JSONKey

	err = jwk.Populate(getJWTPublicKeyResp.Body)

	if err != nil {
		log.Println("microform - getPublicKey - Error parsing json to jwk.")
		return nil, err
	}

	rsaKey, err := jwk.RSA()

	if err != nil {
		log.Println("microform - getPublicKey - Error converting jwk to rsa key.")
		return nil, err
	}

	return rsaKey, nil
}

// Validate the token received and returns the temporary Cybersource card token.
func ValidateToken(token string) (bool, string, error) {

	parsedToken, err := jwt.Parse(token, getJWK)

	if err != nil {
		log.Println("microform.ValidateToken - Error parsing JWT.")
		return false, "Error parsing JWT.", err
	}

	// Get the JWK and save it to validade the generated token in the frontend
	claims := parsedToken.Claims.(jwt.MapClaims)

	return true, claims["jti"].(string), nil

}

func getJWK(token *jwt.Token) (interface{}, error) {

	// Converts the json jwk to rsa key
	var jwk jwtutils.JSONKey

	err := jwk.Populate(lastJwkStr)

	if err != nil {
		log.Println("microform.getJWK - Error parsing json to jwk.")
		return nil, err
	}

	rsaKey, err := jwk.RSA()

	if err != nil {
		log.Println("microform.getJWK - Error converting jwk to rsa key.")
		return nil, err
	}

	return rsaKey, nil
}
