package threeds

import (
	"time"

	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
)

// TransactionModeMOTO - Mail Order Telephone Order
const TransactionModeMOTO = "MOTO"

// TransactionModeRETAIL - Retail
const TransactionModeRETAIL = "RETAIL"

// TransactionModeECOMMERCE - eCommerce
const TransactionModeECOMMERCE = "eCommerce" //"ECOMMERCE"

// TransactionModeMOBILE - Mobile Device
const TransactionModeMOBILE = "MOBILE"

// TransactionModeTABLET - Tablet
const TransactionModeTABLET = "TABLET"

// CardTypeVisa - Visa Card
const CardTypeVisa = "001"

// CardTypeMaster - MasterCard Card
const CardTypeMaster = "002"

// CardTypeAmex - Amex Card
const CardTypeAmex = "003"

// CardTypeDiscover - Discover Card
const CardTypeDiscover = "004"

// CardTypeDinersClub - Diners Club Card
const CardTypeDinersClub = "005"

// CardTypeJBC - JBC Card
const CardTypeJBC = "007"

// CardTypeMaestroUK - Maestro (UK Domestic) Card
const CardTypeMaestroUK = "024"

// CardTypeEncoded - Encoded account number
const CardTypeEncoded = "039"

// CardTypeMaestroInt - Maestro (International) Card
const CardTypeMaestroInt = "042"

// CardTypeEloInt - Elo Card
const CardTypeEloInt = "054"

// SetupPayerAuthRequestData - Data for a 3DS validation request
type SetupPayerAuthRequestData struct {
	ClientReferenceInformation *commons.ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	PaymentInformation         *commons.PaymentInformation         `json:"paymentInformation,omitempty"`
	TokenizedCard              *commons.TokenizedCard              `json:"tokenizedCard,omitempty"`
	FluidData                  *commons.FluidData                  `json:"fluidData,omitempty"`
	Customer                   *commons.Customer                   `json:"customer,omitempty"`
	ProcessingInformation      *commons.ProcessingInformation      `json:"processingInformation,omitempty"`
	TokenInformation           *commons.TokenInformation           `json:"tokenInformation,omitempty"`
}

// SetupPayerAuthResponseData - Data for a 3DS validation request
type SetupPayerAuthResponseData struct {
	ClientReferenceInformation        *commons.ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	PaymentInformation                *commons.PaymentInformation                `json:"paymentInformation,omitempty"`
	TokenizedCard                     *commons.TokenizedCard                     `json:"tokenizedCard,omitempty"`
	FluidData                         *commons.FluidData                         `json:"fluidData,omitempty"`
	Customer                          *commons.Customer                          `json:"customer,omitempty"`
	ProcessingInformation             *commons.ProcessingInformation             `json:"processingInformation,omitempty"`
	TokenInformation                  *commons.TokenInformation                  `json:"tokenInformation,omitempty"`
	ConsumerAuthenticationInformation *commons.ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
	ID                                *string                                    `json:"id,omitempty"`
	Status                            *string                                    `json:"status,omitempty"`
	SubmitTimeUtc                     *time.Time                                 `json:"submitTimeUtc,omitempty"`
}

// ValidationRequestData - Data for a 3DS validation request
type ValidationRequestData struct {
	ClientReferenceInformation        *commons.ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	OrderInformation                  *commons.OrderInformation                  `json:"orderInformation,omitempty"`
	PaymentInformation                *commons.PaymentInformation                `json:"paymentInformation,omitempty"`
	ConsumerAuthenticationInformation *commons.ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
}

// EnrollmentRequestData - Data for a 3DS enrolment request
type EnrollmentRequestData struct {
	ClientReferenceInformation        *commons.ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	OrderInformation                  *commons.OrderInformation                  `json:"orderInformation,omitempty"`
	PaymentInformation                *commons.PaymentInformation                `json:"paymentInformation,omitempty"`
	BuyerInformation                  *commons.BuyerInformation                  `json:"buyerInformation,omitempty"`
	ConsumerAuthenticationInformation *commons.ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
	DeviceInformation                 *DeviceInformation                         `json:"deviceInformation,omitempty"`
	MerchantInformation               *MerchantInformation                       `json:"merchantInformation,omitempty"`
}

// EnrollmentResponse - Enrollment response message
type EnrollmentResponse struct {
	ClientReferenceInformation        *commons.ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	ConsumerAuthenticationInformation *commons.ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
	ID                                *string                                    `json:"id,omitempty"`
	Status                            *string                                    `json:"status,omitempty"`
	SubmitTimeUtc                     *time.Time                                 `json:"submitTimeUtc,omitempty"`
}

// ValidationResponse - Validation response message
type ValidationResponse struct {
	ClientReferenceInformation        *commons.ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	ConsumerAuthenticationInformation *commons.ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
	ErrorInformation                  *commons.ErrorInformation                  `json:"errorInformation,omitempty"`
	ID                                *string                                    `json:"id,omitempty"`
	Status                            *string                                    `json:"status,omitempty"`
	SubmitTimeUtc                     *time.Time                                 `json:"submitTimeUtc,omitempty"`
}

// DeviceInformation - Devicefingerprint information
type DeviceInformation struct {
	IPAddress                    *string `json:"ipAddress,omitempty"`
	HTTPAcceptBrowserValue       *string `json:"httpAcceptBrowserValue,omitempty"`
	HTTPAcceptContent            *string `json:"httpAcceptContent,omitempty"`
	HTTPBrowserLanguage          *string `json:"httpBrowserLanguage,omitempty"`
	HTTPBrowserJavaEnabled       *string `json:"httpBrowserJavaEnabled,omitempty"`
	HTTPBrowserJavaScriptEnabled *string `json:"httpBrowserJavaScriptEnabled,omitempty"`
	HTTPBrowserColorDepth        *int    `json:"httpBrowserColorDepth,omitempty"`
	HTTPBrowserScreenHeight      *int    `json:"httpBrowserScreenHeight,omitempty"`
	HTTPBrowserScreenWidth       *int    `json:"httpBrowserScreenWidth,omitempty"`
	HTTPBrowserTimeDifference    *int    `json:"httpBrowserTimeDifference,omitempty"`
	USERAgentBrowserValue        *string `json:"userAgentBrowserValue,omitempty"`
}

// MerchantInformation - Merchant information data
type MerchantInformation struct {
	MerchantDescriptor *MerchantDescriptor `json:"merchantDescriptor,omitempty"`
	MerchantName       *string             `json:"merchantName,omitempty"`
}

// MerchantDescriptor - Merchant descriptor data
type MerchantDescriptor struct {
	Name *string `json:"name,omitempty"`
	URL  *string `json:"url,omitempty"`
}
