package commons

type ErrorInformation struct {
	Reason  *string `json:"reason,omitempty"`
	Message *string `json:"message,omitempty"`
}

// Credentials - Cybersource + cardinal credentials
type Credentials struct {
	CyberSourceCredential CyberSourceCredential `json:"cyberSourceCredential"`
	CardinalCredential    CardinalCredential    `json:"cardinalCredential"`
}

// CyberSourceCredential - Struct with the cybersource credentials information
type CyberSourceCredential struct {
	SharedSecretKey string `json:"sharedSecretKey"`
	ProfileID       string `json:"profileID"`
	APIKeyID        string `json:"apiKeyID"`
	MID             string `json:"mid"`
}

// CardinalCredential - Struct with the cardinal credentials information
type CardinalCredential struct {
	APIKeyID      string `json:"apiKeyID"`
	APIIdentifier string `json:"apiIdentifier"`
	OrgUnitID     string `json:"orgUnitID"`
	MID           string `json:"mid"`
}

// Common use structs

// PersonalIdentification - Personal ID information
type PersonalIdentification struct {
	Type     *string `json:"type,omitempty"`
	ID       *string `json:"id,omitempty"`
	IssuedBy *string `json:"issuedBy,omitempty"`
}

// BuyerInformation - Buyer informarion
type BuyerInformation struct {
	Username               *string                 `json:"username,omitempty"`
	MerchantCustomerID     *string                 `json:"merchantCustomerId,omitempty"`
	DateOfBirth            *string                 `json:"dateOfBirth,omitempty"`
	VatRegistrationNumber  *string                 `json:"vatRegistrationNumber,omitempty"`
	CompanyTaxID           *string                 `json:"companyTaxId,omitempty"`
	PersonalIdentification *PersonalIdentification `json:"personalIdentification,omitempty"`
	HashedPassword         *string                 `json:"hashedPassword,omitempty"`
	MobilePhone            *string                 `json:"mobilePhone,omitempty"`
	Currency               *string                 `json:"currency,omitempty"`
}

// Partner - Struct for partner information
type Partner struct {
	OriginalTransactionID *string `json:"originalTransactionId,omitempty"`
	DeveloperID           *string `json:"developerId,omitempty"`
	SolutionID            *string `json:"solutionId,omitempty"`
}

// ClientReferenceInformation - Struct for client reference information
type ClientReferenceInformation struct {
	Code               *string  `json:"code,omitempty"`
	ApplicationName    *string  `json:"applicationName,omitempty"`
	ApplicationVersion *string  `json:"applicationVersion,omitempty"`
	TransactionID      *string  `json:"transactionId,omitempty"`
	Comments           *string  `json:"comments,omitempty"`
	Partner            *Partner `json:"partner,omitempty"`
}

// ConsumerAuthenticationInformation - Authentication information (3DS)
type ConsumerAuthenticationInformation struct {
	Cavv                        *string `json:"cavv,omitempty"`
	CavvAlgorithm               *string `json:"cavvAlgorithm,omitempty"`
	EciRaw                      *string `json:"eciRaw,omitempty"`
	ParesStatus                 *string `json:"paresStatus,omitempty"`
	VeresEnrolled               *string `json:"veresEnrolled,omitempty"`
	Xid                         *string `json:"xid,omitempty"`
	UcafAuthenticationData      *string `json:"ucafAuthenticationData,omitempty"`
	UcafCollectionIndicator     *string `json:"ucafCollectionIndicator,omitempty"`
	AuthenticationTransactionID *string `json:"authenticationTransactionId,omitempty"`
	ACSUrl                      *string `json:"acsUrl,omitempty"`
	PAReq                       *string `json:"pareq,omitempty"`
	Eci                         *string `json:"eci,omitempty"`
	ProofXML                    *string `json:"proofXml,omitempty"`
	Token                       *string `json:"token,omitempty"`
	AuthenticationPath          *string `json:"authenticationPath,omitempty"`
	EcommerceIndicator          *string `json:"ecommerceIndicator,omitempty"`
	SpecificationVersion        *string `json:"specificationVersion,omitempty"`
	Indicator                   *string `json:"indicator,omitempty"`
	AuthenticationResult        *string `json:"authenticationResult,omitempty"`
	AuthenticationStatusMsg     *string `json:"authenticationStatusMsg,omitempty"`
	MCC                         *string `json:"mcc,omitempty"`
	MessageCategory             *string `json:"messageCategory,omitempty"`
	ProductCode                 *string `json:"productCode,omitempty"`
	TransactionMode             *string `json:"transactionMode,omitempty"`
	OverridePaymentMethod       *string `json:"overridePaymentMethod,omitempty"`
	ReferenceID                 *string `json:"referenceId,omitempty"`
	AcsWindowSize               *string `json:"acsWindowSize,omitempty"`
	DeviceChannel               *string `json:"deviceChannel,omitempty"`
	RequestorID                 *string `json:"requestorId,omitempty"`
	RequestorName               *string `json:"requestorName,omitempty"`
	ReturnUrl                   *string `json:"returnUrl,omitempty"`
	DeviceDataCollectionUrl     *string `json:"deviceDataCollectionUrl,omitempty"`
	AccessToken                 *string `json:"accessToken,omitempty"`
	StepUpUrl                   *string `json:"stepUpUrl,omitempty"`
}

