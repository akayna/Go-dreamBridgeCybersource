package dm

import (
	"time"

	"github.com/rafaelcunha/Go-CyberSource/cybersourcecommons"
)

// ClientReferenceInformation - ClientReferenceInformation infos struct
type ClientReferenceInformation struct {
	Code string `json:"code"`
}

// CreateCaseResponse - Create DM case response
type CreateCaseResponse struct {
	ClientReferenceInformation *cybersourcecommons.ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	ID                         *string                                        `json:"id,omitempty"`
	RiskInformation            *RiskInformation                               `json:"riskInformation,omitempty"`
	Status                     *string                                        `json:"status,omitempty"`
	SubmitTimeUtc              *time.Time                                     `json:"submitTimeUtc,omitempty"`
	ErrorInformation           *ErrorInformation                              `json:"errorInformation,omitempty"`
	PaymentInformation         *cybersourcecommons.PaymentInformation         `json:"paymentInformation,omitempty"`
}

// Score - Score information
type Score struct {
	Result      *string  `json:"result,omitempty"`
	FactorCodes []string `json:"factorCodes,omitempty"`
	ModelUsed   *string  `json:"modelUsed,omitempty"`
}

// ErrorInformation - Error information
type ErrorInformation struct {
	Reason  *string `json:"reason,omitempty"`
	Message *string `json:"message,omitempty"`
}

// CreateCaseRequest - Create Decision Manager Case request data struct
type CreateCaseRequest struct {
	ClientReferenceInformation *cybersourcecommons.ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	ProcessorInformation       *cybersourcecommons.ProcessorInformation       `json:"processorInformation,omitempty"`
	PaymentInformation         *cybersourcecommons.PaymentInformation         `json:"paymentInformation,omitempty"`
	OrderInformation           *cybersourcecommons.OrderInformation           `json:"orderInformation,omitempty"`
	BuyerInformation           *cybersourcecommons.BuyerInformation           `json:"buyerInformation,omitempty"`
	DeviceInformation          *cybersourcecommons.DeviceInformation          `json:"deviceInformation,omitempty"`
	CardVerification           *cybersourcecommons.CardVerification           `json:"cardVerification,omitempty"`
	RiskInformation            *RiskInformation                               `json:"riskInformation,omitempty"`
	TravelInformation          *TravelInformation                             `json:"travelInformation,omitempty"`
	MerchantDefinedInformation []MerchantDefinedInformation                   `json:"merchantDefinedInformation,omitempty"`
}

// InfoCodes - Info codes
type InfoCodes struct {
	Address        []string `json:"address,omitempty"`
	Phone          []string `json:"phone,omitempty"`
	GlobalVelocity []string `json:"globalVelocity,omitempty"`
	Suspicious     []string `json:"suspicious,omitempty"`
	IdentityChange []string `json:"identityChange,omitempty"`
}

// Profile - Profile name
type Profile struct {
	Name         *string `json:"name,omitempty"`
	SelectorRule *string `json:"selectorRule,omitempty"`
}

// Rules - Triggered rule and its decision.
type Rules struct {
	Decision *string `json:"decision,omitempty"`
	Name     *string `json:"name,omitempty"`
}

// Providers - ?
type Providers struct {
}

// RiskInformation - Risk information for or from the Decision Manager
type RiskInformation struct {
	//Score              *string                                `json:"score,omitempty"`
	LocalTime          *string                                `json:"localTime,omitempty"`
	InfoCodes          *InfoCodes                             `json:"infoCodes,omitempty"`
	Profile            *Profile                               `json:"profile,omitempty"`
	EventType          *string                                `json:"eventType,omitempty"`
	Rules              []Rules                                `json:"rules,omitempty"`
	PaymentInformation *cybersourcecommons.PaymentInformation `json:"paymentInformation,omitempty"`
	Providers          *Providers                             `json:"providers,omitempty"`
	CasePriority       *string                                `json:"casePriority,omitempty"`
	Score              *Score                                 `json:"score,omitempty"`
}

// Legs - Travel lags information
type Legs struct {
	Origination       *string `json:"origination,omitempty"`
	Destination       *string `json:"destination,omitempty"`
	DepartureDateTime *string `json:"departureDateTime,omitempty"`
}

// TravelInformation - Information about the travel
type TravelInformation struct {
	ActualFinalDestination *string `json:"actualFinalDestination,omitempty"`
	CompleteRoute          *string `json:"completeRoute,omitempty"`
	DepartureTime          *string `json:"departureTime,omitempty"`
	JourneyType            *string `json:"journeyType,omitempty"`
	Legs                   *Legs   `json:"legs,omitempty"`
}

// MerchantDefinedInformation - Free field for the merchant to send aditional data
type MerchantDefinedInformation struct {
	Key   *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}
