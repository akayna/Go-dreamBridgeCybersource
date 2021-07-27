// Asks the backend for one KeyID to generate one cardToken
function getCardTokenKey() {
    console.log("getCardTokenKey");

    var settings = {
        "async": true,
        "crossDomain": true,
        "url": "http://localhost:5000/getFlexAPIKey",
        "method": "GET",
        "headers": {
          "Accept": "*/*",
          "Cache-Control": "no-cache",
          //"Host": "localhost:5000",
          //"accept-encoding": "gzip, deflate",
          //"Connection": "keep-alive",
          "cache-control": "no-cache"
        }
      };
      
    $.ajax(settings).done(function (response, status) {

      console.log("Status: "+ status);

      console.log("keyID: "+ response);

      setKeyID(response);

      enableCardInputs();

      return false;
    }).fail(failResponse);

    return false;
}

// Generates one cardtoken using the card form and the KeyID generated into the backend
function generateCardToken() {
  console.log("generateCardToken");

  var cardTokenData = {
      keyId: document.getElementById('keyID').value, 
      cardInfo: {
          cardNumber: document.getElementById("cardNumber").value,
          cardExpirationMonth: document.getElementById("expMonth").value,
          cardExpirationYear: document.getElementById("expYear").value,
          cardType: document.getElementById("type").value
      }
  };

  var payload = JSON.stringify(cardTokenData);
  console.log("Payload: " + payload);

  var settings = {
      "async": true,
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

      setCardToken(response.token);

      disableCardInputs();

      window.alert("CardToken gerado com sucesso:\n"+response.token);

      return false;
    }).fail(failResponse);

  return false;
}

// Asks the backend for one KeyID with cryptography to generate one cardToken
function getCardTokenKeyCrypto() {
  console.log("getCardTokenKey");

  var settings = {
      "async": true,
      "crossDomain": true,
      "url": "http://localhost:5000/getFlexAPIKeyCrypto",
      "method": "GET",
      "headers": {
        "Accept": "*/*",
        "Cache-Control": "no-cache",
        //"Host": "localhost:5000",
        //"accept-encoding": "gzip, deflate",
        //"Connection": "keep-alive",
        "cache-control": "no-cache"
      }
    };
    
  $.ajax(settings).done(function (response, status) {

    console.log("Status: "+ status);

    console.log("keyID: "+ response);

    setKeyID(response);

    enableCardInputs();

    return false;
  }).fail(failResponse);

  return false;
}

function generateCardTokenCrypto() {
  console.log("generateCardTokenCrypto: ");

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
  
      disableCardInputs();
  
      window.alert("CardToken gerado com sucesso:\n"+response.token);
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
function enableCardInputs() {
  document.getElementById("cardNumber").disabled = false;
  document.getElementById("expMonth").disabled = false;
  document.getElementById("expYear").disabled = false;
  document.getElementById("type").disabled = false;

  if (document.getElementById("tokenizeBtn") != null) 
    document.getElementById("tokenizeBtn").disabled = false;
}

// Disable all card inputs and buttons fields
function disableCardInputs() {
  document.getElementById("cardNumber").disabled = true;
  document.getElementById("expMonth").disabled = true;
  document.getElementById("expYear").disabled = true;
  document.getElementById("type").disabled = true;
  
  if (document.getElementById("tokenizeBtn") != null) 
    document.getElementById("tokenizeBtn").disabled = true;
}