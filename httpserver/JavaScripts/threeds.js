function setupPayerAuthTkn(){
    console.log("setupPayerAuthTkn");

    if (getCardTokenKeyCrypto()) {
        enableAuthFields();

    } else {
        console.log("falha ao gerar token key.");
    }

}

function setupPayerAuthToken() {
    console.log("setupPayerAuthToken");

    generateCardTokenCryptoCallBack(callBackGenerateCardTokenSuccess);
}

function callBackGenerateCardTokenSuccess() {
    console.log("callBackGenerateCardTokenSuccess");
    var setupPayerAuthData = {
        clientReferenceInformation: {
            code: document.getElementById("referenceCode").value
        },
        paymentInformation: {
            card: {
                expirationMonth: document.getElementById("expMonth").value,
                expirationYear: document.getElementById("expYear").value,
                type: document.getElementById("type").value
            },
            instrumentIdentifier: {
                id: document.getElementById("cardToken").value
            }
        }
    };

    setupPayerAuth(setupPayerAuthData);    
}

function setupPayerAuthPAN() {
    console.log("setupPayerAuthPAN");

    var setupPayerAuthData = {
        clientReferenceInformation: {
            code: document.getElementById("referenceCode").value
        },
        paymentInformation: {
            card: {
                number: document.getElementById("cardNumber").value,
                expirationMonth: document.getElementById("expMonth").value,
                expirationYear: document.getElementById("expYear").value,
                type: document.getElementById("type").value
            }
        }
    };

    setupPayerAuth(setupPayerAuthData);
    
}

function setupPayerAuth(setupPayerAuthData) {
    console.log("setupPayerAuth");

    disableAuthFields();

    var payload = JSON.stringify(setupPayerAuthData);
    console.log("SetupPayerAuth Payload: " + payload);

    var settings = {
        "async": false,
        "crossDomain": true,
        "url": "http://localhost:5000/setupPayerAuth",
        "method": "POST",
        "headers": {
            "Content-Type": "application/json",
            "Accept": "*/*"
        },
        "processData": false,
        "data": payload
    };
        
    $.ajax(settings).done(function (response, status) {

        console.log("Status: "+ status);

        if (response !== null) {
            console.log("Response: ");
            console.log(response);
    
            var responseObj = JSON.parse(response);
    
            document.getElementById('ddc-form').setAttribute('action',responseObj.consumerAuthenticationInformation.deviceDataCollectionUrl);
            document.getElementById('JWT_ddc').setAttribute('value',responseObj.consumerAuthenticationInformation.accessToken);
            document.getElementById('referenceID').setAttribute('value',responseObj.consumerAuthenticationInformation.referenceId);
    
            createDeviceDataCollectionListener();
            challengeFormSend();
        } else {
            console.log("Response is null.");
        }

        return true;
    }).fail(failResponse);

    return false;
}

function enableAuthFields() {
    document.getElementById("cardNumber").disabled = false;
    document.getElementById("expMonth").disabled = false;
    document.getElementById("expYear").disabled = false;
    document.getElementById("type").disabled = false;
    document.getElementById("overridePaymentMethod").disabled = false;
    document.getElementById("cpfNumber").disabled = false;
    document.getElementById("authBtn").disabled = false;
    document.getElementById("referenceCode").disabled = false;
    document.getElementById("firstName").disabled = false;
    document.getElementById("lastName").disabled = false;
    document.getElementById("address1").disabled = false;
    document.getElementById("address2").disabled = false;
    document.getElementById("locality").disabled = false;
    document.getElementById("administrativeArea").disabled = false;
    document.getElementById("country").disabled = false;
    document.getElementById("postalCode").disabled = false;
    document.getElementById("email").disabled = false;
    document.getElementById("mobilePhone").disabled = false;
    document.getElementById("totalAmount").disabled = false;
    document.getElementById("currency").disabled = false;
}

function disableAuthFields() {
    document.getElementById("cardNumber").disabled = true;
    document.getElementById("expMonth").disabled = true;
    document.getElementById("expYear").disabled = true;
    document.getElementById("type").disabled = true;
    document.getElementById("overridePaymentMethod").disabled = true;
    document.getElementById("cpfNumber").disabled = true;
    document.getElementById("authBtn").disabled = true;
    document.getElementById("referenceCode").disabled = true;
    document.getElementById("firstName").disabled = true;
    document.getElementById("lastName").disabled = true;
    document.getElementById("address1").disabled = true;
    document.getElementById("address2").disabled = true;
    document.getElementById("locality").disabled = true;
    document.getElementById("administrativeArea").disabled = true;
    document.getElementById("country").disabled = true;
    document.getElementById("postalCode").disabled = true;
    document.getElementById("email").disabled = true;
    document.getElementById("mobilePhone").disabled = true;
    document.getElementById("totalAmount").disabled = true;
    document.getElementById("currency").disabled = true;
}

