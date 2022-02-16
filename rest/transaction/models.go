package transaction

import (
	"time"

	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
)

// Transaction - Transaction structure
type Transaction struct {
	ID                                *string                                    `json:"id,omitempty"`
	RootID                            *string                                    `json:"rootId,omitempty"` //-
	ReconciliationID                  *string                                    `json:"reconciliationId,omitempty"`
	SubmitTimeUtc                     *time.Time                                 `json:"submitTimeUtc,omitempty"`
	MerchantID                        *string                                    `json:"merchantId,omitempty"`             //-
	ApplicationInformation            *ApplicationInformation                    `json:"applicationInformation,omitempty"` //-
	BuyerInformation                  *commons.BuyerInformation                  `json:"buyerInformation,omitempty"`
	ClientReferenceInformation        *commons.ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	ConsumerAuthenticationInformation *commons.ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
	DeviceInformation                 *commons.DeviceInformation                 `json:"deviceInformation,omitempty"`
	InstallmentInformation            *commons.InstallmentInformation            `json:"installmentInformation,omitempty"`
	FraudMarkingInformation           *FraudMarkingInformation                   `json:"fraudMarkingInformation,omitempty"` //-
	MerchantInformation               *commons.MerchantInformation               `json:"merchantInformation,omitempty"`
	OrderInformation                  *commons.OrderInformation                  `json:"orderInformation,omitempty"`
	PaymentInformation                *commons.PaymentInformation                `json:"paymentInformation,omitempty"`
	ProcessingInformation             *commons.ProcessingInformation             `json:"processingInformation,omitempty"`
	ProcessorInformation              *commons.ProcessorInformation              `json:"processorInformation,omitempty"`
	PointOfSaleInformation            *commons.PointOfSaleInformation            `json:"pointOfSaleInformation,omitempty"`
	RiskInformation                   *RiskInformation                           `json:"riskInformation,omitempty"`   //-
	SenderInformation                 *SenderInformation                         `json:"senderInformation,omitempty"` //-
	Links                             *commons.Links                             `json:"_links,omitempty"`
}

// Applications - Applications
type Applications struct {
	Name             *string `json:"name,omitempty"`
	Status           *string `json:"status,omitempty"`
	ReasonCode       *string `json:"reasonCode,omitempty"`
	RCode            *string `json:"rCode,omitempty"`
	RFlag            *string `json:"rFlag,omitempty"`
	ReconciliationID *string `json:"reconciliationId,omitempty"`
	RMessage         *string `json:"rMessage,omitempty"`
	ReturnCode       *int    `json:"returnCode,omitempty"`
}

// ApplicationInformation - Application information
type ApplicationInformation struct {
	Status       *string        `json:"status,omitempty"`
	ReasonCode   *int           `json:"reasonCode,omitempty"`
	Applications []Applications `json:"applications,omitempty"`
}

// FraudMarkingInformation - Fraud marking information
type FraudMarkingInformation struct {
}

// AmountDetails - Amount details
type AmountDetails struct {
	TotalAmount      *string `json:"totalAmount,omitempty"`
	Currency         *string `json:"currency,omitempty"`
	TaxAmount        *string `json:"taxAmount,omitempty"`
	AuthorizedAmount *string `json:"authorizedAmount,omitempty"`
}

// ShippingDetails - Shipping details
type ShippingDetails struct {
}

// Score - Score information
type Score struct {
}

// RiskInformation - Risk information
type RiskInformation struct {
	Score Score `json:"score"`
}

// SenderInformation - Sender information
type SenderInformation struct {
}