// DeviceInformation - Puchase device information
type DeviceInformation struct {
	HostName             *string `json:"hostName,omitempty"`
	IPAddress            *string `json:"ipAddress,omitempty"`
	UserAgent            *string `json:"userAgent,omitempty"`
	CookiesAccepted      *string `json:"cookiesAccepted,omitempty"`
	FingerprintSessionID *string `json:"fingerprintSessionId,omitempty"`
	HTTPBrowserEmail     *string `json:"httpBrowserEmail,omitempty"`
}

// InstallmentInformation - Installment information
type InstallmentInformation struct {
	Amount               *string `json:"amount,omitempty"`
	Frequency            *string `json:"frequency,omitempty"`
	PlanType             *string `json:"planType,omitempty"`
	Sequence             *string `json:"sequence,omitempty"`
	TotalAmount          *string `json:"totalAmount,omitempty"`
	TotalCount           *string `json:"totalCount,omitempty"`
	FirstInstallmentDate *string `json:"firstInstallmentDate,omitempty"`
	InvoiceData          *string `json:"invoiceData,omitempty"`
	PaymentType          *string `json:"paymentType,omitempty"`
	EligibilityInquiry   *string `json:"eligibilityInquiry,omitempty"`
}

// MerchantDescriptor - Merchant information
type MerchantDescriptor struct {
	Name               *string `json:"name,omitempty"`
	AlternateName      *string `json:"alternateName,omitempty"`
	Contact            *string `json:"contact,omitempty"`
	Address1           *string `json:"address1,omitempty"`
	Locality           *string `json:"locality,omitempty"`
	Country            *string `json:"country,omitempty"`
	PostalCode         *string `json:"postalCode,omitempty"`
	AdministrativeArea *string `json:"administrativeArea,omitempty"`
}

// ServiceFeeDescriptor - Informatin about the service provider that is collecting the service fee
type ServiceFeeDescriptor struct {
	Name    *string `json:"name,omitempty"`
	Contact *string `json:"contact,omitempty"`
	State   *string `json:"state,omitempty"`
}

// MerchantInformation - Merchant information
type MerchantInformation struct {
	MerchantDescriptor          *MerchantDescriptor   `json:"merchantDescriptor,omitempty"`
	SalesOrganizationID         *string               `json:"salesOrganizationId,omitempty"`
	CategoryCode                *string               `json:"categoryCode,omitempty"`
	CategoryCodeDomestic        *string               `json:"categoryCodeDomestic,omitempty"`
	TaxID                       *string               `json:"taxId,omitempty"`
	VatRegistrationNumber       *string               `json:"vatRegistrationNumber,omitempty"`
	CardAcceptorReferenceNumber *string               `json:"cardAcceptorReferenceNumber,omitempty"`
	TransactionLocalDateTime    *string               `json:"transactionLocalDateTime,omitempty"`
	ServiceFeeDescriptor        *ServiceFeeDescriptor `json:"serviceFeeDescriptor,omitempty"`
}

