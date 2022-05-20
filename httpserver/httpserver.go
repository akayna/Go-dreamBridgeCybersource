package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync/atomic"
	"syscall"
	"time"

	"github.com/akayna/Go-dreamBridgeCybersource/rest/commons"
	"github.com/akayna/Go-dreamBridgeCybersource/rest/flexAPI"
	"github.com/akayna/Go-dreamBridgeCybersource/rest/microform"
	"github.com/akayna/Go-dreamBridgeCybersource/rest/threeds"
	"github.com/akayna/Go-dreamBridgeUtils/jsonfile"
)

type middleware func(http.Handler) http.Handler
type middlewares []middleware

func (mws middlewares) apply(hdlr http.Handler) http.Handler {
	if len(mws) == 0 {
		return hdlr
	}
	return mws[1:].apply(mws[0](hdlr))
}

func (c *controller) shutdown(ctx context.Context, server *http.Server) context.Context {
	ctx, done := context.WithCancel(ctx)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		defer done()

		<-quit
		signal.Stop(quit)
		close(quit)

		atomic.StoreInt64(&c.healthy, 0)
		server.ErrorLog.Printf("Server is shutting down...\n")

		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			server.ErrorLog.Fatalf("Could not gracefully shutdown the server: %s\n", err)
		}
	}()

	return ctx
}

type controller struct {
	logger        *log.Logger
	nextRequestID func() string
	healthy       int64
}

var credentials commons.Credentials

func main() {

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Printf("Loading credentials...")

	// Credenciais teste
	err := jsonfile.ReadJSONFile2("/home/rafaelsonhador/Documents/Credenciais Cybersource/", "rafaelcunha.json", &credentials)

	// Credenciais live
	//err := jsonfile.ReadJSONFile2("/home/rafaelsonhador/Documents/Credenciais Cybersource/", "cybsbrdemo.json", &credentials)

	if err != nil {
		log.Println("Erro ao ler credenciais.")
		log.Println("Erro: ", err)

		return
	}

	logger.Println("Credentials loaded: " + credentials.CyberSourceCredential.MID)

	logger.Printf("Server is starting...")

	listenAddr := ":5000"
	if len(os.Args) == 2 {
		listenAddr = os.Args[1]
	}

	c := &controller{logger: logger, nextRequestID: func() string { return strconv.FormatInt(time.Now().UnixNano(), 36) }}
	router := http.NewServeMux()

	//router.HandleFunc("/", c.index)
	router.HandleFunc("/healthz", c.healthz)

	// FlexAPI services
	router.HandleFunc("/getFlexAPIKey", getFlexAPIKey)
	router.HandleFunc("/getFlexAPIKeyCrypto", getFlexAPIKeyCrypto)

	// Microform services
	router.HandleFunc("/getMicroformContext", getMicroformContext)
	router.HandleFunc("/validateMicroformToken", validateMicroformToken)

	// 3DS
	router.HandleFunc("/setupPayerAuth", setupPayerAuth)
	router.HandleFunc("/doEnrollment", doEnrollment)
	router.HandleFunc("/validate", validate)

	directory := flag.String("d", "./", "the directory of static file to host")
	router.Handle("/", http.StripPrefix(strings.TrimRight("/", "/"), http.FileServer(http.Dir(*directory))))

	flag.Parse()

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      (middlewares{c.tracing, c.logging}).apply(router),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}
	ctx := c.shutdown(context.Background(), server)

	logger.Printf("Server is ready to handle requests at %q\n", listenAddr)
	atomic.StoreInt64(&c.healthy, time.Now().UnixNano())

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %q: %s\n", listenAddr, err)
	}
	<-ctx.Done()
	logger.Printf("Server stopped\n")
}

func (c *controller) index(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	fmt.Fprintf(w, "Olá, Mundo!\n")
}

func (c *controller) healthz(w http.ResponseWriter, req *http.Request) {
	if h := atomic.LoadInt64(&c.healthy); h == 0 {
		w.WriteHeader(http.StatusServiceUnavailable)
	} else {
		fmt.Fprintf(w, "uptime: %s\n", time.Since(time.Unix(0, h)))
	}
}

func (c *controller) logging(hdlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		defer func(start time.Time) {
			requestID := w.Header().Get("X-Request-Id")
			if requestID == "" {
				requestID = "unknown"
			}
			c.logger.Println(requestID, req.Method, req.URL.Path, req.RemoteAddr, req.UserAgent(), time.Since(start))
		}(time.Now())
		hdlr.ServeHTTP(w, req)
	})
}

func (c *controller) tracing(hdlr http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		requestID := req.Header.Get("X-Request-Id")
		if requestID == "" {
			requestID = c.nextRequestID()
		}
		w.Header().Set("X-Request-Id", requestID)
		hdlr.ServeHTTP(w, req)
	})
}

