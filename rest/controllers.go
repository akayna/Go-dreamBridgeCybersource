package rest

import (
	"Go-dreamBridgeCybersource/rest/commons"
	"Go-dreamBridgeUtils/digest"
	"Go-dreamBridgeUtils/timeutils"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// Constants
var headerSignatureAlgorithm = "HmacSHA256"

//var headerHeadersPostPut = "host date (request-target) digest v-c-merchant-id"
//var headerHeadersGet = "host date (request-target) v-c-merchant-id"

// Utitlizando o console, percebi que o campo date não é enviado.
var headerHeadersPostPut = "host (request-target) digest v-c-merchant-id"
var headerHeadersGet = "host (request-target) v-c-merchant-id"

var host = "apitest.cybersource.com" // Ambiente de teste
//var host = "api.cybersource.com" // Ambiente produtivo

// Functions

// RestFullSimplePOST - Execute a simple Post call to an endpoint
func RestFullSimplePOST(endpoint, payload string) (*RequestResponse, error) {
	url := "https://" + host + endpoint
	//url := "http://" + host + endpoint

	req, _ := http.NewRequest("POST", url, strings.NewReader(payload))

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("cybersourcerest - RestFullSimplePOST - Error executing POST request.")
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var response = RequestResponse{
		StatusCode: res.StatusCode,
		Body:       string(body),
	}

	return &response, nil
}

// RestFullPOST - Execute a Post call to an endpoint
func RestFullPOST(credentials *commons.CyberSourceCredential, endpoint, payload string) (*RequestResponse, error) {
	url := "https://" + host + endpoint
	//url := "http://" + host + endpoint

	req, _ := http.NewRequest("POST", url, strings.NewReader(payload))

	header, err := getHeader(credentials, host, payload, "POST", endpoint)
	if err != nil {
		log.Println("cybersourcerest - RestFullGET - Error generating Get headers.")
		return nil, err
	}

	headerMap := header.getMapString()

	log.Println("Header: ")

	for key, val := range headerMap {
		if val != "" {
			req.Header.Add(key, val)
			log.Println(key + ": " + val)
		}
	}

	log.Println("Payload:")
	log.Println(payload)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("cybersourcerest - RestFullGET - Error executing GET request.")
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	log.Println(res)

	var response = RequestResponse{
		StatusCode: res.StatusCode,
		Body:       string(body),
	}

	return &response, nil
}

// RestFullDELETE - Execute a Delete call to an endpoint
func RestFullDELETE(credentials *commons.CyberSourceCredential, endpoint string) (*RequestResponse, error) {
	url := "https://" + host + endpoint

	req, _ := http.NewRequest("DELETE", url, nil)

	header, err := getHeader(credentials, host, "", "DELETE", endpoint)
	if err != nil {
		log.Println("cybersourcerest - RestFullDELETE - Error generating headers.")
		return nil, err
	}

	headerMap := header.getMapString()

	for key, val := range headerMap {
		if val != "" {
			req.Header.Add(key, val)
			//fmt.Println(key + ": " + val)
		}
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("cybersourcerest - RestFullDELETE - Error executing DELETE request.")
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var response = RequestResponse{
		StatusCode: res.StatusCode,
		Body:       string(body),
	}

	return &response, nil
}

// RestFullGET - Execute a Get call to an endpoint
func RestFullGET(credentials *commons.CyberSourceCredential, endpoint string) (*RequestResponse, error) {
	url := "https://" + host + endpoint

	log.Println("Get URL: " + url)

	req, _ := http.NewRequest("GET", url, nil)

	header, err := getHeader(credentials, host, "", "GET", endpoint)
	if err != nil {
		log.Println("cybersourcerest - RestFullGET - Error generating Get headers.")
		return nil, err
	}

	headerMap := header.getMapString()

	for key, val := range headerMap {
		if val != "" {
			req.Header.Add(key, val)
			//fmt.Println(key + ": " + val)
		}
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("cybersourcerest - RestFullGET - Error executing GET request.")
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var response = RequestResponse{
		StatusCode: res.StatusCode,
		Body:       string(body),
	}

	return &response, nil
}

/*
// GetPostHeader -Return the header of a POST request
func GetPostHeader(credentials *cybersourcecommons.CyberSourceCredential, endpoint, payload string) (*RestfullHeader, error) {
	return getHeader(credentials, host, payload, "post", endpoint)
}

// GetPutHeader -Return the header of a PUT request
func GetPutHeader(credentials *cybersourcecommons.CyberSourceCredential, endpoint, payload string) (*RestfullHeader, error) {
	return getHeader(credentials, host, payload, "put", endpoint)
}

// GetGetHeader -Return the header of a GET request
func GetGetHeader(credentials *cybersourcecommons.CyberSourceCredential, endpoint, id string) (*RestfullHeader, error) {
	return getHeader(credentials, host, "", "get", (endpoint + id))
}
*/
// GetMapString - Return a string map of the RESTFull header
func (header *RestfullHeader) getMapString() map[string]string {
	headerMap := map[string]string{
		"v-c-merchant-id": header.MerchantID,
		"v-c-date":        header.Date, // Na documentação esse campo seria Date:
		"host":            header.Host,
		"digest":          header.Digest,
		"signature":       header.Signature.getString(),
		"Content-Type":    header.ContentType,
		"profile-id":      header.ProfileID,
	}

	return headerMap
}

// getHeader - Return a struct with CyberSuource header
func getHeader(credentials *commons.CyberSourceCredential, host, payload, verb, endpoint string) (*RestfullHeader, error) {

	var header RestfullHeader

	// Get actual system time into the RFC1123 format
	actualDateTime := timeutils.GetActualGMTDate()
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
		genDigest, err := digest.GenerateDigest(payload)
		if err != nil {
			log.Println("cybersourcerest - getHeader: Error generating Digest signature.")
			return nil, err
		}

		header.Digest = "SHA-256=" + genDigest
	}

	// Mounts the header signature parameter
	headerSignature, err := generateSignature(credentials, verb, actualDateTime, payload, header.Digest, endpoint)

	if err != nil {
		log.Println("cybersourcerest - getHeader: Error generating header signature.")
		return nil, err
	}

	header.Signature = headerSignature

	return &header, nil
}

// generateSignature - Generate the requisition signature param
func generateSignature(credentials *commons.CyberSourceCredential, verb, actualDateTime, payload, digestString, endpoint string) (*headerSignature, error) {
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
		//"\ndate: " + date +
		"\n(request-target): " + strings.ToLower(verb) + " " + target

	switch verb {
	case "PUT":
		signatureString += "\ndigest: " + digestString

		headers = headerHeadersPostPut
		break
	case "POST":
		signatureString += "\ndigest: " + digestString

		headers = headerHeadersPostPut
		break
	case "GET":
		headers = headerHeadersGet
		break
	case "DELETE":
		headers = headerHeadersGet
		break

	default:
		return "", "", errors.New("Unknown HTTP verbe: " + verb)
	}

	signatureString += "\nv-c-merchant-id: " + mid

	fmt.Println("Signature String: " + signatureString)

	signature, err := digest.GenerateSignature(sharedSecretKey, signatureString)
	if err != nil {
		log.Println("cybersourcerest - calculateSignature: Error generating the signature param.")
		log.Println("error: ", err)
		return "", "", err
	}

	return headers, signature, nil
}

// GetString - Returns the string to the signature header field
func (header headerSignature) getString() string {
	signature := "keyid=\"" + header.APIKey + "\", algorithm=\"" + header.Algorithm + "\", headers=\"" + header.Headers + "\", signature=\"" + header.Signature + "\""

	//fmt.Println("signature:\n" + signature)

	return signature
}
