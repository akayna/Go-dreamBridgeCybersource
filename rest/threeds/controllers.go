package threeds

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/akayna/Go-dreamBridgeCybersource/rest"
	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
)

var setupPayerAuthEndpoint = "/risk/v1/authentication-setups"
var enrollmentEndpoint = "/risk/v1/authentications"
var validationEndpoint = "/risk/v1/authentication-results"

// SetupPayerAuthRequest - Realize the setup for cruize integration.
func SetupPayerAuthRequest(credentials *commons.CyberSourceCredential, data *SetupPayerAuthRequestData) (*SetupPayerAuthResponseData, string, error) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("cybersource3ds - SetupPayerAuthRequest - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	fmt.Println("Json Request: ")
	fmt.Println(string(jsonData))

	setupRawResp, err := rest.RestFullPOST(credentials, setupPayerAuthEndpoint, string(jsonData))

	if err != nil {
		log.Println("cybersource3ds - SetupPayerAuthRequest - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	return treatSetupPayerAuthResponse(setupRawResp)
}

// treatSetupPayerAuthResponse - Verify and treat the HTTP Post Response to a Setup Request
func treatSetupPayerAuthResponse(response *rest.RequestResponse) (*SetupPayerAuthResponseData, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var setupPayerAuthResp SetupPayerAuthResponseData

	err := json.Unmarshal([]byte(response.Body), &setupPayerAuthResp)
	if err != nil {
		log.Println("cybersource3ds - treatEnrollmentRequestResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &setupPayerAuthResp, "Enrollment succssesfully done.", nil
}

// EnrollmentRequest - Realize one enrollment request
func EnrollmentRequest(credentials *commons.CyberSourceCredential, data *EnrollmentRequestData) (*EnrollmentResponse, string, error) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("cybersource3ds - EnrollmentRequest - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	fmt.Println("Json Request: ")
	fmt.Println(string(jsonData))

	enrollmentRawResp, err := rest.RestFullPOST(credentials, enrollmentEndpoint, string(jsonData))

	if err != nil {
		log.Println("cybersource3ds - EnrollmentRequest - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	return treatEnrollmentRequestResponse(enrollmentRawResp)
}

// treatEnrollmentRequestResponse - Verify and treat the HTTP Post Response to a Enrollment Request
func treatEnrollmentRequestResponse(response *rest.RequestResponse) (*EnrollmentResponse, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var enrollmentResp EnrollmentResponse

	err := json.Unmarshal([]byte(response.Body), &enrollmentResp)
	if err != nil {
		log.Println("cybersource3ds - treatEnrollmentRequestResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &enrollmentResp, "Enrollment succssesfully done.", nil
}

// ValidationtRequest - Realize one validation request
func ValidationtRequest(credentials *commons.CyberSourceCredential, data *ValidationRequestData) (*ValidationResponse, string, error) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("cybersource3ds - ValidationtRequest - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	log.Println("Json Request: ")
	log.Println(string(jsonData))

	validationRawResp, err := rest.RestFullPOST(credentials, validationEndpoint, string(jsonData))

	if err != nil {
		log.Println("cybersource3ds - ValidationtRequest - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	return treatValidationRequestResponse(validationRawResp)
}

// treatValidationRequestResponse - Verify and treat the HTTP Post Response to a Validation Request
func treatValidationRequestResponse(response *rest.RequestResponse) (*ValidationResponse, string, error) {
	log.Println("Validation Response:")
	log.Printf("%+v\n", response)

	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var validationResponse ValidationResponse

	err := json.Unmarshal([]byte(response.Body), &validationResponse)
	if err != nil {
		log.Println("cybersource3ds - treatValidationRequestResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &validationResponse, "Validation succssesfully done.", nil
}