// getFlexAPIKey - Generate ont key for the FlexAPI and send it to the browser
func getFlexAPIKey(w http.ResponseWriter, req *http.Request) {
	log.Println("getFlexAPIKey")

	generatedKey, msg, err := flexAPI.GenerateKey(&credentials.CyberSourceCredential, nil)

	if err != nil {
		log.Println("main - Error generating key.")
		log.Println(err)
		return
	}

	log.Println(msg)
	log.Printf("Key: %+v\n", generatedKey)
	log.Printf("KeyID: %+v\n", *generatedKey.KeyID)

	w.Write([]byte(*generatedKey.KeyID))
}

// getFlexAPIKey - Generate ont key for the FlexAPI and send it to the browser
func getFlexAPIKeyCrypto(w http.ResponseWriter, req *http.Request) {
	var origin = "https://www.teste.com"
	generatedKey, msg, err := flexAPI.GenerateRsaOaep256Key(&credentials.CyberSourceCredential, &origin)

	if err != nil {
		log.Println("main - Error generating key.")
		log.Println(err)
		return
	}

	fmt.Println(msg)
	fmt.Printf("Key: %+v\n", *generatedKey)
	fmt.Printf("KeyID: %+v\n", *generatedKey.KeyID)
	fmt.Printf("JWK: %+v\n", *generatedKey.Jwk)

	responseJSON, err := json.Marshal(generatedKey)

	if err != nil {
		errorString := "getFlexAPIKeyCrypto - Erro converting struct to JSON - " + err.Error()
		log.Printf("getFlexAPIKeyCrypto - Erro converting struct to JSON - %s", err)
		w.Write([]byte(errorString))
		return
	}

	fmt.Println("Json: ", string(responseJSON))

	w.Write([]byte(responseJSON))
}

// getMicroformKey - Generate one microfom key and send back to the frontend
func getMicroformContext(w http.ResponseWriter, req *http.Request) {
	log.Println("getMicroformContext")

	var targetOrigin []string

	targetOrigin = append(targetOrigin, "http://localhost:5000")

	context, message, err := microform.GenerateMicroformContext(&credentials.CyberSourceCredential, targetOrigin)

	if err != nil {
		log.Println("getMicroformContext - Error getting context.")
		log.Println("getMicroformContext - Message: ", message)
		log.Println("Error: ", err)
		http.Error(w, message, http.StatusInternalServerError)
		return
	}

	w.Write([]byte(context))
}

func validateMicroformToken(w http.ResponseWriter, req *http.Request) {
	log.Println("validateMicroformToken")
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		log.Println("main.validateMicroformToken - Error reading POST request body.")
		w.Write([]byte("Error reading POST request body."))
		return
	}

	validated, cyberToken, err := microform.ValidateToken(string(body))

	if err != nil {
		log.Println("main.validateMicroformToken - Error validating token received.")
		log.Println(err)
		w.Write([]byte(cyberToken))
	}

	if validated {
		log.Println("main.validateMicroformToken - Token validated.")
		log.Println("main.validateMicroformToken - Token Cybersource: ", cyberToken)
		w.Write([]byte("Token validated: " + cyberToken))
	} else {
		w.Write([]byte("Token not validated."))
	}
}

// setupPayerAuth - Execute the setup payer auth call to Cubersource API
func setupPayerAuth(w http.ResponseWriter, req *http.Request) {
	log.Println("setupPayerAuth")

	defer req.Body.Close()

	var setupPayerAuthRequestData threeds.SetupPayerAuthRequestData

	err := json.NewDecoder(req.Body).Decode(&setupPayerAuthRequestData)
	if err != nil {
		log.Printf("setupPayerAuth - Erro converting Post Body to JSON - %s", err)
		return
	}

	log.Println("setupPayerAuth - Executing Setup Payer Auth call.")

	setupPayerResp, msg, err := threeds.SetupPayerAuthRequest(&credentials.CyberSourceCredential, &setupPayerAuthRequestData)

	if err != nil {
		log.Println("setupPayerAuth - Error on SetupPayerAuthRequest.")
		log.Println(err)

		errorString := "setupPayerAuth - Error: " + err.Error() + " Msg string: " + msg
		http.Error(w, errorString, http.StatusBadRequest)
		return
	}

	log.Println("setupPayerAuth - Converting setup payer auth response to string.")

	setupPayerAuthResponseJSON, err := json.Marshal(setupPayerResp)

	if err != nil {
		errorString := "setupPayerAuth - Erro converting struct to JSON - " + err.Error()
		log.Printf("setupPayerAuth - Erro converting struct to JSON - %s", err)
		http.Error(w, errorString, http.StatusBadRequest)
		return
	}

	log.Println("setupPayerAuth response:")
	log.Println(string(setupPayerAuthResponseJSON))

	w.Write(setupPayerAuthResponseJSON)
}