function setRandomReferenceCode() {
    document.getElementById("referenceCode").value = "test_Rafael_" + Date.now();
}

function challengeFormSend() {
    console.log('challengeFormSend:');

    var ddcForm = document.querySelector('#ddc-form');
    if (ddcForm) {
        ddcForm.submit();
    }
}

function createDeviceDataCollectionListener() {
    console.log('createDeviceDataCollectionListener:');

    window.addEventListener("message", function (event) {

            console.log('window event listener: ');
            console.log(event);

            if (event.origin === "https://centinelapistag.cardinalcommerce.com") {
            //if (event.origin === "https://centinelapi.cardinalcommerce.com") {
                var data = JSON.parse(event.data);
                console.log('Merchant received a message:', data);

                if (data.Status) {
                    console.log('Songbird ran DF successfully');

                    // Initiate enrollment process
                    doEnrollment();
                } else {
                    console.log('Message from different type.');
                }

            } else {
                console.log('Message from different origin.');
            }

        }, false);
}

function createChallengeValidatedEventLiestener() {
    console.log('createChallengeValidatedEvent:');

    var challengeIframe = document.getElementById("step-up-iframe");

    challengeIframe.addEventListener("click", function (event) {
            console.log('Event:');
            console.log(event);
        }, false);
}

function doEnrollment() {
    console.log("doEnrollment.");
  
    var authenticationData = JSON.stringify(getAuthenticationData());
  
    console.log("Data sent to erollment:");
    console.log(authenticationData);

    var settings = {
        "async": false,
        "crossDomain": true,
        "url": "http://localhost:5000/doEnrollment",
        "method": "POST",
        "headers": {
            "Content-Type": "application/json",
            "Accept": "*/*",
            "Cache-Control": "no-cache",
            "cache-control": "no-cache"
        },
        "processData": false,
        "data": authenticationData,
    };

    $.ajax(settings).done(function (response, status) {

        console.log("Status: "+ status);

        console.log("Enrollment Response: "+ response);

        treatEnrollmentResponse(response);
    
    }).fail(failResponse);
}

function testeStepupcallback() {
    console.log("testeStepupcallback");
  
}

function treatEnrollmentResponse(enrollmentResponse) {
    var objEnrollment = JSON.parse(enrollmentResponse);
  
    console.log("3DS protocol version: "+ objEnrollment.consumerAuthenticationInformation.specificationVersion);
  
    switch(objEnrollment.consumerAuthenticationInformation.veresEnrolled) {
      case "Y":
          switch(objEnrollment.consumerAuthenticationInformation.paresStatus) {
            case "Y":
              console.log("Successful silent authentication.");
              window.alert("Successful silent authentication.");
              break;
            case "N":
              console.log("Payer cannot be authenticated - Unsuccessful.");
              window.alert("Payer cannot be authenticated - Unsuccessful.");
              break;
            case "A":
                console.log("Stand-in silent authentication.");
                window.alert("Stand-in silent authentication.");
              break;
            case "U":
              console.log("Payer cannot be authenticated - Unavailable.");
              window.alert("Payer cannot be authenticated - Unavailable.");
              break;
            case "R":
              console.log("Payer cannot be authenticated - Rejected.");
              window.alert("Payer cannot be authenticated - Rejected.");
              break;
            default: //case "C":
              console.log("Step-up.");

              var stepUpForm = document.querySelector('#step-up-form');

              if (stepUpForm) {
                  
                    // Convert the pareq from base64 to string
                  pareqString = atob(objEnrollment.consumerAuthenticationInformation.pareq);

                  var objPareq = JSON.parse(pareqString);
          
                  console.log("Pareq: " + pareqString);
          
                  adjustStepUpWindowsSize(objPareq.challengeWindowSize);
          
                  stepUpForm.setAttribute('Action',objEnrollment.consumerAuthenticationInformation.stepUpUrl);
          
                  var inputs = stepUpForm.elements;
                  var JWTInput = inputs.JWT_step;
          
                  JWTInput.setAttribute('value',objEnrollment.consumerAuthenticationInformation.accessToken);

                  showStepUpButton();
              }
              break;
          }
        break;
      case "U":
        console.log("Payer cannot be authenticated - not Available.");
        window.alert("Payer cannot be authenticated - not Available.");
        break;
      case "B":
        console.log("Bypassed Authentication.");
        window.alert("Bypassed Authentication.");
        break;
      case "R":
        console.log("Payer cannot be authenticated - Rejected.");
        window.alert("Payer cannot be authenticated - Rejected.");
        break;
      default:
        console.log("Authentication failure. Unknown response received.");
        break;
    }
  }