// OrderInformation - Order information
type OrderInformation struct {
	AmountDetails   *AmountDetails   `json:"amountDetails,omitempty"`
	BillTo          *BillTo          `json:"billTo,omitempty"`
	ShipTo          *ShipTo          `json:"shipTo,omitempty"`
	LineItems       []LineItems      `json:"lineItems,omitempty"`
	InvoiceDetails  *InvoiceDetails  `json:"invoiceDetails,omitempty"`
	ShippingDetails *ShippingDetails `json:"shippingDetails,omitempty"`
	ReturnsAccepted *string          `json:"returnsAccepted,omitempty"`
}

// BillTo - Information about the customer
type BillTo struct {
	FirstName          *string `json:"firstName,omitempty"`
	LastName           *string `json:"lastName,omitempty"`
	MiddleName         *string `json:"middleName,omitempty"`
	NameSuffix         *string `json:"nameSuffix,omitempty"`
	Title              *string `json:"title,omitempty"`
	Company            *string `json:"company,omitempty"`
	Address1           *string `json:"address1,omitempty"`
	Address2           *string `json:"address2,omitempty"`
	Address3           *string `json:"address3,omitempty"`
	Address4           *string `json:"address4,omitempty"`
	Locality           *string `json:"locality,omitempty"`
	AdministrativeArea *string `json:"administrativeArea,omitempty"`
	PostalCode         *string `json:"postalCode,omitempty"`
	Country            *string `json:"country,omitempty"`
	District           *string `json:"district,omitempty"`
	BuildingNumber     *string `json:"buildingNumber,omitempty"`
	Email              *string `json:"email,omitempty"`
	PhoneNumber        *string `json:"phoneNumber,omitempty"`
	PhoneType          *string `json:"phoneType,omitempty"`
}

// ShipTo - Information about the recipient
type ShipTo struct {
	FirstName          *string `json:"firstName,omitempty"`
	LastName           *string `json:"lastName,omitempty"`
	Address1           *string `json:"address1,omitempty"`
	Address2           *string `json:"address2,omitempty"`
	Locality           *string `json:"locality,omitempty"`
	AdministrativeArea *string `json:"administrativeArea,omitempty"`
	PostalCode         *string `json:"postalCode,omitempty"`
	Country            *string `json:"country,omitempty"`
	District           *string `json:"district,omitempty"`
	BuildingNumber     *string `json:"buildingNumber,omitempty"`
	PhoneNumber        *string `json:"phoneNumber,omitempty"`
	Company            *string `json:"company,omitempty"`
}

// Surcharge - Surcharge information
type Surcharge struct {
	Amount      *string `json:"amount,omitempty"`
	Description *string `json:"description,omitempty"`
}

// AmexAdditionalAmounts - American Express Direct information
type AmexAdditionalAmounts struct {
	Code   *string `json:"code,omitempty"`
	Amount *string `json:"amount,omitempty"`
}

// TaxDetails - Taxes detail information
type TaxDetails struct {
	Type          *string `json:"type,omitempty"`
	Amount        *string `json:"amount,omitempty"`
	Rate          *string `json:"rate,omitempty"`
	Code          *string `json:"code,omitempty"`
	TaxID         *string `json:"taxId,omitempty"`
	Applied       *string `json:"applied,omitempty"`
	ExemptionCode *string `json:"exemptionCode,omitempty"`
}

