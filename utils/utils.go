package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"log"
)

// GenerateSignature - Generate the CyberSource signature using the key and data
func GenerateSignature(key, data string) (string, error) {

	//log.Println("Key base 64 encoded: " + key)
	decodedKey, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		log.Println("digest - GenerateSignature: Error decoding key from Base64.")
		log.Println("error:", err)
		return "", err
	}

	//log.Println("Key base 64 decoded: " + string(decodedKey))

	hasher := hmac.New(sha256.New, decodedKey)
	hasher.Write([]byte(data))

	signatureEncodedSTD := base64.StdEncoding.EncodeToString(hasher.Sum(nil))

	log.Println("Calculated Signature econded:", signatureEncodedSTD)

	return signatureEncodedSTD, nil
}
