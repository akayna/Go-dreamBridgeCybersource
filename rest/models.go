package rest

// RestfullHeader - Struct conteining all the data necessary to the post or put header
type RestfullHeader struct {
	MerchantID  string
	Date        string
	Host        string
	Digest      string
	Signature   *headerSignature
	ContentType string
	ProfileID   string
}

// headerSignature - Contains all the information to generate the signature field on the RESTFull request
type headerSignature struct {
	APIKey    string `json:"keyid"`     // Cybersource API Key
	Algorithm string `json:"algorithm"` // HmacSHA256
	Headers   string `json:"headers"`   // POST & PUT: "host date (request-target) digest v-c-merchant-id" / GET: "host date (request-target) v-c-merchant-id"
	Signature string `json:"signature"` // A Base64-encoded hash based on the headers value. Each header's name and its associated value are included in a string. This string is converted to a hash value (HMACSHA256) and Base64-encoded.
}

// RequestResponse - Strunct containg some request response field
type RequestResponse struct {
	StatusCode int
	Body       string
}
