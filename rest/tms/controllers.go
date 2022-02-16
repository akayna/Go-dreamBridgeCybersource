package tms

import (
	"encoding/json"
	"log"

	"github.com/akayna/Go-dreamBridgeCybersource/rest"
	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
)

var tmsPaymentInstrumentEndpoint_v1 = "/tms/v1/paymentinstruments"
var tmsInstrumentIdentifierEndpoint_v1 = "/tms/v1/instrumentidentifiers"
var tmsTokensEndpoint_v2 = "/tms/v2/tokens"

// DeleteInstrumentIdentifier - Deletes the instrument identifier
func DeleteInstrumentIdentifier(credentials *commons.CyberSourceCredential, instrumentIdentifierID string) (bool, string, error) {

	response, err := rest.RestFullDELETE(credentials, (tmsInstrumentIdentifierEndpoint_v1 + "/" + instrumentIdentifierID))

	if err != nil {
		log.Println("cybersourcetms - DeleteInstrumentIdentifier - Error executing Delete request.")
		return false, "", err
	}

	return treateInstrumentIdentifierDeleteResponse(response)
}

// treateInstrumentIdentifierDeleteResponse - Verify the HTTP Response and treat the delete response.
func treateInstrumentIdentifierDeleteResponse(response *rest.RequestResponse) (bool, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {
		return false, response.Body, nil
	}

	return true, "Instrument Identifier deleted.", nil

}

// RetrieveInstrumentIdentifier - Retrieves the instrument identifier
func RetrieveInstrumentIdentifier(credentials *commons.CyberSourceCredential, instrumentIdentifierID string) (*InstrumentIdentifierResponse, string, error) {
	requestResp, err := rest.RestFullGET(credentials, (tmsInstrumentIdentifierEndpoint_v1 + "/" + instrumentIdentifierID))

	if err != nil {
		log.Println("cybersourcetms - RetrieveInstrumentIdentifier - Error executing GET request.")
		return nil, "Error executing GET request.", err
	}

	return treateRetrieveInstrumentIdentifierResponse(requestResp)
}

// treateRetrieveInstrumentIdentifierResponse - Verify and treat the HTTP Post Response
func treateRetrieveInstrumentIdentifierResponse(response *rest.RequestResponse) (*InstrumentIdentifierResponse, string, error) {
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
func CreateInstrumentIdentifier(credentials *commons.CyberSourceCredential, instrumentIdentifierRequestData *CreateInstrumentIdentifierRequest) (*InstrumentIdentifierResponse, string, error) {

	payload, err := json.Marshal(instrumentIdentifierRequestData)
	if err != nil {
		log.Println("cybersourcetms - CreateInstrumentIdentifier - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	createResp, err := rest.RestFullPOST(credentials, tmsInstrumentIdentifierEndpoint_v1, string(payload))

	if err != nil {
		log.Println("cybersourcetms - CreateInstrumentIdentifier - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	return treateCreateInstrumentIdentifierResponse(createResp)
}

// treateCreateInstrumentIdentifierResponse - Verify and treat the HTTP Post Response
func treateCreateInstrumentIdentifierResponse(response *rest.RequestResponse) (*InstrumentIdentifierResponse, string, error) {
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
func CreatePaymentInstrument(credentials *commons.CyberSourceCredential, paymentInstrument *CreatePaymentInstrumentRequest) (*PaymentInstrumentResponse, string, error) {

	payload, err := json.Marshal(paymentInstrument)
	if err != nil {
		log.Println("cybersourcetms - CeratePaymentInstrument - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	//fmt.Println("Payload:\n" + string(payload))

	createResp, err := rest.RestFullPOST(credentials, tmsPaymentInstrumentEndpoint_v1, string(payload))

	if err != nil {
		log.Println("cybersourcetms - CeratePaymentInstrument - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	return treateCreatePaymentInstrumentResponse(createResp)
}

// treateCreatePaymentInstrumentResponse - Verify and treat the HTTP Post Response
func treateCreatePaymentInstrumentResponse(response *rest.RequestResponse) (*PaymentInstrumentResponse, string, error) {
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
func RetrievePaymentInstrument(credentials *commons.CyberSourceCredential, paymentInstrumentID string) (*PaymentInstrumentResponse, string, error) {

	requestResp, err := rest.RestFullGET(credentials, (tmsPaymentInstrumentEndpoint_v1 + "/" + paymentInstrumentID))

	if err != nil {
		log.Println("cybersourcetms - RetrievePaymentInstrument - Error executing GET request.")
		return nil, "Error executing GET request.", err
	}

	return treateRetrievePaymentInstrumentResponse(requestResp)
}

// treateRetrievePaymentInstrumentResponse - Verify the HTTP Response and treat the get response.
func treateRetrievePaymentInstrumentResponse(response *rest.RequestResponse) (*PaymentInstrumentResponse, string, error) {
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
func DeletePaymentInstrument(credentials *commons.CyberSourceCredential, paymentInstrumentID string) (bool, string, error) {

	response, err := rest.RestFullDELETE(credentials, (tmsPaymentInstrumentEndpoint_v1 + "/" + paymentInstrumentID))

	if err != nil {
		log.Println("cybersourcetms - RetrievePaymentInstrument - Error executing Delete request.")
		return false, "", err
	}

	return treatePaymentInstrumentDeleteResponse(response)
}

// treatePaymentInstrumentDeleteResponse - Verify the HTTP Response and treat the delete response.
func treatePaymentInstrumentDeleteResponse(response *rest.RequestResponse) (bool, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {
		return false, response.Body, nil
	}

	return true, "Payment instrument deleted.", nil

}

// GenerateCryptogram - Execute the call to generate one cryptogram from an instrument identifier
func GenerateCryptogram(credentials *commons.CyberSourceCredential, instrumentIdentifierID string) (string, string, error) {
	createCryptogramResp, err := rest.RestFullPOST(credentials, (tmsTokensEndpoint_v2 + "/" + instrumentIdentifierID + "/cryptograms"), "{}")

	if err != nil {
		log.Println("cybersourcetms - GenerateCryptogram - Error creating a cryptogram.")
		return "", "Error executing GET request.", err
	}

	return createCryptogramResp.Body, "Cryptogram generated.", nil
}