function treatValidationResponse(validationResponse) {
    var objValidation = JSON.parse(validationResponse);

    saveAutenticationData(objValidation);

    console.log("3DS protocol version: "+ objValidation.consumerAuthenticationInformation.specificationVersion);

    switch(objValidation.consumerAuthenticationInformation.paresStatus) {
        case "Y":
            console.log("Successful authentication.");
            window.alert("Successful authentication.");
            break;
        case "N":
            console.log("Payer could be authenticated - Unsuccessful.");
            window.alert("Payer could be authenticated - Unsuccessful.");
            break;
        case "U":
            console.log("Payer could be authenticated - Unavailable.");
            window.alert("Payer could be authenticated - Unavailable.");
            break;
        case "B":
            console.log("Payer could be authenticated - MerchantBypass.");
            window.alert("Payer could be authenticated - MerchantBypass.");
            break;
        default:
            console.log("Unexpected response. paresStatus: " + objValidation.consumerAuthenticationInformation.paresStatus);
            window.alert("Unexpected response. paresStatus: " + objValidation.consumerAuthenticationInformation.paresStatus);
            break;
    }
}

function executeChallenge() {
    console.log("executeChallenge.");

    var stepUpForm = document.querySelector('#step-up-form');

    if (stepUpForm) {
        blockStepUpButton();
        console.log("opening challenge...");
        //createChallengeValidatedEventLiestener();
        stepUpForm.submit();
    }
}

function adjustStepUpWindowsSize(size) {
    console.log("adjustStepUpWindowsSize.");

    var width;
    var height;

    switch(size) {
        case "01":
            width = 250;
            height = 400;
            break;
        case "02":
            width = 390;
            height = 400;
            break;
        case "03":
            width = 500;
            height = 600;
            break;
        case "04":
            width = 600;
            height = 400;
            break;
        case "05":
            width = "100%";
            height = "100%";
            break;
        default:
            console.log("Error - invalid challengeWindowSize");

    }

    document.getElementById('step-up-iframe').setAttribute('width',width);
    document.getElementById('step-up-iframe').setAttribute('height',height);
}

function getAuthenticationData() {
    console.log("getAuthenticationData.");

    var authenticationData;

    authenticationData = {
        clientReferenceInformation: {
            code: document.getElementById("referenceCode").value,
        },
        orderInformation: {
            billTo: {
                address1: document.getElementById("address1").value,
                address2: document.getElementById("address2").value,
                administrativeArea: document.getElementById("administrativeArea").value,
                country: document.getElementById("country").value,
                locality: document.getElementById("locality").value,
                firstName: document.getElementById("firstName").value,
                lastName: document.getElementById("lastName").value,
                phoneNumber: document.getElementById("mobilePhone").value,
                email: document.getElementById("email").value,
                postalCode: document.getElementById("postalCode").value,
            },
            amountDetails: {
                currency: document.getElementById("currency").value,
                totalAmount: document.getElementById("totalAmount").value,
            },
        },
        paymentInformation: {
            card: {
                expirationMonth: document.getElementById("expMonth").value,
                expirationYear: document.getElementById("expYear").value,
                number: document.getElementById("cardNumber").value,
                type: document.getElementById("type").value,
            },
        },
        buyerInformation: {
            merchantCustomerId: document.getElementById("cpfNumber").value,
            mobilePhone: document.getElementById("mobilePhone").value,
        },
        consumerAuthenticationInformation: {
            overridePaymentMethod: document.getElementById("overridePaymentMethod").value,
            referenceId: document.getElementById("referenceID").value,
            deviceChannel: "Browser",
        },
        deviceInformation: mountDeviceFingerprint()
    };

    return authenticationData;
}

function mountDeviceFingerprint() {
    console.log("mountDeviceFingerprint.");

    var deviceInformation;

    httpBrowserJavaEnabledBool = "";
    if (navigator.javaEnabled() == true)
    {
        httpBrowserJavaEnabledBool= "Y";
    }
    else
    {
        httpBrowserJavaEnabledBool= "N";
    }

    var date = new Date();

    deviceInformation = {
        httpBrowserColorDepth: screen.colorDepth,
        httpBrowserJavaEnabled: httpBrowserJavaEnabledBool,
        httpBrowserJavaScriptEnabled: "Y",
        httpBrowserLanguage: navigator.language || navigator.userLanguage,
        httpBrowserScreenHeight: window.innerHeight,
        httpBrowserScreenWidth: window.innerWidth,
        httpBrowserTimeDifference: date.getTimezoneOffset(),
    };

    return deviceInformation;
  }

function showStepUpButton() {
    console.log("showStepUpButton.");

    document.getElementById("stepUpBtn").disabled = false;
    document.getElementById("stepUpBtn").style.display = "block";
}

function hideStepUpButton() {
    document.getElementById("stepUpBtn").disabled = true;
    document.getElementById("stepUpBtn").style.display = "none";
}

function blockStepUpButton() {
    document.getElementById("stepUpBtn").disabled = true;
}