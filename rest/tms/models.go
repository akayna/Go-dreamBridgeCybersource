package tms

import (
	"github.com/rafaelcunha/Go-CyberSource/cybersourcecommons"
)

// CreateInstrumentIdentifierRequest - Contains all fields to request an Instrument Identifier Token
type CreateInstrumentIdentifierRequest struct {
	Card *cybersourcecommons.Card `json:"card,omitempty"`
}

// InstrumentIdentifierResponse - Contains all response fields of an instrument identifier creation request
type InstrumentIdentifierResponse struct {
	Links                 *cybersourcecommons.Links                 `json:"_links,omitempty"`
	ID                    *string                                   `json:"id,omitempty"`
	Object                *string                                   `json:"object,omitempty"`
	State                 *string                                   `json:"state,omitempty"`
	Card                  *cybersourcecommons.Card                  `json:"card,omitempty"`
	ProcessingInformation *cybersourcecommons.ProcessingInformation `json:"processingInformation,omitempty"`
	Metadata              *Metadata                                 `json:"metadata,omitempty"`
}

// PaymentInstrumentResponse - COntains all field into the Payment Instrument Get Response
type PaymentInstrumentResponse struct {
	Links    *cybersourcecommons.Links `json:"_links,omitempty"`
	ID       *string                   `json:"id,omitempty"`
	Object   *string                   `json:"object,omitempty"`
	State    *string                   `json:"state,omitempty"`
	Card     *cybersourcecommons.Card  `json:"card,omitempty"`
	Metadata *Metadata                 `json:"metadata,omitempty"`
	Embedded *Embedded                 `json:"_embedded,omitempty"`
}

// CreatePaymentInstrumentRequest - Contains all possible fields to create a payment instrument
type CreatePaymentInstrumentRequest struct {
	BankAccount          *BankAccount                         `json:"bankAccount,omitempty"`
	Card                 *cybersourcecommons.Card             `json:"card,omitempty"`
	BuyerInformation     *cybersourcecommons.BuyerInformation `json:"buyerInformation,omitempty"`
	BillTo               *cybersourcecommons.BillTo           `json:"billTo,omitempty"`
	InstrumentIdentifier *InstrumentIdentifier                `json:"instrumentIdentifier,omitempty"`
}

// Embedded - Embedded info structure
type Embedded struct {
	InstrumentIdentifier *InstrumentIdentifier `json:"instrumentIdentifier,omitempty"`
}

// InstrumentIdentifier - InstrumentIdentifier infos struct
type InstrumentIdentifier struct {
	Links                 *cybersourcecommons.Links                 `json:"_links,omitempty"`
	ID                    *string                                   `json:"id,omitempty"`
	Object                *string                                   `json:"object,omitempty"`
	State                 *string                                   `json:"state,omitempty"`
	Card                  *cybersourcecommons.Card                  `json:"card,omitempty"`
	ProcessingInformation *cybersourcecommons.ProcessingInformation `json:"processingInformation,omitempty"`
	Metadata              *Metadata                                 `json:"metadata,omitempty"`
}

// Metadata - Metadata information
type Metadata struct {
	Creator *string `json:"creator,omitempty"`
}

// BankAccount - Struct with information about the bank account.
type BankAccount struct {
	Type string `json:"type,omitempty"`
}