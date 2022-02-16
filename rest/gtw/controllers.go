package gtw

import (
	"encoding/json"
	"log"

	"github.com/akayna/Go-dreamBridgeCybersource/rest"
	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
)

var paymentsEndpoint = "/pts/v2/payments"
var capturesEndpoint = "/pts/v2/captures"
var refundsEndpoint = "/pts/v2/refunds"

// VoidRefundByID - Void a Refund by its ID
func VoidRefundByID(credentials *commons.CyberSourceCredential, refundID string, voidRefundData *Payment) (*Payment, string, error) {
	voidRefundEndpoint := refundsEndpoint + "/" + refundID + "/voids"
	return voidTransaction(credentials, voidRefundEndpoint, voidRefundData)
}

// VoidRefund - Void a Refund
func VoidRefund(credentials *commons.CyberSourceCredential, voidRefundEndpoint string, voidRefundData *Payment) (*Payment, string, error) {
	return voidTransaction(credentials, voidRefundEndpoint, voidRefundData)
}

// VoidCaptureByID - Void a Capture by its ID
func VoidCaptureByID(credentials *commons.CyberSourceCredential, captureID string, voidCaptureData *Payment) (*Payment, string, error) {
	voidCaptureEndpoint := capturesEndpoint + "/" + captureID + "/voids"
	return voidTransaction(credentials, voidCaptureEndpoint, voidCaptureData)
}

// VoidCapture - Void a Capture
func VoidCapture(credentials *commons.CyberSourceCredential, voidCaptureEndpoint string, voidCaptureData *Payment) (*Payment, string, error) {
	return voidTransaction(credentials, voidCaptureEndpoint, voidCaptureData)
}

// VoidPaymentByID - Void a Payment by its ID
func VoidPaymentByID(credentials *commons.CyberSourceCredential, paymentID string, voidPaymentData *Payment) (*Payment, string, error) {
	voidPaymentEndpoint := paymentsEndpoint + "/" + paymentID + "/voids"
	return voidTransaction(credentials, voidPaymentEndpoint, voidPaymentData)
}

// VoidPayment - Void a Payment
func VoidPayment(credentials *commons.CyberSourceCredential, voidPaymentEndpoint string, voidPaymentData *Payment) (*Payment, string, error) {
	return voidTransaction(credentials, voidPaymentEndpoint, voidPaymentData)
}

