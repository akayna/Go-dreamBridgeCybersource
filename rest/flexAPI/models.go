package flexAPI

// generateCardTokenRequest - Struct with the necessary information to execute a generate card request.
type generateCardTokenRequest struct {
	KeyID    string    `json:"keyId"`
	CardInfo *CardInfo `json:"cardInfo"`
}

// CardInfo - Card information to a Token Generate Request
type CardInfo struct {
	CardNumber          string `json:"cardNumber"`
	CardExpirationMonth string `json:"cardExpirationMonth"`
	CardExpirationYear  string `json:"cardExpirationYear"`
	CardType            string `json:"cardType"`
}

// GenerateCardTokenResponse - Struct with the response of a Generate Card Token Request
type GenerateCardTokenResponse struct {
	KeyID                string               `json:"keyId"`
	Token                string               `json:"token"`
	MaskedPan            string               `json:"maskedPan"`
	CardType             string               `json:"cardType"`
	Timestamp            int64                `json:"timestamp"`
	SignedFields         string               `json:"signedFields"`
	Signature            string               `json:"signature"`
	DiscoverableServices DiscoverableServices `json:"discoverableServices"`
	Embedded             Embedded             `json:"_embedded"`
}

// Embedded - Not specified into the Cybersource documentations
type Embedded struct {
	IcsReply IcsReply `json:"icsReply"`
}

// IcsReply - Not specified into the Cybersource documentations
type IcsReply struct {
	RequestID            string               `json:"requestId"`
	InstrumentIdentifier InstrumentIdentifier `json:"instrumentIdentifier"`
	Links                Links                `json:"_links"`
}

// InstrumentIdentifier - Not specified into the Cybersource documentations
type InstrumentIdentifier struct {
	ID    string `json:"id"`
	New   string `json:"new"`
	State string `json:"state"`
}

// Self - Not specified into the Cybersource documentations
type Self struct {
	Href string `json:"href"`
}

// Links - Not specified into the Cybersource documentations
type Links struct {
	Self Self `json:"self"`
}

// DiscoverableServices - Object specified into the Cybersource documentations
type DiscoverableServices struct {
}

// GenerateKeyRequest - Struct for generate key request
type GenerateKeyRequest struct {
	Encryptiontype *string `json:"encryptionType,omitempty"`
	Targetorigin   *string `json:"targetOrigin,omitempty"`
}

// GenerateKeyResponse - Struct with the generate key response information.
type GenerateKeyResponse struct {
	KeyID *string `json:"keyId,omitempty"`
	Der   *Der    `json:"der,omitempty"`
	Jwk   *Jwk    `json:"jwk,omitempty"`
}

// Der - Struct with the generated key information
type Der struct {
	Format    string `json:"format"`
	Algorithm string `json:"algorithm"`
	PublicKey string `json:"publicKey"`
}

// Jwk - Struct with the encrypted generated key
type Jwk struct {
	Kty string `json:"kty"`
	Use string `json:"use"`
	Kid string `json:"kid"`
	N   string `json:"n"`
	E   string `json:"e"`
}