// AmountDetails - Amount details information
type AmountDetails struct {
	TotalAmount             *string                `json:"totalAmount,omitempty"`
	AuthorizedAmount        *string                `json:"authorizedAmount,omitempty"`
	ReversedAmount          *string                `json:"reversedAmount,omitempty"`
	Currency                *string                `json:"currency,omitempty"`
	VoidAmount              *string                `json:"voidAmount,omitempty"`
	DiscountAmount          *string                `json:"discountAmount,omitempty"`
	DutyAmount              *string                `json:"dutyAmount,omitempty"`
	GratuityAmount          *string                `json:"gratuityAmount,omitempty"`
	TaxAmount               *string                `json:"taxAmount,omitempty"`
	NationalTaxIncluded     *string                `json:"nationalTaxIncluded,omitempty"`
	TaxAppliedAfterDiscount *string                `json:"taxAppliedAfterDiscount,omitempty"`
	TaxAppliedLevel         *string                `json:"taxAppliedLevel,omitempty"`
	TaxTypeCode             *string                `json:"taxTypeCode,omitempty"`
	FreightAmount           *string                `json:"freightAmount,omitempty"`
	ForeignAmount           *string                `json:"foreignAmount,omitempty"`
	ForeignCurrency         *string                `json:"foreignCurrency,omitempty"`
	ExchangeRate            *string                `json:"exchangeRate,omitempty"`
	ExchangeRateTimeStamp   *string                `json:"exchangeRateTimeStamp,omitempty"`
	Surcharge               *Surcharge             `json:"surcharge,omitempty"`
	SettlementAmount        *string                `json:"settlementAmount,omitempty"`
	SettlementCurrency      *string                `json:"settlementCurrency,omitempty"`
	AmexAdditionalAmounts   *AmexAdditionalAmounts `json:"amexAdditionalAmounts,omitempty"`
	TaxDetails              *TaxDetails            `json:"taxDetails,omitempty"`
	ServiceFeeAmount        *string                `json:"serviceFeeAmount,omitempty"`
	OriginalAmount          *string                `json:"originalAmount,omitempty"`
	OriginalCurrency        *string                `json:"originalCurrency,omitempty"`
	CashbackAmount          *string                `json:"cashbackAmount,omitempty"`
}

// Passenger - Information about the passenger
type Passenger struct {
	Type        *string `json:"type,omitempty"`
	Status      *string `json:"status,omitempty"`
	Phone       *string `json:"phone,omitempty"`
	FirstName   *string `json:"firstName,omitempty"`
	LastName    *string `json:"lastName,omitempty"`
	ID          *string `json:"id,omitempty"`
	Email       *string `json:"email,omitempty"`
	Nationality *string `json:"nationality,omitempty"`
}

// LineItems - Information about the items
type LineItems struct {
	ProductCode             *string     `json:"productCode,omitempty"`
	ProductName             *string     `json:"productName,omitempty"`
	ProductSku              *string     `json:"productSku,omitempty"`
	ProductRisk             *string     `json:"productRisk,omitempty"`
	DistributorProductSku   *string     `json:"distributorProductSku,omitempty"`
	Gift                    *string     `json:"gift,omitempty"`
	Quantity                *int        `json:"quantity,omitempty"`
	UnitPrice               *float32    `json:"unitPrice,omitempty"`
	UnitOfMeasure           *string     `json:"unitOfMeasure,omitempty"`
	TotalAmount             *string     `json:"totalAmount,omitempty"`
	TaxAmount               *float32    `json:"taxAmount,omitempty"`
	TaxRate                 *string     `json:"taxRate,omitempty"`
	TaxAppliedAfterDiscount *string     `json:"taxAppliedAfterDiscount,omitempty"`
	TaxStatusIndicator      *string     `json:"taxStatusIndicator,omitempty"`
	TaxTypeCode             *string     `json:"taxTypeCode,omitempty"`
	AmountIncludesTax       *string     `json:"amountIncludesTax,omitempty"`
	TypeOfSupply            *string     `json:"typeOfSupply,omitempty"`
	CommodityCode           *string     `json:"commodityCode,omitempty"`
	DiscountAmount          *string     `json:"discountAmount,omitempty"`
	DiscountApplied         *string     `json:"discountApplied,omitempty"`
	DiscountRate            *string     `json:"discountRate,omitempty"`
	InvoiceNumber           *string     `json:"invoiceNumber,omitempty"`
	TaxDetails              *TaxDetails `json:"taxDetails,omitempty"`
	FulfillmentType         *string     `json:"fulfillmentType,omitempty"`
	Weight                  *string     `json:"weight,omitempty"`
	WeightIdentifier        *string     `json:"weightIdentifier,omitempty"`
	WeightUnit              *string     `json:"weightUnit,omitempty"`
	ReferenceDataCode       *string     `json:"referenceDataCode,omitempty"`
	ReferenceDataNumber     *string     `json:"referenceDataNumber,omitempty"`
	Passenger               *Passenger  `json:"passenger,omitempty"`
}