// RefundCapture - Refund a capture
func voidTransaction(credentials *commons.CyberSourceCredential, voidEndpoint string, voidData *Payment) (*Payment, string, error) {

	jsonData, err := json.Marshal(voidData)
	if err != nil {
		log.Println("cybersourcegateway - voidTransaction - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	//log.Println("Void Json Request: ")
	//log.Println(string(jsonData))

	voidRawResp, err := rest.RestFullPOST(credentials, voidEndpoint, string(jsonData))

	if err != nil {
		log.Println("cybersourcegateway - voidTransaction - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	//log.Println("Void Response Data:")
	//log.Printf("%+v\n", voidRawResp)

	return treatsVoidResponse(voidRawResp)
}

// treatsRefundCaptureResponse - Verify and treat the HTTP POST Response of a refund capture request
func treatsVoidResponse(response *rest.RequestResponse) (*Payment, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var voidResp Payment

	err := json.Unmarshal([]byte(response.Body), &voidResp)
	if err != nil {
		log.Println("cybersourcegateway - treatsVoidResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &voidResp, "Void succssesfull.", nil
}

// RefundCaptureByID - Refund a capture by its id
func RefundCaptureByID(credentials *commons.CyberSourceCredential, captureID string, refundData *Payment) (*Payment, string, error) {
	refundEndpoint := capturesEndpoint + "/" + captureID + "/refunds"
	return RefundCapture(credentials, refundEndpoint, refundData)
}

// RefundCapture - Refund a capture
func RefundCapture(credentials *commons.CyberSourceCredential, refundEndpoint string, refundData *Payment) (*Payment, string, error) {

	jsonData, err := json.Marshal(refundData)
	if err != nil {
		log.Println("cybersourcegateway - RefundCapture - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	log.Println("Refund Capture Json Request: ")
	log.Println(string(jsonData))

	refundCaptureRawResp, err := rest.RestFullPOST(credentials, refundEndpoint, string(jsonData))

	if err != nil {
		log.Println("cybersourcegateway - RefundCapture - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	log.Println("Refund Capture Response Data:")
	log.Printf("%+v\n", refundCaptureRawResp)

	return treatsRefundCaptureResponse(refundCaptureRawResp)
}

// treatsRefundCaptureResponse - Verify and treat the HTTP POST Response of a refund capture request
func treatsRefundCaptureResponse(response *rest.RequestResponse) (*Payment, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var refundCaptureResp Payment

	err := json.Unmarshal([]byte(response.Body), &refundCaptureResp)
	if err != nil {
		log.Println("cybersourcegateway - treatsRefundCaptureResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &refundCaptureResp, "Capture Refund succssesfull.", nil
}

// RefundPaymentByID - Refund a payment by its id
func RefundPaymentByID(credentials *commons.CyberSourceCredential, paymentID string, refundData *Payment) (*Payment, string, error) {
	refundEndpoint := paymentsEndpoint + "/" + paymentID + "/refunds"
	return RefundPayment(credentials, refundEndpoint, refundData)
}

// RefundPayment - Refund a payment
func RefundPayment(credentials *commons.CyberSourceCredential, refundEndpoint string, refundData *Payment) (*Payment, string, error) {

	jsonData, err := json.Marshal(refundData)
	if err != nil {
		log.Println("cybersourcegateway - RefundPayment - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	log.Println("Refund Payment Json Request: ")
	log.Println(string(jsonData))

	refundPaymentRawResp, err := rest.RestFullPOST(credentials, refundEndpoint, string(jsonData))

	if err != nil {
		log.Println("cybersourcegateway - RefundPayment - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	log.Println("Refund Payment Response Data:")
	log.Printf("%+v\n", refundPaymentRawResp)

	return treatsRefundPaymentResponse(refundPaymentRawResp)
}

// treatsRefundPaymentResponse - Verify and treat the HTTP POST Response of a refund payment request
func treatsRefundPaymentResponse(response *rest.RequestResponse) (*Payment, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var refundPaymentResp Payment

	err := json.Unmarshal([]byte(response.Body), &refundPaymentResp)
	if err != nil {
		log.Println("cybersourcegateway - treatsRefundPaymentResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &refundPaymentResp, "Payment Refund succssesfull.", nil
}

// CapturePaymentByID - Capture payment by its id
func CapturePaymentByID(credentials *commons.CyberSourceCredential, paymentID string, captureData *Payment) (*Payment, string, error) {
	captureEndpoint := paymentsEndpoint + "/" + paymentID + "/captures"
	return CapturePayment(credentials, captureEndpoint, captureData)
}

// CapturePayment - Capture a payment
func CapturePayment(credentials *commons.CyberSourceCredential, captureEndpoint string, captureData *Payment) (*Payment, string, error) {

	jsonData, err := json.Marshal(captureData)
	if err != nil {
		log.Println("cybersourcegateway - CapturePayment - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	log.Println("Capture Payment Json Request: ")
	log.Println(string(jsonData))

	capturePaymentRawResp, err := rest.RestFullPOST(credentials, captureEndpoint, string(jsonData))

	if err != nil {
		log.Println("cybersourcegateway - CapturePayment - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	log.Println("Capture Payment Response Data:")
	log.Printf("%+v\n", capturePaymentRawResp)

	return treatsCapturePaymentResponse(capturePaymentRawResp)
}

// treatsCapturePaymentResponse - Verify and treat the HTTP POST Response of a payment capture
func treatsCapturePaymentResponse(response *rest.RequestResponse) (*Payment, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var captureInfoResp Payment

	err := json.Unmarshal([]byte(response.Body), &captureInfoResp)
	if err != nil {
		log.Println("cybersourcegateway - treatsCapturePaymentResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &captureInfoResp, "Payment capture succssesfull.", nil
}

// ProcessAuthorizationReversalByID - Process one authorization reversal using its ID
func ProcessAuthorizationReversalByID(credentials *commons.CyberSourceCredential, paymentID string, reversalData *Payment) (*Payment, string, error) {
	reversalEndpoint := paymentsEndpoint + "/" + paymentID + "/reversals"

	return ProcessAuthorizationReversal(credentials, reversalEndpoint, reversalData)
}

// ProcessAuthorizationReversal - Process one authorization reversal
func ProcessAuthorizationReversal(credentials *commons.CyberSourceCredential, reversalEndpoint string, reversalData *Payment) (*Payment, string, error) {

	jsonData, err := json.Marshal(reversalData)
	if err != nil {
		log.Println("cybersourcegateway - ProcessAuthorizationReversal - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	log.Println("Json Request: ")
	log.Println(string(jsonData))

	authorizationReversalRawResp, err := rest.RestFullPOST(credentials, reversalEndpoint, string(jsonData))

	if err != nil {
		log.Println("cybersourcegateway - ProcessAuthorizationReversal - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	log.Println("Reverse Authorization Response Data:")
	log.Printf("%+v\n", authorizationReversalRawResp)

	return treatsProcessAuthorizationReversalResponse(authorizationReversalRawResp)
}

// treatsProcessAuthorizationReversal - Verify and treat the HTTP POST Response of an authorization reversal
func treatsProcessAuthorizationReversalResponse(response *rest.RequestResponse) (*Payment, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var paymentInfoResp Payment

	err := json.Unmarshal([]byte(response.Body), &paymentInfoResp)
	if err != nil {
		log.Println("cybersourcegateway - treatsProcessAuthorizationReversal - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &paymentInfoResp, "Reversal realized succssesfully.", nil
}

// GetPaymentInfoByID - Retrieves payment info by payment id
func GetPaymentInfoByID(credentials *commons.CyberSourceCredential, paymentID string) (*Payment, string, error) {
	return GetPaymentInfo(credentials, paymentsEndpoint+"/"+paymentID)
}

// GetPaymentInfo - Retrieves payment information
func GetPaymentInfo(credentials *commons.CyberSourceCredential, endpointURL string) (*Payment, string, error) {

	paymentInfoRawResp, err := rest.RestFullGET(credentials, endpointURL)

	if err != nil {
		log.Println("cybersourcegateway - GetPaymentInfo - Error executing GET request.")
		return nil, "Error executing GET request.", err
	}

	//log.Println("Response Data:")
	//log.Printf("%+v\n", paymentInfoRawResp)

	return treatsGetPaymentInfoResponse(paymentInfoRawResp)
}

// treatsGetPaymentInfoResponse - Verify and treat the HTTP GET Response of a get payment info
func treatsGetPaymentInfoResponse(response *rest.RequestResponse) (*Payment, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var paymentInfoResp Payment

	err := json.Unmarshal([]byte(response.Body), &paymentInfoResp)
	if err != nil {
		log.Println("cybersourcegateway - treatsGetPaymentInfoResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &paymentInfoResp, "Information retrieved succssesfully.", nil
}

// ProcessPayment - Post one payment request
func ProcessPayment(credentials *commons.CyberSourceCredential, paymentData *Payment) (*Payment, string, error) {

	jsonData, err := json.Marshal(paymentData)
	if err != nil {
		log.Println("cybersourcegateway - AuthorizePayment - Error converting struct to json string.")
		return nil, "Error converting struct to json string.", err
	}

	log.Println("Json Request: ")
	log.Println(string(jsonData))

	paymentRawResp, err := rest.RestFullPOST(credentials, paymentsEndpoint, string(jsonData))

	if err != nil {
		log.Println("cybersourcegateway - ProcessPayment - Error executing POST request.")
		return nil, "Error executing POST request.", err
	}

	log.Println("Authorization Response Data:")
	log.Printf("%+v\n", paymentRawResp)

	return treatePaymentRequestResponse(paymentRawResp)
}

// treatePaymentRequestResponse - Verify and treat the HTTP Post Response to authorization request
func treatePaymentRequestResponse(response *rest.RequestResponse) (*Payment, string, error) {
	if response.StatusCode > 299 || response.StatusCode < 200 {

		return nil, response.Body, nil
	}

	var paymentResp Payment

	err := json.Unmarshal([]byte(response.Body), &paymentResp)
	if err != nil {
		log.Println("cybersourcegateway - treatePaymentRequestResponse - Error converting json string to struct.")
		return nil, "Error converting json string to struct.", err
	}

	return &paymentResp, "Authorization succssesfully done.", nil
}
