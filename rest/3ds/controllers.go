package 3ds

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/rafaelcunha/Go-CyberSource/cybersourcerest"

	"github.com/rafaelcunha/Go-CyberSource/cybersourcecommons"
)

var enrollmentEndpoint = "/risk/v1/authentications"
var validationEndpoint = "/risk/v1/authentication-results"

// EnrollmentRequest - Realize one enrollment request
func EnrollmentRequest(credentials *cybersourcecommons.CyberSourceCredential, data *EnrollmentRequestData) (*EnrollmentResponse, string, error) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("cybersource3ds - EnrollmentRequest - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	fmt.Println("Json Request: ")
	fmt.Println(string(jsonData))

	enrollmentRawResp, err := cybersourcerest.RestFullPOST(credentials, enrollmentEndpoint, string(jsonData))

	if err != nil {
		log.Println("cybersource3ds - EnrollmentRequest - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	return treateEnrollmentRequestResponse(enrollmentRawResp)
}

// treateEnrollmentRequestResponse - Verify and treat the HTTP Post Response to a Enrollment Request
func treateEnrollmentRequestResponse(response *cybersourcerest.RequestResponse) (*EnrollmentResponse, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var enrollmentResp EnrollmentResponse

	err := json.Unmarshal([]byte(response.Body), &enrollmentResp)
	if err != nil {
		log.Println("cybersource3ds - treateEnrollmentRequestResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &enrollmentResp, "Enrollment succssesfully done.", nil
}

// ValidationtRequest - Realize one validation request
func ValidationtRequest(credentials *cybersourcecommons.CyberSourceCredential, data *ValidationRequestData) (*ValidationResponse, string, error) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Println("cybersource3ds - ValidationtRequest - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	log.Println("Json Request: ")
	log.Println(string(jsonData))

	validationRawResp, err := cybersourcerest.RestFullPOST(credentials, validationEndpoint, string(jsonData))

	if err != nil {
		log.Println("cybersource3ds - ValidationtRequest - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	return treateValidationRequestResponse(validationRawResp)
}

// treateValidationRequestResponse - Verify and treat the HTTP Post Response to a Validation Request
func treateValidationRequestResponse(response *cybersourcerest.RequestResponse) (*ValidationResponse, string, error) {
	log.Println("Validation Response:")
	log.Printf("%+v\n", response)

	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var validationResponse ValidationResponse

	err := json.Unmarshal([]byte(response.Body), &validationResponse)
	if err != nil {
		log.Println("cybersource3ds - treateValidationRequestResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &validationResponse, "Validation succssesfully done.", nil
}