// TransactionAdviceAddendum - Transaction Advice Addendum (TAA) fields
type TransactionAdviceAddendum struct {
	Data *string `json:"data,omitempty"`
}

// InvoiceDetails - Invoice details information
type InvoiceDetails struct {
	InvoiceNumber             *string                    `json:"invoiceNumber,omitempty"`
	BarcodeNumber             *string                    `json:"barcodeNumber,omitempty"`
	ExpirationDate            *string                    `json:"expirationDate,omitempty"`
	PurchaseOrderNumber       *string                    `json:"purchaseOrderNumber,omitempty"`
	PurchaseOrderDate         *string                    `json:"purchaseOrderDate,omitempty"`
	PurchaseContactName       *string                    `json:"purchaseContactName,omitempty"`
	Taxable                   *string                    `json:"taxable,omitempty"`
	VatInvoiceReferenceNumber *string                    `json:"vatInvoiceReferenceNumber,omitempty"`
	CommodityCode             *string                    `json:"commodityCode,omitempty"`
	MerchandiseCode           *string                    `json:"merchandiseCode,omitempty"`
	TransactionAdviceAddendum *TransactionAdviceAddendum `json:"transactionAdviceAddendum,omitempty"`
	ReferenceDataCode         *string                    `json:"referenceDataCode,omitempty"`
	ReferenceDataNumber       *string                    `json:"referenceDataNumber,omitempty"`
}

// ShippingDetails - Shipping datails
type ShippingDetails struct {
	GiftWrap           *string `json:"giftWrap,omitempty"`
	ShippingMethod     *string `json:"shippingMethod,omitempty"`
	ShipFromPostalCode *string `json:"shipFromPostalCode,omitempty"`
}

// Card - Card information
type Card struct {
	Suffix                *string `json:"suffix,omitempty"`
	Prefix                *string `json:"prefix,omitempty"`
	Number                *string `json:"number,omitempty"`
	ExpirationMonth       *string `json:"expirationMonth,omitempty"`
	ExpirationYear        *string `json:"expirationYear,omitempty"`
	Type                  *string `json:"type,omitempty"`
	UseAs                 *string `json:"useAs,omitempty"`
	SourceAccountType     *string `json:"sourceAccountType,omitempty"`
	SecurityCode          *string `json:"securityCode,omitempty"`
	SecurityCodeIndicator *string `json:"securityCodeIndicator,omitempty"`
	AccountEncoderID      *string `json:"accountEncoderId,omitempty"`
	IssueNumber           *string `json:"issueNumber,omitempty"`
	StartMonth            *string `json:"startMonth,omitempty"`
	StartYear             *string `json:"startYear,omitempty"`
	ProductName           *string `json:"productName,omitempty"`
	Bin                   *string `json:"bin,omitempty"`
}

// Invoice - Payment invoice information
type Invoice struct {
}

// AccountFeatures - Account fatures information
type AccountFeatures struct {
}

// TokenizedCard - Token information
type TokenizedCard struct {
	Number          *string `json:"number,omitempty"`
	ExpirationMonth *string `json:"expirationMonth,omitempty"`
	ExpirationYear  *string `json:"expirationYear,omitempty"`
	Type            *string `json:"type,omitempty"`
	Cryptogram      *string `json:"cryptogram,omitempty"`
	RequestorID     *string `json:"requestorId,omitempty"`
	TransactionType *string `json:"transactionType,omitempty"`
	AssuranceLevel  *string `json:"assuranceLevel,omitempty"`
	StorageMethod   *string `json:"storageMethod,omitempty"`
	SecurityCode    *string `json:"securityCode,omitempty"`
}

// TokenInformation - Token information
type TokenInformation struct {
	TransientToken *string `json:"transientToken,omitempty"`
}

