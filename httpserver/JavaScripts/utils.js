/*window.onload = function(){
    document.getElementById('authorizeBtn').onclick = sendAuthorizationRequest;
    document.getElementById('authenticateBtn').onclick = initiateAuthentication;
    document.getElementById('tokenizeBtn').onclick = getCardTokenKey;

    mountDeviceFingerprint();

    configureSongbird();

    getJWT();
};*/

// Treats the one fail response.
function failResponse(response, status) {
  console.log("Status: "+ status);

  console.log("Response object:");

  console.log(response);

}

// Enable all inputs and buttons fields
function enablePaymentInputs() {
    document.getElementById("ccnum").disabled = false;
    document.getElementById("cname").disabled = false;
    document.getElementById("cvv").disabled = false;
    document.getElementById("expmonth").disabled = false;
    document.getElementById("expyear").disabled = false;
    document.getElementById("type").disabled = false;
    document.getElementById("overridePaymentMethod").disabled = false;
    document.getElementById("cpf").disabled = false;

    document.getElementById("authenticateBtn").disabled = false;
    document.getElementById('tokenizeBtn').disabled = false;
    document.getElementById('authorizeBtn').disabled = false;

    document.getElementById("firstName").disabled = false;
    document.getElementById("lastName").disabled = false;
    document.getElementById("address1").disabled = false;
    document.getElementById("address2").disabled = false;
    document.getElementById("locality").disabled = false;
    document.getElementById("country").disabled = false;
    document.getElementById("administrativeArea").disabled = false;
    document.getElementById("postalCode").disabled = false;
    document.getElementById("email").disabled = false;
    document.getElementById("phoneNumber").disabled = false;
    document.getElementById("mobilePhone").disabled = false;

    document.getElementById("totalAmount").disabled = false;
    document.getElementById("currency").disabled = false;
    document.getElementById("orderNumber").disabled = false;
}

// Disable all inputs and buttons fields
function disablePaymentInputs() {
    document.getElementById("ccnum").disabled = true;
    document.getElementById("cname").disabled = true;
    document.getElementById("cvv").disabled = true;
    document.getElementById("expmonth").disabled = true;
    document.getElementById("expyear").disabled = true;
    document.getElementById("type").disabled = true;
    document.getElementById("overridePaymentMethod").disabled = true;
    document.getElementById("cpf").disabled = true;

    document.getElementById("authenticateBtn").disabled = true;
    document.getElementById('tokenizeBtn').disabled = true;
    document.getElementById('authorizeBtn').disabled = true;

    document.getElementById("firstName").disabled = true;
    document.getElementById("lastName").disabled = true;
    document.getElementById("address1").disabled = true;
    document.getElementById("address2").disabled = true;
    document.getElementById("locality").disabled = true;
    document.getElementById("country").disabled = true;
    document.getElementById("administrativeArea").disabled = true;
    document.getElementById("postalCode").disabled = true;
    document.getElementById("email").disabled = true;
    document.getElementById("phoneNumber").disabled = true;
    document.getElementById("mobilePhone").disabled = true;

    document.getElementById("totalAmount").disabled = true;
    document.getElementById("currency").disabled = true;
    document.getElementById("orderNumber").disabled = true;
}

function saveAuthenticationInfo(eci_ucafind, cavv_ucafdata, xid) {
    document.getElementById("ECI_UCAFIND").setAttribute('value',eci_ucafind);
    document.getElementById("CAVV_UCAFDATA").setAttribute('value',cavv_ucafdata);
    document.getElementById("XID").setAttribute('value',xid);
}

function setCardToken(cardToken) {
    document.getElementById("cardToken").setAttribute('value',cardToken);
}

function getAuthenticationData() {

    var authenticationData;
  
    if (document.getElementById("cardToken").value != "") {
  
      authenticationData = {
        clientReferenceInformation: {
          code: document.getElementById("orderNumber").value,
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
            phoneNumber: document.getElementById("phoneNumber").value,
            email: document.getElementById("email").value,
            postalCode: document.getElementById("postalCode").value,
          },
          amountDetails: {
            currency: document.getElementById("currency").value,
            totalAmount: document.getElementById("totalAmount").value,
          },
        },
        paymentInformation: {
          customer: {
            customerId: document.getElementById("cardToken").value,
          },
        },
        buyerInformation: {
          mobilePhone: document.getElementById("mobilePhone").value,
        },
        consumerAuthenticationInformation: {
          overridePaymentMethod: document.getElementById("overridePaymentMethod").value,
          referenceId: document.getElementById("orderNumber").value,//document.getElementById("referenceID").value,
          deviceChannel: "Browser",
        },
        deviceInformation: {
          ipAddress: "198.241.159.102",
          httpAcceptBrowserValue: "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
          httpAcceptContent: "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
          httpBrowserLanguage: "en-US",
          httpBrowserJavaEnabled: "Y",
          httpBrowserJavaScriptEnabled: "Y",
          httpBrowserColorDepth: "24",
          httpBrowserScreenHeight: "864",
          httpBrowserScreenWidth: "1536",
          httpBrowserTimeDifference: "300",
          userAgentBrowserValue: "Mozilla/5.0 (X11; CrOS x86_64 8172.45.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.64 Safari/537.36",
        },
      };
  
    } else {
      authenticationData = {
        clientReferenceInformation: {
          code: document.getElementById("orderNumber").value,
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
            phoneNumber: document.getElementById("phoneNumber").value,
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
            expirationMonth: document.getElementById("expmonth").value,
            expirationYear: document.getElementById("expyear").value,
            number: document.getElementById("ccnum").value,
            type: document.getElementById("type").value,
          },
        },
        buyerInformation: {
          mobilePhone: document.getElementById("mobilePhone").value,
        },
        consumerAuthenticationInformation: {
          overridePaymentMethod: document.getElementById("overridePaymentMethod").value,
          referenceId: document.getElementById("orderNumber").value,//document.getElementById("referenceID").value,
          deviceChannel: "Browser",
        },
        deviceInformation: {
          ipAddress: "198.241.159.102",
          httpAcceptBrowserValue: "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
          httpAcceptContent: "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8",
          httpBrowserLanguage: "en-US",
          httpBrowserJavaEnabled: "Y",
          httpBrowserJavaScriptEnabled: "Y",
          httpBrowserColorDepth: "24",
          httpBrowserScreenHeight: "864",
          httpBrowserScreenWidth: "1536",
          httpBrowserTimeDifference: "300",
          userAgentBrowserValue: "Mozilla/5.0 (X11; CrOS x86_64 8172.45.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/51.0.2704.64 Safari/537.36",
        },
      };
    }
  
    return authenticationData;
  }

  function getPaymentData() {

    var paymentData;
  
      paymentData = {
        clientReferenceInformation: {
          code: document.getElementById("orderNumber").value,
        },
        processingInformation: {
          commerceIndicator: "internet",
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
            phoneNumber: document.getElementById("phoneNumber").value,
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
            expirationMonth: document.getElementById("expmonth").value,
            expirationYear: document.getElementById("expyear").value,
            number: document.getElementById("ccnum").value,
            type: document.getElementById("type").value,
            useAs: document.getElementById("overridePaymentMethod").value,
            securityCode: document.getElementById("cvv").value,
          },
        },
        buyerInformation: {
          mobilePhone: document.getElementById("mobilePhone").value,
        },
        /*consumerAuthenticationInformation: {
          overridePaymentMethod: document.getElementById("overridePaymentMethod").value,
          referenceId: document.getElementById("orderNumber").value,//document.getElementById("referenceID").value,
          deviceChannel: "Browser",
        },*/
      };
    return paymentData;
  }