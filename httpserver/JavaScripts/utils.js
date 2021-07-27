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

  // Enable all card inputs and buttons fields
function enableCardInputs() {
  document.getElementById("cardNumber").disabled = false;
  document.getElementById("expMonth").disabled = false;
  document.getElementById("expYear").disabled = false;
  document.getElementById("type").disabled = false;
  document.getElementById("overridePaymentMethod").disabled = false;
}

// Disable all card inputs and buttons fields
function disableCardInputs() {
  document.getElementById("cardNumber").disabled = true;
  document.getElementById("expMonth").disabled = true;
  document.getElementById("expYear").disabled = true;
  document.getElementById("type").disabled = true;
  document.getElementById("overridePaymentMethod").disabled = true;
}