// FluidData - Information fo the payment solution
type FluidData struct {
	KeySerialNumber *string `json:"keySerialNumber,omitempty"`
	Descriptor      *string `json:"descriptor,omitempty"`
	Value           *string `json:"value,omitempty"`
	Encoding        *string `json:"encoding,omitempty"`
}

// Customer - Customer information
type Customer struct {
	CustomerID *string `json:"customerId,omitempty"`
}

// Account - Information about the banck account
type Account struct {
	Type                      *string `json:"type,omitempty"`
	Number                    *string `json:"number,omitempty"`
	EncoderID                 *string `json:"encoderId,omitempty"`
	CheckNumber               *string `json:"checkNumber,omitempty"`
	CheckImageReferenceNumber *string `json:"checkImageReferenceNumber,omitempty"`
}

// Bank - Bank information
type Bank struct {
	Account       *Account `json:"account,omitempty"`
	RoutingNumber *string  `json:"routingNumber,omitempty"`
}

// PaymentType - Payment type information
type PaymentType struct {
	Name        *string `json:"name,omitempty"`
	Type        *string `json:"type,omitempty"`
	SubTypeName *string `json:"subTypeName,omitempty"`
	Method      *string `json:"method,omitempty"`
}

// PaymentInformation - Payment information
type PaymentInformation struct {
	Card                 *Card                 `json:"card,omitempty"`
	Invoice              *Invoice              `json:"invoice,omitempty"`
	AccountFeatures      *AccountFeatures      `json:"accountFeatures,omitempty"`
	TokenizedCard        *TokenizedCard        `json:"tokenizedCard,omitempty"`
	FluidData            *FluidData            `json:"fluidData,omitempty"`
	Customer             *Customer             `json:"customer,omitempty"`
	Bank                 *Bank                 `json:"bank,omitempty"`
	PaymentType          *PaymentType          `json:"paymentType,omitempty"`
	Scheme               *string               `json:"scheme,omitempty"`
	Bin                  *string               `json:"bin,omitempty"`
	AccountType          *string               `json:"accountType,omitempty"`
	Issuer               *string               `json:"issuer,omitempty"`
	BinCountry           *string               `json:"binCountry,omitempty"`
	InstrumentIdentifier *InstrumentIdentifier `json:"instrumentIdentifier,omitempty"`
}

// InstrumentIdentifier - Instrument identifier information
type InstrumentIdentifier struct {
	ID *string `json:"id,omitempty"`
}

// MerchantInitiatedTransaction - Information about the previous transaction.
type MerchantInitiatedTransaction struct {
	Reason                   *string `json:"reason,omitempty"`
	PreviousTransactionID    *string `json:"previousTransactionId,omitempty"`
	OriginalAuthorizedAmount *string `json:"originalAuthorizedAmount,omitempty"`
}

// Initiator - Information about who has initiated the transaction
type Initiator struct {
	Type                         *string                       `json:"type,omitempty"`
	CredentialStoredOnFile       *string                       `json:"credentialStoredOnFile,omitempty"`
	StoredCredentialUsed         *string                       `json:"storedCredentialUsed,omitempty"`
	MerchantInitiatedTransaction *MerchantInitiatedTransaction `json:"merchantInitiatedTransaction,omitempty"`
}

// AuthorizationOptions - Options to the authorization process
type AuthorizationOptions struct {
	AuthType                *string    `json:"authType,omitempty"`
	VerbalAuthCode          *string    `json:"verbalAuthCode,omitempty"`
	VerbalAuthTransactionID *string    `json:"verbalAuthTransactionId,omitempty"`
	AuthIndicator           *string    `json:"authIndicator,omitempty"`
	PartialAuthIndicator    *string    `json:"partialAuthIndicator,omitempty"`
	BalanceInquiry          *string    `json:"balanceInquiry,omitempty"`
	IgnoreAvsResult         *string    `json:"ignoreAvsResult,omitempty"`
	DeclineAvsFlags         *string    `json:"declineAvsFlags,omitempty"`
	IgnoreCvResult          *string    `json:"ignoreCvResult,omitempty"`
	Initiator               *Initiator `json:"initiator,omitempty"`
	BillPayment             *string    `json:"billPayment,omitempty"`
	BillPaymentType         *string    `json:"billPaymentType,omitempty"`
}

