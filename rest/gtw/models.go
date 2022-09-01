package gtw

import (
	"time"

	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
)

// PAYMENT AUTHORIZATION //

// Payment - Authorization struct
type Payment struct {
	ClientReferenceInformation        *commons.ClientReferenceInformation        `json:"clientReferenceInformation,omitempty"`
	ID                                *string                                    `json:"id,omitempty"`
	ProcessingInformation             *commons.ProcessingInformation             `json:"processingInformation,omitempty"`
	IssuerInformation                 *IssuerInformation                         `json:"issuerInformation,omitempty"`
	PaymentInformation                *commons.PaymentInformation                `json:"paymentInformation,omitempty"`
	ProcessorInformation              *commons.ProcessorInformation              `json:"processorInformation,omitempty"`
	ReconciliationID                  *string                                    `json:"reconciliationId,omitempty"`
	ReversalInformation               *ReversalInformation                       `json:"reversalInformation,omitempty"`
	StatusInformation                 *StatusInformation                         `json:"statusInformation,omitempty"`
	OrderInformation                  *commons.OrderInformation                  `json:"orderInformation,omitempty"`
	Status                            *string                                    `json:"status,omitempty"`
	SubmitTimeUtc                     *time.Time                                 `json:"submitTimeUtc,omitempty"`
	VoidAmountDetails                 *commons.AmountDetails                     `json:"voidAmountDetails,omitempty"`
	PaymentAccountInformation         *PaymentAccountInformation                 `json:"paymentAccountInformation,omitempty"`
	BuyerInformation                  *commons.BuyerInformation                  `json:"buyerInformation,omitempty"`
	RecipientInformation              *RecipientInformation                      `json:"recipientInformation,omitempty"`
	DeviceInformation                 *commons.DeviceInformation                 `json:"deviceInformation,omitempty"`
	MerchantInformation               *commons.MerchantInformation               `json:"merchantInformation,omitempty"`
	AggregatorInformation             *AggregatorInformation                     `json:"aggregatorInformation,omitempty"`
	ConsumerAuthenticationInformation *commons.ConsumerAuthenticationInformation `json:"consumerAuthenticationInformation,omitempty"`
	PointOfSaleInformation            *commons.PointOfSaleInformation            `json:"pointOfSaleInformation,omitempty"`
	MerchantDefinedInformation        []MerchantDefinedInformation               `json:"merchantDefinedInformation,omitempty"`
	InstallmentInformation            *commons.InstallmentInformation            `json:"installmentInformation,omitempty"`
	Links                             *commons.Links                             `json:"_links,omitempty"`
	Reason                            *string                                    `json:"reason,omitempty"`
}

// ReversalInformation
type ReversalInformation struct {
	AmountDetails *commons.AmountDetails `json:"amountDetails,omitempty"`
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
	Card *commons.Card `json:"card,omitempty"`
}

// StatusInformation - Payment status information
type StatusInformation struct {
	Reason  *string `json:"reason,omitempty"`
	Message *string `json:"message,omitempty"`
}
