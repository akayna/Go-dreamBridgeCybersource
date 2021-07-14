package gtw

import (
	"time"

	"github.com/rafaelcunha/Go-CyberSource/cybersourcecommons"
)

// PAYMENT AUTHORIZATION //

// Payment - Authorization struct
type Payment struct {
	ClientReferenceInformation        *cybersourcecommons.ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	ID                                *string                                               `json:"id,omitempty"`
	ProcessingInformation             *cybersourcecommons.ProcessingInformation             `json:"processingInformation,omitempty"`
	IssuerInformation                 *IssuerInformation                                    `json:"issuerInformation,omitempty"`
	PaymentInformation                *cybersourcecommons.PaymentInformation                `json:"paymentInformation,omitempty"`
	ProcessorInformation              *cybersourcecommons.ProcessorInformation              `json:"processorInformation,omitempty"`
	ReconciliationID                  *string                                               `json:"reconciliationId,omitempty"`
	ReversalAmountDetails             *cybersourcecommons.AmountDetails                     `json:"reversalAmountDetails,omitempty"`
	StatusInformation                 *StatusInformation                                    `json:"statusInformation,omitempty"`
	OrderInformation                  *cybersourcecommons.OrderInformation                  `json:"orderInformation,omitempty"`
	Status                            *string                                               `json:"status,omitempty"`
	SubmitTimeUtc                     *time.Time                                            `json:"submitTimeUtc,omitempty"`
	VoidAmountDetails                 *cybersourcecommons.AmountDetails                     `json:"voidAmountDetails,omitempty"`
	PaymentAccountInformation         *PaymentAccountInformation                            `json:"paymentAccountInformation,omitempty"`
	BuyerInformation                  *cybersourcecommons.BuyerInformation                  `json:"buyerInformation,omitempty"`
	RecipientInformation              *RecipientInformation                                 `json:"recipientInformation,omitempty"`
	DeviceInformation                 *cybersourcecommons.DeviceInformation                 `json:"deviceInformation,omitempty"`
	MerchantInformation               *cybersourcecommons.MerchantInformation               `json:"merchantInformation,omitempty"`
	AggregatorInformation             *AggregatorInformation                                `json:"aggregatorInformation,omitempty"`
	ConsumerAuthenticationInformation *cybersourcecommons.ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
	PointOfSaleInformation            *cybersourcecommons.PointOfSaleInformation            `json:"pointOfSaleInformation,omitempty"`
	MerchantDefinedInformation        []MerchantDefinedInformation                          `json:"merchantDefinedInformation,omitempty"`
	InstallmentInformation            *cybersourcecommons.InstallmentInformation            `json:"installmentInformation,omitempty"`
	Links                             *cybersourcecommons.Links                             `json:"_links,omitempty"`
}

// IssuerInformation - Additional data about the issuer
type IssuerInformation struct {
	DiscretionaryData *string `json:"discretionaryData,omitempty"`
}

// RecipientInformation - Recipent´s account information
type RecipientInformation struct {
	AccountID  *string `json:"accountId,omitempty"`
	LastName   *string `json:"lastName,omitempty"`
	PostalCode *string `json:"postalCode,omitempty"`
}

// SubMerchant - SubMerchant information
type SubMerchant struct {
	CardAcceptorID     *string `json:"cardAcceptorId,omitempty"`
	Name               *string `json:"name,omitempty"`
	Address1           *string `json:"address1,omitempty"`
	Locality           *string `json:"locality,omitempty"`
	AdministrativeArea *string `json:"administrativeArea,omitempty"`
	Region             *string `json:"region,omitempty"`
	PostalCode         *string `json:"postalCode,omitempty"`
	Country            *string `json:"country,omitempty"`
	Email              *string `json:"email,omitempty"`
	PhoneNumber        *string `json:"phoneNumber,omitempty"`
}

// AggregatorInformation - Payment aggregator information
type AggregatorInformation struct {
	AggregatorID *string      `json:"aggregatorId,omitempty"`
	Name         *string      `json:"name,omitempty"`
	SubMerchant  *SubMerchant `json:"subMerchant,omitempty"`
}

// MerchantDefinedInformation - Merchant-defined information
type MerchantDefinedInformation struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// PaymentAccountInformation - Payment account information
type PaymentAccountInformation struct {
	Card *cybersourcecommons.Card `json:"card,omitempty"`
}

// StatusInformation - Payment status information
type StatusInformation struct {
	Reason  *string `json:"reason,omitempty"`
	Message *string `json:"message,omitempty"`
}