// doEnrollment - Initiate enrollment process
func doEnrollment(w http.ResponseWriter, req *http.Request) {
	log.Println("doEnrollment")

	defer req.Body.Close()

	var enrollmentData threeds.EnrollmentRequestData

	err := json.NewDecoder(req.Body).Decode(&enrollmentData)
	if err != nil {
		log.Printf("doEnrollment - Erro converting Post Body to JSON - %s", err)
		return
	}

	// Add missing data
	var messageCategory = "01"
	var productCode = "01"
	var transactionMode = threeds.TransactionModeECOMMERCE
	var acsWindowSize = "02"

	var requestorID = "CARDCYBS_5b16ebc085282c2b20313e7b"
	var requestorName = "Braspag"

	var merchantName = "Brazil Test"
	var merchantURL = "https://merchantrul.com"
	var mcc = "5399"

	var returnUrl = "https://webhook.dreambridge.net/webhook"

	var ipAddress = req.RemoteAddr
	var httpAcceptBrowserValueStr = "*/*"

	enrollmentData.ConsumerAuthenticationInformation.MCC = &mcc
	enrollmentData.ConsumerAuthenticationInformation.RequestorID = &requestorID
	enrollmentData.ConsumerAuthenticationInformation.RequestorName = &requestorName
	enrollmentData.ConsumerAuthenticationInformation.MessageCategory = &messageCategory
	enrollmentData.ConsumerAuthenticationInformation.ProductCode = &productCode
	enrollmentData.ConsumerAuthenticationInformation.TransactionMode = &transactionMode
	enrollmentData.ConsumerAuthenticationInformation.AcsWindowSize = &acsWindowSize

	enrollmentData.ConsumerAuthenticationInformation.ReturnUrl = &returnUrl

	MerchantInformationData := new(threeds.MerchantInformation)
	MerchantDescriptorData := new(threeds.MerchantDescriptor)

	MerchantInformationData.MerchantName = &merchantName
	MerchantDescriptorData.Name = &merchantName
	MerchantDescriptorData.URL = &merchantURL

	MerchantInformationData.MerchantDescriptor = MerchantDescriptorData
	enrollmentData.MerchantInformation = MerchantInformationData

	enrollmentData.DeviceInformation.IPAddress = &ipAddress
	enrollmentData.DeviceInformation.HTTPAcceptBrowserValue = &httpAcceptBrowserValueStr

	enrollmentResponse, returnString, err := threeds.EnrollmentRequest(&credentials.CyberSourceCredential, &enrollmentData)

	if err != nil || enrollmentResponse == nil {
		log.Println("doEnrollment - Error during enrollment request: " + returnString)
		http.Error(w, returnString, http.StatusBadRequest)
		return
	}

	enrollmentResponseJSON, err := json.Marshal(enrollmentResponse)

	if err != nil {
		errorString := "doEnrollment - Erro converting struct to JSON - " + err.Error()
		log.Printf("doEnrollment - Erro converting struct to JSON - %s", err)
		http.Error(w, errorString, http.StatusBadRequest)
		return
	}

	log.Println("Enrollment response:")
	log.Println(string(enrollmentResponseJSON))

	w.Write(enrollmentResponseJSON)
}

func validate(w http.ResponseWriter, req *http.Request) {
	log.Println("validate")

	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("validate - Erro reading request body.")
		log.Fatalln(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var bodyMap map[string]string
	var bodySplit []string = strings.Split(string(body), "&")

	bodyMap = make(map[string]string)
	for _, pair := range bodySplit {
		z := strings.Split(pair, "=")
		bodyMap[z[0]] = z[1]
	}

	log.Println("TransactionId: " + bodyMap["TransactionId"])

	authenticationTransactionID := bodyMap["TransactionId"]

	validationRequestData := threeds.ValidationRequestData{
		ConsumerAuthenticationInformation: &commons.ConsumerAuthenticationInformation{
			AuthenticationTransactionID: &authenticationTransactionID,
		},
	}

	validationResponse, returnString, err := threeds.ValidationtRequest(&credentials.CyberSourceCredential, &validationRequestData)

	if err != nil || validationResponse == nil {
		log.Println("validate - Error during validation request: " + returnString)
		log.Println(err)
		http.Error(w, returnString, http.StatusBadRequest)
		return
	}

	validationResponseJSON, err := json.Marshal(validationResponse)

	if err != nil {
		errorString := "validate - Erro converting struct to JSON - " + err.Error()
		log.Printf("validate - Erro converting struct to JSON.")
		log.Println(err)
		http.Error(w, errorString, http.StatusInternalServerError)
		return
	}

	w.Write(validationResponseJSON)
}

// main_test.go
var (
	_ http.Handler = http.HandlerFunc((&controller{}).index)
	_ http.Handler = http.HandlerFunc((&controller{}).healthz)
	_ middleware   = (&controller{}).logging
	_ middleware   = (&controller{}).tracing
)
