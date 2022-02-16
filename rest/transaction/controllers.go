package transaction

import (
	"encoding/json"
	"log"

	"github.com/akayna/Go-dreamBridgeCybersource/rest"
	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
)

var transactionsEndpoint = "/tss/v2/transactions/"

// GetTransactionByID - Retrieves a transaction info by its id
func GetTransactionByID(credentials *commons.CyberSourceCredential, transactionID string) (*Transaction, string, error) {
	transactionsEndpoint := transactionsEndpoint + transactionID
	return GetTransaction(credentials, transactionsEndpoint)
}

// GetTransaction - Retrieves a transaction
func GetTransaction(credentials *commons.CyberSourceCredential, endpointURL string) (*Transaction, string, error) {

	transactionRawResp, err := rest.RestFullGET(credentials, endpointURL)

	if err != nil {
		log.Println("cybersourcegateway - GetTransaction - Error executing GET request.")
		return nil, "Error executing GET request.", err
	}

	//log.Println("Response Data:")
	//log.Printf("%+v\n", transactionRawResp)

	return treatsGetTransactionResponse(transactionRawResp)
}

// treatsGetTransactionResponse - Verify and treat the HTTP GET Response of a get transaction
func treatsGetTransactionResponse(response *rest.RequestResponse) (*Transaction, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var transactionResp Transaction

	err := json.Unmarshal([]byte(response.Body), &transactionResp)
	if err != nil {
		log.Println("cybersourcegateway - treatsGetTransactionResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &transactionResp, "Transaction retrieved succssesfully.", nil
}
