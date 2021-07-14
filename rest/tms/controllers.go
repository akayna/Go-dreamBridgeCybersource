package tms

import (
	"encoding/json"
	"log"

	"github.com/rafaelcunha/Go-CyberSource/cybersourcecommons"
	"github.com/rafaelcunha/Go-CyberSource/cybersourcerest"
)

var tmsPaymentInstrumentEndpoint = "/tms/v1/paymentinstruments"
var tmsInstrumentIdentifierEndpoint = "/tms/v1/instrumentidentifiers"

// DeleteInstrumentIdentifier - Deletes the instrument identifier
func DeleteInstrumentIdentifier(credentials *cybersourcecommons.CyberSourceCredential, instrumentIdentifierID string) (bool, string, error) {

	response, err := cybersourcerest.RestFullDELETE(credentials, (tmsInstrumentIdentifierEndpoint + "/" + instrumentIdentifierID))

	if err != nil {
		log.Println("cybersourcetms - DeleteInstrumentIdentifier - Error executing Delete request.")
		return false, "", err
	}

	return treateInstrumentIdentifierDeleteResponse(response)
}

// treateInstrumentIdentifierDeleteResponse - Verify the HTTP Response and treat the delete response.
func treateInstrumentIdentifierDeleteResponse(response *cybersourcerest.RequestResponse) (bool, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {
		return false, response.Body, nil
	}

	return true, "Instrument Identifier deleted.", nil

}

// RetrieveInstrumentIdentifier - Retrieves the instrument identifier
func RetrieveInstrumentIdentifier(credentials *cybersourcecommons.CyberSourceCredential, instrumentIdentifierID string) (*InstrumentIdentifierResponse, string, error) {
	requestResp, err := cybersourcerest.RestFullGET(credentials, (tmsInstrumentIdentifierEndpoint + "/" + instrumentIdentifierID))

	if err != nil {
		log.Println("cybersourcetms - RetrieveInstrumentIdentifier - Error executing GET request.")
		return nil, "Error executing GET request.", err
	}

	return treateRetrieveInstrumentIdentifierResponse(requestResp)
}

// treateRetrieveInstrumentIdentifierResponse - Verify and treat the HTTP Post Response
func treateRetrieveInstrumentIdentifierResponse(response *cybersourcerest.RequestResponse) (*InstrumentIdentifierResponse, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	instrumentIdentifierResponse := InstrumentIdentifierResponse{}
	err := json.Unmarshal([]byte(response.Body), &instrumentIdentifierResponse)

	if err != nil {
		log.Println("cybersourcetms - treateRetrieveInstrumentIdentifierResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &instrumentIdentifierResponse, "Payment Instrument retrieved.", nil
}

// CreateInstrumentIdentifier - Execute the Create Instrument Identifier Request.
func CreateInstrumentIdentifier(credentials *cybersourcecommons.CyberSourceCredential, instrumentIdentifierRequestData *CreateInstrumentIdentifierRequest) (*InstrumentIdentifierResponse, string, error) {

	payload, err := json.Marshal(instrumentIdentifierRequestData)
	if err != nil {
		log.Println("cybersourcetms - CreateInstrumentIdentifier - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	createResp, err := cybersourcerest.RestFullPOST(credentials, tmsInstrumentIdentifierEndpoint, string(payload))

	if err != nil {
		log.Println("cybersourcetms - CreateInstrumentIdentifier - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	return treateCreateInstrumentIdentifierResponse(createResp)
}

// treateCreateInstrumentIdentifierResponse - Verify and treat the HTTP Post Response
func treateCreateInstrumentIdentifierResponse(response *cybersourcerest.RequestResponse) (*InstrumentIdentifierResponse, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	instrumentIdentifierResponse := InstrumentIdentifierResponse{}
	err := json.Unmarshal([]byte(response.Body), &instrumentIdentifierResponse)

	if err != nil {
		log.Println("cybersourcetms - treateCreatePaymentInstrumentResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &instrumentIdentifierResponse, "Instrument Identifier created.", nil
}

// CreatePaymentInstrument - Create a Payment Instrument
func CreatePaymentInstrument(credentials *cybersourcecommons.CyberSourceCredential, paymentInstrument *CreatePaymentInstrumentRequest) (*PaymentInstrumentResponse, string, error) {

	payload, err := json.Marshal(paymentInstrument)
	if err != nil {
		log.Println("cybersourcetms - CeratePaymentInstrument - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	//fmt.Println("Payload:\n" + string(payload))

	createResp, err := cybersourcerest.RestFullPOST(credentials, tmsPaymentInstrumentEndpoint, string(payload))

	if err != nil {
		log.Println("cybersourcetms - CeratePaymentInstrument - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	return treateCreatePaymentInstrumentResponse(createResp)
}

// treateCreatePaymentInstrumentResponse - Verify and treat the HTTP Post Response
func treateCreatePaymentInstrumentResponse(response *cybersourcerest.RequestResponse) (*PaymentInstrumentResponse, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	paymentInstrumentResponse := PaymentInstrumentResponse{}
	err := json.Unmarshal([]byte(response.Body), &paymentInstrumentResponse)

	if err != nil {
		log.Println("cybersourcetms - treateCreatePaymentInstrumentResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &paymentInstrumentResponse, "Payment Instrument created.", nil

}

// RetrievePaymentInstrument - Return the payment instrument json
func RetrievePaymentInstrument(credentials *cybersourcecommons.CyberSourceCredential, paymentInstrumentID string) (*PaymentInstrumentResponse, string, error) {

	requestResp, err := cybersourcerest.RestFullGET(credentials, (tmsPaymentInstrumentEndpoint + "/" + paymentInstrumentID))

	if err != nil {
		log.Println("cybersourcetms - RetrievePaymentInstrument - Error executing GET request.")
		return nil, "Error executing GET request.", err
	}

	return treateRetrievePaymentInstrumentResponse(requestResp)
}

// treateRetrievePaymentInstrumentResponse - Verify the HTTP Response and treat the get response.
func treateRetrievePaymentInstrumentResponse(response *cybersourcerest.RequestResponse) (*PaymentInstrumentResponse, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	paymentInstrumentGetResponse := PaymentInstrumentResponse{}
	err := json.Unmarshal([]byte(response.Body), &paymentInstrumentGetResponse)

	if err != nil {
		log.Println("cybersourcetms - treateRetrievePaymentInstrumentResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &paymentInstrumentGetResponse, "Payment Instrument retrieved.", nil
}

// DeletePaymentInstrument - Deletes the payment instrument
func DeletePaymentInstrument(credentials *cybersourcecommons.CyberSourceCredential, paymentInstrumentID string) (bool, string, error) {

	response, err := cybersourcerest.RestFullDELETE(credentials, (tmsPaymentInstrumentEndpoint + "/" + paymentInstrumentID))

	if err != nil {
		log.Println("cybersourcetms - RetrievePaymentInstrument - Error executing Delete request.")
		return false, "", err
	}

	return treatePaymentInstrumentDeleteResponse(response)
}

// treatePaymentInstrumentDeleteResponse - Verify the HTTP Response and treat the delete response.
func treatePaymentInstrumentDeleteResponse(response *cybersourcerest.RequestResponse) (bool, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {
		return false, response.Body, nil
	}

	return true, "Payment instrument deleted.", nil

}