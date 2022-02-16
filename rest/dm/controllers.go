package dm

import (
	"encoding/json"
	"log"

	"github.com/akayna/Go-dreamBridgeCybersource/rest"
	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
)

var createCaseEndpoint = "/risk/v1/decisions"
var listManagementEndpoint = "/risk/v1/lists"

// CreateCase - Creates one Decicision Manager case
func CreateCase(credentials *commons.CyberSourceCredential, dmRequestData *DMRequest) (*DMResponse, string, error) {

	jsonData, err := json.Marshal(dmRequestData)
	if err != nil {
		log.Println("cybersourcedm - CreateCase - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	log.Println("Create Case Json Request: ")
	log.Println(string(jsonData))

	createCaseRawResp, err := rest.RestFullPOST(credentials, createCaseEndpoint, string(jsonData))

	if err != nil {
		log.Println("cybersourcedm - CreateCase - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	log.Println("Create Case Response Data:")
	log.Printf("%+v\n", createCaseRawResp)

	return treatsDMResponse(createCaseRawResp)
}

// treatsDMResponse - Receive and treat a DM response
func treatsDMResponse(response *rest.RequestResponse) (*DMResponse, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var createCaseResp DMResponse

	err := json.Unmarshal([]byte(response.Body), &createCaseResp)
	if err != nil {
		log.Println("cybersourcedm - treatsDMResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &createCaseResp, "DM request succssesfull.", nil
}

// listManagementRequest - Execute a list management request
func listManagementRequest(credentials *commons.CyberSourceCredential, dmRequestData *DMRequest, listName string) (*DMResponse, string, error) {

	jsonData, err := json.Marshal(dmRequestData)
	if err != nil {
		log.Println("cybersourcedm - listManagement - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	log.Println("listManagement Request: ")
	log.Println(string(jsonData))

	listManagementEndpoint := listManagementEndpoint + "/" + listName + "/entries"

	listManagementRawResp, err := rest.RestFullPOST(credentials, listManagementEndpoint, string(jsonData))

	if err != nil {
		log.Println("cybersourcedm - listManagement - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	log.Println("Create Case Response Data:")
	log.Printf("%+v\n", listManagementRawResp)

	return treatsDMResponse(listManagementRawResp)
}

// AddDataToNegativeList - add Data to the negative list
func AddDataToNegativeList(credentials *commons.CyberSourceCredential, dmRequestData *DMRequest) (*DMResponse, string, error) {

	action := "add"

	var markingDetails MarkingDetails
	markingDetails.Action = &action

	var riskInformation RiskInformation
	riskInformation.MarkingDetails = &markingDetails

	dmRequestData.RiskInformation = &riskInformation

	return listManagementRequest(credentials, dmRequestData, "negative")
}

// DeleteDataFromNegativeList - delete Data to the negative list
func DeleteDataFromNegativeList(credentials *commons.CyberSourceCredential, dmRequestData *DMRequest) (*DMResponse, string, error) {

	action := "delete"

	var markingDetails MarkingDetails
	markingDetails.Action = &action

	var riskInformation RiskInformation
	riskInformation.MarkingDetails = &markingDetails

	dmRequestData.RiskInformation = &riskInformation

	return listManagementRequest(credentials, dmRequestData, "negative")
}
