package dm

import (
	"time"

	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
)

// DMResponse - DM case responses
type DMResponse struct {
	ClientReferenceInformation *commons.ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	ID                         *string                             `json:"id,omitempty"`
	RiskInformation            *RiskInformation                    `json:"riskInformation,omitempty"`
	Status                     *string                             `json:"status,omitempty"`
	SubmitTimeUtc              *time.Time                          `json:"submitTimeUtc,omitempty"`
	ErrorInformation           *commons.ErrorInformation           `json:"errorInformation,omitempty"`
	PaymentInformation         *commons.PaymentInformation         `json:"paymentInformation,omitempty"`
}

// DMRequest -Decision Manager request data struct
type DMRequest struct {
	ClientReferenceInformation *commons.ClientReferenceInformation `json:"clientReferenceInformation,omitempty"`
	ProcessorInformation       *commons.ProcessorInformation       `json:"processorInformation,omitempty"`
	PaymentInformation         *commons.PaymentInformation         `json:"paymentInformation,omitempty"`
	OrderInformation           *commons.OrderInformation           `json:"orderInformation,omitempty"`
	BuyerInformation           *commons.BuyerInformation           `json:"buyerInformation,omitempty"`
	DeviceInformation          *commons.DeviceInformation          `json:"deviceInformation,omitempty"`
	CardVerification           *commons.CardVerification           `json:"cardVerification,omitempty"`
	RiskInformation            *RiskInformation                    `json:"riskInformation,omitempty"`
	TravelInformation          *TravelInformation                  `json:"travelInformation,omitempty"`
	MerchantDefinedInformation []MerchantDefinedInformation        `json:"merchantDefinedInformation,omitempty"`
}

// Score - Score information
type Score struct {
	Result      *string  `json:"result,omitempty"`
	FactorCodes []string `json:"factorCodes,omitempty"`
	ModelUsed   *string  `json:"modelUsed,omitempty"`
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
	LocalTime          *string                     `json:"localTime,omitempty"`
	InfoCodes          *InfoCodes                  `json:"infoCodes,omitempty"`
	Profile            *Profile                    `json:"profile,omitempty"`
	EventType          *string                     `json:"eventType,omitempty"`
	Rules              []Rules                     `json:"rules,omitempty"`
	PaymentInformation *commons.PaymentInformation `json:"paymentInformation,omitempty"`
	Providers          *Providers                  `json:"providers,omitempty"`
	CasePriority       *string                     `json:"casePriority,omitempty"`
	Score              *Score                      `json:"score,omitempty"`
	MarkingDetails     *MarkingDetails             `json:"markingDetails,omitempty"`
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

// MarkingDetails - Field for list management actions
type MarkingDetails struct {
	Notes      *string `json:"notes,omitempty"`
	Reason     *string `json:"reason,omitempty"`
	RecordName *string `json:"recordName,omitempty"`
	Action     *string `json:"action,omitempty"`
}