// CaptureOptions - Options to control the partuials capture
type CaptureOptions struct {
	CaptureSequenceNumber *string `json:"captureSequenceNumber,omitempty"`
	TotalCaptureCount     *string `json:"totalCaptureCount,omitempty"`
	DateToCapture         *string `json:"dateToCapture,omitempty"`
}

// RecurringOptions - Information about the recurrence.
type RecurringOptions struct {
	LoanPayment           *string `json:"loanPayment,omitempty"`
	FirstRecurringPayment *string `json:"firstRecurringPayment,omitempty"`
}

// BankTransferOptions - Used for bank transfer
type BankTransferOptions struct {
	DeclineAvsFlags     *string `json:"declineAvsFlags,omitempty"`
	SecCode             *string `json:"secCode,omitempty"`
	TerminalCity        *string `json:"terminalCity,omitempty"`
	TerminalState       *string `json:"terminalState,omitempty"`
	EffectiveDate       *string `json:"effectiveDate,omitempty"`
	PartialPaymentID    *string `json:"partialPaymentId,omitempty"`
	CustomerMemo        *string `json:"customerMemo,omitempty"`
	PaymentCategoryCode *string `json:"paymentCategoryCode,omitempty"`
	SettlementMethod    *string `json:"settlementMethod,omitempty"`
	FraudScreeningLevel *string `json:"fraudScreeningLevel,omitempty"`
}

// PurchaseOptions - Purchase information
type PurchaseOptions struct {
	IsElectronicBenefitsTransfer *string `json:"isElectronicBenefitsTransfer,omitempty"`
	Type                         *string `json:"type,omitempty"`
}

// ElectronicBenefitsTransfer - Information about the EBT transaction
type ElectronicBenefitsTransfer struct {
	Category            *string `json:"category,omitempty"`
	VoucherSerialNumber *string `json:"voucherSerialNumber,omitempty"`
}

// ProcessingInformation - Information to process the transaction
type ProcessingInformation struct {
	Capture                    *bool                       `json:"capture,omitempty"`
	ProcessorID                *string                     `json:"processorId,omitempty"`
	BusinessApplicationID      *string                     `json:"businessApplicationId,omitempty"`
	CommerceIndicator          *string                     `json:"commerceIndicator,omitempty"`
	PaymentSolution            *string                     `json:"paymentSolution,omitempty"`
	ReconciliationID           *string                     `json:"reconciliationId,omitempty"`
	LinkID                     *string                     `json:"linkId,omitempty"`
	PurchaseLevel              *string                     `json:"purchaseLevel,omitempty"`
	ReportGroup                *string                     `json:"reportGroup,omitempty"`
	VisaCheckoutID             *string                     `json:"visaCheckoutId,omitempty"`
	IndustryDataType           *string                     `json:"industryDataType,omitempty"`
	AuthorizationOptions       *AuthorizationOptions       `json:"authorizationOptions,omitempty"`
	CaptureOptions             *CaptureOptions             `json:"captureOptions,omitempty"`
	RecurringOptions           *RecurringOptions           `json:"recurringOptions,omitempty"`
	BankTransferOptions        *BankTransferOptions        `json:"bankTransferOptions,omitempty"`
	PurchaseOptions            *PurchaseOptions            `json:"purchaseOptions,omitempty"`
	ElectronicBenefitsTransfer *ElectronicBenefitsTransfer `json:"electronicBenefitsTransfer,omitempty"`
	ActionList                 []string                    `json:"actionList,omitempty"`
}

// Processor - Processor information
type Processor struct {
	Name *string `json:"name,omitempty"`
}

// Avs - Avs information
type Avs struct {
	Code    *string `json:"code,omitempty"`
	CodeRaw *string `json:"codeRaw,omitempty"`
}

