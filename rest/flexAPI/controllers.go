package flexAPI

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/akayna/Go-dreamBridgeCybersource/rest"
	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
)

var generateKeyEndpoint = "/flex/v1/keys"
var generateCardTokenEndpoint = "/flex/v1/tokens"

// GenerateCardTokenFrontend - Example function to generate the card token. This function must be implemented on the frontend
func GenerateCardTokenFrontend(keyID string, cardInfo *CardInfo) (*GenerateCardTokenResponse, string, error) {
	generateCardRequest := generateCardTokenRequest{
		KeyID:    keyID,
		CardInfo: cardInfo,
	}

	payload, err := json.Marshal(generateCardRequest)
	if err != nil {
		log.Println("cybersourceflex - GenerateCardTokenFrontend - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	generateCardTokenResp, err := rest.RestFullSimplePOST(generateCardTokenEndpoint, string(payload))

	if err != nil {
		log.Println("cybersourceflex - GenerateCardTokenFrontend - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	fmt.Printf("Card Token Response: %+v\n", generateCardTokenResp)

	return treateGenerateCardTokenResponse(generateCardTokenResp)
}

// treateGenerateCardTokenResponse - Treate the response to the Generate card token POST request
func treateGenerateCardTokenResponse(response *rest.RequestResponse) (*GenerateCardTokenResponse, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	generateCardTokenResponse := GenerateCardTokenResponse{}
	err := json.Unmarshal([]byte(response.Body), &generateCardTokenResponse)

	if err != nil {
		log.Println("cybersourceflex - treateGenerateCardTokenResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &generateCardTokenResponse, "Key generated.", nil
}

// GenerateRsaOaep256Key - Execute the post to generate Key with RsaOaep256 encryption
func GenerateRsaOaep256Key(credentials *commons.CyberSourceCredential, targetorigin *string) (*GenerateKeyResponse, string, error) {
	var generateKeyData GenerateKeyRequest

	encryptiontype := "RsaOaep256"

	generateKeyData.Encryptiontype = &encryptiontype
	generateKeyData.Targetorigin = targetorigin

	return generateKey(credentials, &generateKeyData)
}

// GenerateKey - Execute the post to generate Key without any encryption
func GenerateKey(credentials *commons.CyberSourceCredential, targetorigin *string) (*GenerateKeyResponse, string, error) {
	var generateKeyData GenerateKeyRequest

	encryptiontype := "None"

	generateKeyData.Encryptiontype = &encryptiontype
	generateKeyData.Targetorigin = targetorigin

	return generateKey(credentials, &generateKeyData)
}

// GenerateMicroformKey - Execute the post to generate the key to use with the microform
func GenerateMicroformKey(credentials *commons.CyberSourceCredential, targetorigin string) (*GenerateKeyResponse, string, error) {
	var generateKeyData GenerateKeyRequest

	encryptiontype := "RsaOaep256"

	generateKeyData.Encryptiontype = &encryptiontype
	generateKeyData.Targetorigin = &targetorigin

	return generateKeyJWT(credentials, &generateKeyData)
}

// generateKeyJWT - Execute the post to generate Key on JWT format
func generateKeyJWT(credentials *commons.CyberSourceCredential, generateKeyData *GenerateKeyRequest) (*GenerateKeyResponse, string, error) {
	payload, err := json.Marshal(generateKeyData)
	if err != nil {
		log.Println("cybersourceflex - generateKeyJWT - Error converting struct to json string.")
		return nil, "Error converting struct generateKeyData to json.", err
	}

	generateKeyResp, err := rest.RestFullPOST(credentials, generateKeyEndpoint+"?format=JWT", string(payload))

	if err != nil {
		log.Println("cybersourceflex - generateKeyJWT - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	return treateGenerateKeyResponse(generateKeyResp)
}

// generateKey - Execute the post to generate Key
func generateKey(credentials *commons.CyberSourceCredential, generateKeyData *GenerateKeyRequest) (*GenerateKeyResponse, string, error) {
	payload, err := json.Marshal(generateKeyData)
	if err != nil {
		log.Println("cybersourceflex - GenerateKey - Error converting struct to json string.")
		return nil, "Error converting struct generateKeyData to json.", err
	}

	generateKeyResp, err := rest.RestFullPOST(credentials, generateKeyEndpoint, string(payload))

	if err != nil {
		log.Println("cybersourceflex - GenerateKey - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	return treateGenerateKeyResponse(generateKeyResp)
}

// treateGenerateKeyResponse - Treate the response to the Generate Key POST request
func treateGenerateKeyResponse(response *rest.RequestResponse) (*GenerateKeyResponse, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	generateKeyResponse := GenerateKeyResponse{}
	err := json.Unmarshal([]byte(response.Body), &generateKeyResponse)

	if err != nil {
		log.Println("cybersourceflex - treateGenerateKeyResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &generateKeyResponse, "Key generated.", nil
}
