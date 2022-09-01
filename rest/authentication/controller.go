package authentication

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
	"github.com/akayna/Go-dreamBridgeCybersource/utils"
	"github.com/akayna/Go-dreamBridgeUtils/digest"
)

// Constants
var headerSignatureAlgorithm = "HmacSHA256"

// Utitlizando o console, percebi que o campo date não é enviado.
var headerHeadersPostPut = "host date (request-target) digest v-c-merchant-id"
var headerHeadersGet = "host date (request-target) v-c-merchant-id"

// getHeader - Return a struct with CyberSuource header
func GetHeader(credentials *commons.CyberSourceCredential, host, payload, verb, endpoint string) (*RestfullHeader, error) {

	var header RestfullHeader

	// Get actual system time into the RFC1123 format
	var actualDateTime = "Thu, 14 Jul 2022 17:04:36 GMT" //timeutils.GetActualGMTDate() //"Wed, 13 Oct 2021 12:48:33 GMT"
	header.Date = actualDateTime

	// Set the MID
	header.MerchantID = credentials.MID

	// Set the host
	header.Host = host

	// Set the Profile ID
	header.ProfileID = credentials.ProfileID

	if verb == "POST" || verb == "PUT" {
		// Set the header content type
		header.ContentType = "application/json"

		// Generate the digest signature
		genDigest, err := digest.GenerateDigestSHA256(payload)
		if err != nil {
			log.Println("cybersourcerest - getHeader: Error generating Digest signature.")
			return nil, err
		}

		header.Digest = "SHA-256=" + genDigest
	}

	// Mounts the header signature parameter
	headerSignature, err := generateSignature(credentials, verb, actualDateTime, payload, header.Digest, host, endpoint)

	if err != nil {
		log.Println("cybersourcerest - getHeader: Error generating header signature.")
		return nil, err
	}

	header.Signature = headerSignature

	return &header, nil
}

// generateSignature - Generate the requisition signature param
func generateSignature(credentials *commons.CyberSourceCredential, verb, actualDateTime, payload, digestString, host, endpoint string) (*headerSignature, error) {
	var signature headerSignature

	var err error
	signature.Headers, signature.Signature, err = calculateSignature(credentials.SharedSecretKey, host, actualDateTime, endpoint, credentials.MID, verb, digestString)

	if err != nil {
		return nil, err
	}

	signature.APIKey = credentials.APIKeyID
	signature.Algorithm = headerSignatureAlgorithm

	return &signature, nil

}

// calculateSignature - Calculate the signature
func calculateSignature(sharedSecretKey, host, date, target, mid, verb, digestString string) (string, string, error) {

	var headers string

	signatureString := "host: " + host +
		"\ndate: " + date +
		"\n(request-target): " + strings.ToLower(verb) + " " + target

	switch verb {
	case "PUT":
		signatureString += "\ndigest: " + digestString

		headers = headerHeadersPostPut
	case "POST":
		signatureString += "\ndigest: " + digestString

		headers = headerHeadersPostPut
	case "GET":
		headers = headerHeadersGet
	case "DELETE":
		headers = headerHeadersGet
	default:
		return "", "", errors.New("Unknown HTTP verbe: " + verb)
	}

	signatureString += "\nv-c-merchant-id: " + mid

	fmt.Println("Signature String:\n" + signatureString)

	signature, err := utils.GenerateSignature(sharedSecretKey, signatureString)
	if err != nil {
		log.Println("cybersourcerest - calculateSignature: Error generating the signature param.")
		log.Println("error: ", err)
		return "", "", err
	}

	return headers, signature, nil
}

// GetMapString - Return a string map of the RESTFull header
func (header *RestfullHeader) GetMapString() map[string]string {
	headerMap := map[string]string{
		"v-c-merchant-id": header.MerchantID,
		"date":            header.Date,
		"host":            header.Host,
		"digest":          header.Digest,
		"signature":       header.Signature.getString(),
		"Content-Type":    header.ContentType,
		"profile-id":      header.ProfileID,
	}

	return headerMap
}

// GetString - Returns the string to the signature header field
func (header headerSignature) getString() string {
	signature := "keyid=\"" + header.APIKey + "\", algorithm=\"" + header.Algorithm + "\", headers=\"" + header.Headers + "\", signature=\"" + header.Signature + "\""

	return signature
}
