package rest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/akayna/Go-dreamBridgeCybersource/rest/authentication"
	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
)

var host = "apitest.cybersource.com" // Ambiente de teste
//var host = "api.cybersource.com" // Ambiente produtivo

// Functions

// RestFullSimplePOST - Execute a simple Post call to an endpoint
func RestFullSimplePOST(endpoint, payload string) (*RequestResponse, error) {
	url := "https://" + host + endpoint

	req, _ := http.NewRequest("POST", url, strings.NewReader(payload))

	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("cybersourcerest - RestFullSimplePOST - Error executing POST request.")
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var response = RequestResponse{
		StatusCode: res.StatusCode,
		Body:       string(body),
	}

	return &response, nil
}

// RestFullPOST - Execute a Post call to an endpoint
func RestFullPOST(credentials *commons.CyberSourceCredential, endpoint, payload string) (*RequestResponse, error) {
	url := "https://" + host + endpoint

	req, _ := http.NewRequest("POST", url, strings.NewReader(payload))

	header, err := authentication.GetHeader(credentials, host, payload, "POST", endpoint)
	if err != nil {
		log.Println("cybersourcerest - RestFullGET - Error generating Get headers.")
		return nil, err
	}

	headerMap := header.GetMapString()

	log.Println("Header: ")

	for key, val := range headerMap {
		if val != "" {
			req.Header.Add(key, val)
			log.Println(key + ": " + val)
		}
	}

	log.Println("Payload:")
	log.Println(payload)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("cybersourcerest - RestFullGET - Error executing GET request.")
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	log.Println(res)

	var response = RequestResponse{
		StatusCode: res.StatusCode,
		Body:       string(body),
	}

	return &response, nil
}

// RestFullDELETE - Execute a Delete call to an endpoint
func RestFullDELETE(credentials *commons.CyberSourceCredential, endpoint string) (*RequestResponse, error) {
	url := "https://" + host + endpoint

	req, _ := http.NewRequest("DELETE", url, nil)

	header, err := authentication.GetHeader(credentials, host, "", "DELETE", endpoint)
	if err != nil {
		log.Println("cybersourcerest - RestFullDELETE - Error generating headers.")
		return nil, err
	}

	headerMap := header.GetMapString()

	for key, val := range headerMap {
		if val != "" {
			req.Header.Add(key, val)
			//fmt.Println(key + ": " + val)
		}
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("cybersourcerest - RestFullDELETE - Error executing DELETE request.")
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var response = RequestResponse{
		StatusCode: res.StatusCode,
		Body:       string(body),
	}

	return &response, nil
}

// RestFullGET - Execute a Get call to an endpoint
func RestFullGET(credentials *commons.CyberSourceCredential, endpoint string) (*RequestResponse, error) {
	url := "https://" + host + endpoint

	log.Println("Get URL: " + url)

	req, _ := http.NewRequest("GET", url, nil)

	header, err := authentication.GetHeader(credentials, host, "", "GET", endpoint)
	if err != nil {
		log.Println("cybersourcerest - RestFullGET - Error generating Get headers.")
		return nil, err
	}

	headerMap := header.GetMapString()

	log.Println("REQUEST HEADERS")
	for key, val := range headerMap {
		if val != "" {
			req.Header.Add(key, val)
			fmt.Println(key + ": " + val)
		}
	}

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("cybersourcerest - RestFullGET - Error executing GET request.")
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var response = RequestResponse{
		StatusCode: res.StatusCode,
		Body:       string(body),
	}

	return &response, nil
}

// RestFullGETNoCerdentials - Execute a Get call to an endpoint without the credentials
func RestFullGETNoCerdentials(endpoint string) (*RequestResponse, error) {
	url := "https://" + host + endpoint

	log.Println("Get URL: " + url)

	req, _ := http.NewRequest("GET", url, nil)

	res, err := http.DefaultClient.Do(req)

	if err != nil {
		log.Println("cybersourcerest - RestFullGETNoCerdentials - Error executing GET request.")
		return nil, err
	}

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	var response = RequestResponse{
		StatusCode: res.StatusCode,
		Body:       string(body),
	}

	return &response, nil
}
