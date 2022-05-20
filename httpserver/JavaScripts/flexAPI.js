// Asks the backend for one KeyID to generate one cardToken
function getCardTokenKey() {
    console.log("getCardTokenKey");

    var settings = {
        "async": false,
        "crossDomain": true,
        "url": "http://localhost:5000/getFlexAPIKey",
        "method": "GET",
        "headers": {
          "Accept": "*/*",
          "Cache-Control": "no-cache",
          "cache-control": "no-cache"
        }
      };
      
    $.ajax(settings).done(function (response, status) {

      console.log("Status: "+ status);

      console.log("keyID: "+ response);

      setKeyID(response);

      enableInputs();

    }).fail(failResponse);
}

// Generates one cardtoken using the card form and the KeyID generated into the backend
function generateCardToken() {
  console.log("generateCardToken");

  disableInputs();

  var cardTokenData = {
      keyId: document.getElementById('keyID').value, 
      cardInfo: {
          cardNumber: document.getElementById("cardNumber").value,
          cardType: document.getElementById("type").value,
      }
  };

  var payload = JSON.stringify(cardTokenData);
  console.log("Payload: " + payload);

  var settings = {
      "async": false,
      "crossDomain": true,
      "url": "https://apitest.cybersource.com/flex/v1/tokens",
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

      console.log("Response: ");
      console.log(response);

      setCardToken(response._embedded.icsReply.instrumentIdentifier.id);

      window.alert("CardToken gerado com sucesso:\n"+response._embedded.icsReply.instrumentIdentifier.id);

    }).fail(failResponse);

}

// Asks the backend for one KeyID with cryptography to generate one cardToken
function getCardTokenKeyCrypto() {
  console.log("getCardTokenKeyCrypto");

  var ret = false;

  var settings = {
      "async": false,
      "crossDomain": true,
      "url": "http://localhost:5000/getFlexAPIKeyCrypto",
      "method": "GET",
      "headers": {
        "Accept": "*/*",
        "Cache-Control": "no-cache"
      }
    };
    
  $.ajax(settings).done(function (response, status) {

    console.log("Status: "+ status);

    console.log("keyID: "+ response);

    setKeyID(response);

    enableInputs();

    ret = true;

  }).fail(failResponse);

  return ret;
}

function generateCardTokenCrypto() {
  console.log("generateCardTokenCrypto: ");

  disableInputs();

  var tokenKeyObj = JSON.parse(document.getElementById('keyID').value);

  var options = {
    kid: tokenKeyObj.keyId,
    keystore: tokenKeyObj.jwk,
    encryptionType: 'RsaOaep256', // ensure this matches the encryptionType you specified when creating your key
    cardInfo: {
      cardNumber: document.getElementById("cardNumber").value,
      cardType: document.getElementById("type").value,
    }
  };
 
  FLEX.createToken(options, response => {
    if (response.error) {

      failResponse(response, status);
      
    } else {
      console.log("Response object:");
      console.log(response);
  
      setCardToken(response._embedded.icsReply.instrumentIdentifier.id);
  
      window.alert("CardToken gerado com sucesso:\n"+response._embedded.icsReply.instrumentIdentifier.id);
    } 
  });
}

function generateCardTokenCryptoCallBack(callBack) {
  console.log("generateCardTokenCryptoCallBack: ");

  disableCardInputs();

  var tokenKeyObj = JSON.parse(document.getElementById('keyID').value);

  const options = {
    kid: tokenKeyObj.keyId,
    keystore: tokenKeyObj.jwk,
    encryptionType: 'RsaOaep256', // ensure this matches the encryptionType you specified when creating your key
    cardInfo: {
      cardNumber: document.getElementById("cardNumber").value,
      cardExpirationMonth: document.getElementById("expMonth").value,
      cardExpirationYear: document.getElementById("expYear").value,
      cardType: document.getElementById("type").value
    }
  };
 
  FLEX.createToken(options, response => {
    if (response.error) {

      failResponse(response, status);
      
    } else {
      console.log("Response object:");
      console.log(response);
  
      setCardToken(response.token);
  
      callBack();
    } 
  });
}

function setKeyID(keyID) {
  document.getElementById("keyID").setAttribute('value',keyID);
}

function setCardToken(cardToken) {
  document.getElementById("cardToken").setAttribute('value',cardToken);
}

  // Enable all card inputs and buttons fields
  function enableInputs() {
    document.getElementById("cardNumber").disabled = false;
    document.getElementById("type").disabled = false;
    document.getElementById("tokenizeBtn").disabled = false;
  }
  
  // Disable all card inputs and buttons fields
  function disableInputs() {
    document.getElementById("cardNumber").disabled = true;
    document.getElementById("type").disabled = true;
    document.getElementById("tokenizeBtn").disabled = true;
  }