package dm

import (
	"encoding/json"
	"log"

	"github.com/rafaelcunha/Go-CyberSource/cybersourcecommons"
	"github.com/rafaelcunha/Go-CyberSource/cybersourcerest"
)

var createCaseEndpoint = "/risk/v1/decisions"

// CreateCase - Creates one Decicision Manager case
func CreateCase(credentials *cybersourcecommons.CyberSourceCredential, createCaseData *CreateCaseRequest) (*CreateCaseResponse, string, error) {

	jsonData, err := json.Marshal(createCaseData)
	if err != nil {
		log.Println("cybersourcedm - CreateCase - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	log.Println("Create Case Json Request: ")
	log.Println(string(jsonData))

	createCaseRawResp, err := cybersourcerest.RestFullPOST(credentials, createCaseEndpoint, string(jsonData))

	if err != nil {
		log.Println("cybersourcedm - CreateCase - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	log.Println("Create Case Response Data:")
	log.Printf("%+v\n", createCaseRawResp)

	return treatsCreateCaseResponse(createCaseRawResp)
}

func treatsCreateCaseResponse(response *cybersourcerest.RequestResponse) (*CreateCaseResponse, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var createCaseResp CreateCaseResponse

	err := json.Unmarshal([]byte(response.Body), &createCaseResp)
	if err != nil {
		log.Println("cybersourcedm - treatsCreateCaseResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &createCaseResp, "Create case succssesfull.", nil
}
