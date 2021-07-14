package 3ds

import (
	"time"

	"github.com/rafaelcunha/Go-CyberSource/cybersourcecommons"
)

// TransactionModeMOTO - Mail Order Telephone Order
const TransactionModeMOTO = "MOTO"

// TransactionModeRETAIL - Retail
const TransactionModeRETAIL = "RETAIL"

// TransactionModeECOMMERCE - eCommerce
const TransactionModeECOMMERCE = "ECOMMERCE"

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

// ValidationRequestData - Data for a 3DS validation request
type ValidationRequestData struct {
	ClientReferenceInformation        *cybersourcecommons.ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	OrderInformation                  *cybersourcecommons.OrderInformation                  `json:"orderInformation,omitempty"`
	PaymentInformation                *cybersourcecommons.PaymentInformation                `json:"paymentInformation,omitempty"`
	ConsumerAuthenticationInformation *cybersourcecommons.ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
}

// EnrollmentRequestData - Data for a 3DS enrolment request
type EnrollmentRequestData struct {
	ClientReferenceInformation        *cybersourcecommons.ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	OrderInformation                  *cybersourcecommons.OrderInformation                  `json:"orderInformation,omitempty"`
	PaymentInformation                *cybersourcecommons.PaymentInformation                `json:"paymentInformation,omitempty"`
	BuyerInformation                  *cybersourcecommons.BuyerInformation                  `json:"buyerInformation,omitempty"`
	ConsumerAuthenticationInformation *cybersourcecommons.ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
	DeviceInformation                 *DeviceInformation                                    `json:"deviceInformation,omitempty"`
	MerchantInformation               *MerchantInformation                                  `json:"merchantInformation,omitempty"`
}

// EnrollmentResponse - Enrollment response message
type EnrollmentResponse struct {
	ClientReferenceInformation        *cybersourcecommons.ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	ConsumerAuthenticationInformation *cybersourcecommons.ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
	ID                                *string                                               `json:"id,omitempty"`
	Status                            *string                                               `json:"status,omitempty"`
	SubmitTimeUtc                     *time.Time                                            `json:"submitTimeUtc,omitempty"`
}

// ValidationResponse - Validation response message
type ValidationResponse struct {
	ClientReferenceInformation        *cybersourcecommons.ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	ConsumerAuthenticationInformation *cybersourcecommons.ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
	ID                                *string                                               `json:"id,omitempty"`
	Status                            *string                                               `json:"status,omitempty"`
	SubmitTimeUtc                     *time.Time                                            `json:"submitTimeUtc,omitempty"`
}

// DeviceInformation - Devicefingerprint information
type DeviceInformation struct {
	IPAddress                    *string `json:"ipAddress,omitempty"`
	HTTPAcceptBrowserValue       *string `json:"httpAcceptBrowserValue,omitempty"`
	HTTPAcceptContent            *string `json:"httpAcceptContent,omitempty"`
	HTTPBrowserLanguage          *string `json:"httpBrowserLanguage,omitempty"`
	HTTPBrowserJavaEnabled       *string `json:"httpBrowserJavaEnabled,omitempty"`
	HTTPBrowserJavaScriptEnabled *string `json:"httpBrowserJavaScriptEnabled,omitempty"`
	HTTPBrowserColorDepth        *string `json:"httpBrowserColorDepth,omitempty"`
	HTTPBrowserScreenHeight      *string `json:"httpBrowserScreenHeight,omitempty"`
	HTTPBrowserScreenWidth       *string `json:"httpBrowserScreenWidth,omitempty"`
	HTTPBrowserTimeDifference    *string `json:"httpBrowserTimeDifference,omitempty"`
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