// CardVerification - Card verification information
type CardVerification struct {
	ResultCode *string `json:"resultCode,omitempty"`
}

// AchVerification - Ach verification information
type AchVerification struct {
	ResultCodeRaw *string `json:"resultCodeRaw,omitempty"`
}

// ElectronicVerificationResults - Electronic verification results information
type ElectronicVerificationResults struct {
}

// ProcessorInformation - Processor information
type ProcessorInformation struct {
	Processor                     *Processor                     `json:"processor,omitempty"`
	ApprovalCode                  *string                        `json:"approvalCode,omitempty"`
	ResponseCode                  *string                        `json:"responseCode,omitempty"`
	Avs                           *Avs                           `json:"avs,omitempty"`
	CardVerification              *CardVerification              `json:"cardVerification,omitempty"`
	AchVerification               *AchVerification               `json:"achVerification,omitempty"`
	ElectronicVerificationResults *ElectronicVerificationResults `json:"electronicVerificationResults,omitempty"`
	EventStatus                   *string                        `json:"eventStatus,omitempty"`
}

// Emv - EMV data
type Emv struct {
	Tags                             *string `json:"tags,omitempty"`
	CardholderVerificationMethodUsed *string `json:"cardholderVerificationMethodUsed,omitempty"`
	CardSequenceNumber               *string `json:"cardSequenceNumber,omitempty"`
	Fallback                         *string `json:"fallback,omitempty"`
	FallbackCondition                *string `json:"fallbackCondition,omitempty"`
}

// PointOfSaleInformation - POS information
type PointOfSaleInformation struct {
	TerminalID                    *string `json:"terminalId,omitempty"`
	TerminalSerialNumber          *string `json:"terminalSerialNumber,omitempty"`
	LaneNumber                    *string `json:"laneNumber,omitempty"`
	CatLevel                      *string `json:"catLevel,omitempty"`
	EntryMode                     *string `json:"entryMode,omitempty"`
	TerminalCapability            *string `json:"terminalCapability,omitempty"`
	PinEntryCapability            *string `json:"pinEntryCapability,omitempty"`
	OperatingEnvironment          *string `json:"operatingEnvironment,omitempty"`
	Emv                           *Emv    `json:"emv,omitempty"`
	AmexCapnData                  *string `json:"amexCapnData,omitempty"`
	TrackData                     *string `json:"trackData,omitempty"`
	StoreAndForwardIndicator      *string `json:"storeAndForwardIndicator,omitempty"`
	CardholderVerificationMethod  *string `json:"cardholderVerificationMethod,omitempty"`
	TerminalInputCapability       *string `json:"terminalInputCapability,omitempty"`
	TerminalCardCaptureCapability *string `json:"terminalCardCaptureCapability,omitempty"`
	TerminalOutputCapability      *string `json:"terminalOutputCapability,omitempty"`
	TerminalPinCapability         *string `json:"terminalPinCapability,omitempty"`
	DeviceID                      *string `json:"deviceId,omitempty"`
	PinBlockEncodingFormat        *string `json:"pinBlockEncodingFormat,omitempty"`
	EncryptedPin                  *string `json:"encryptedPin,omitempty"`
	EncryptedKeySerialNumber      *string `json:"encryptedKeySerialNumber,omitempty"`
	PartnerSdkVersion             *string `json:"partnerSdkVersion,omitempty"`
}

// RelatedTransaction - Related transaction information
type RelatedTransaction struct {
	Href   *string `json:"href,omitempty"`
	Method *string `json:"method,omitempty"`
}

// Links - Request links information
type Links struct {
	Self                *RelatedTransaction  `json:"self,omitempty"`
	AuthReversal        *RelatedTransaction  `json:"authReversal,omitempty"`
	Capture             *RelatedTransaction  `json:"capture,omitempty"`
	Refund              *RelatedTransaction  `json:"refund,omitempty"`
	Void                *RelatedTransaction  `json:"void,omitempty"`
	PaymentInstrument   *RelatedTransaction  `json:"paymentInstruments,omitempty"`
	RelatedTransactions []RelatedTransaction `json:"relatedTransactions,omitempty"`
}
