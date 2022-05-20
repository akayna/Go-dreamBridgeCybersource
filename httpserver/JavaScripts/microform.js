function setupMicroform() {
    console.log("setupMicroform");

    var tokenizeButton = document.querySelector('#tokenizeBtn');
    var flexResponse = document.querySelector('#flexresponse');
    var expMonth = document.querySelector('#expMonth');
    var expYear = document.querySelector('#expYear');
    var errorsOutput = document.querySelector('#errors-output');

    // the capture context that was requested server-side for this transaction
    var captureContext = document.getElementById('microformContext').value;

    // custom styles that will be applied to each field we create using Microform
    var myStyles = {
        'input': {    
            'font-size': '14px',    
            'font-family': 'helvetica, tahoma, calibri, sans-serif',    
            'color': '#555'  
    },  
        ':focus': { 'color': 'blue' },  
        ':disabled': { 'cursor': 'not-allowed' },  
        'valid': { 'color': '#3c763d' },  
        'invalid': { 'color': '#a94442' }
    };

    // setup
    var flex = new Flex(captureContext);
    var microform = flex.microform({ styles: myStyles });
    var number = microform.createField('number', { placeholder: 'Digite o número do cartão' });
    var securityCode = microform.createField('securityCode', { placeholder: '••••' });

    number.load('#number-container');
    securityCode.load('#securityCode-container');

    tokenizeButton.addEventListener('click', function() {
        console.log("tokenizeButton");

        var options = {    
          expirationMonth: expMonth.value,  
          expirationYear: expYear.value
        };

        microform.createToken(options, function (err, token) {
          if (err) {
            // handle error
            console.error(err);
            errorsOutput.textContent = err.message;
          } else {
            // At this point you may pass the token back to your server as you wish.
            // In this example we append a hidden input to the form and submit it. 
            console.log("JWT Cardtoken:");     
            console.log(JSON.stringify(token));
            flexResponse.value = JSON.stringify(token);

            validateToken(token);

            //console.log("Status: "+ status);

            //console.log("Response: ");
            //console.log(response);
      
            disableInputs();
      
            window.alert("CardToken gerado com sucesso.");
          }
        });
      });
}

function validateToken(token) {
  console.log("validateToken");

  var settings = {
    "async": false,
    "crossDomain": true,
    "url": "http://localhost:5000/validateMicroformToken",
    "method": "POST",
    "headers": {
        "Content-Type": "application/jwt;charset=UTF-8",
        "Accept": "*/*",
        "Cache-Control": "no-cache",
        "cache-control": "no-cache"
    },
    "processData": false,
    "data": token,
  };

  $.ajax(settings).done(function (response, status) {

    console.log("Status: "+ status);

    console.log("Context: "+ response);

    return false;
  }).fail(failResponse);

  return false;
}


function getMicroformContext() {
    console.log("getMicroformContext");

    var settings = {
        "async": false,
        "crossDomain": true,
        "url": "http://localhost:5000/getMicroformContext",
        "method": "GET",
        "headers": {
          "Accept": "*/*",
          "Cache-Control": "no-cache",
          "cache-control": "no-cache"
        }
      };
      
    $.ajax(settings).done(function (response, status) {

      console.log("Status: "+ status);

      console.log("Context: "+ response);

      document.getElementById("microformContext").setAttribute('value',response);

      enableInputs();

      return false;
    }).fail(failResponse);

    return false;
}

// Enable all card inputs and buttons fields
function enableInputs() {
    document.getElementById("expMonth").disabled = false;
    document.getElementById("expYear").disabled = false;
    document.getElementById("tokenizeBtn").disabled = false;
  }
  
  // Disable all card inputs and buttons fields
  function disableInputs() {
    document.getElementById("expMonth").disabled = true;
    document.getElementById("expYear").disabled = true;
    document.getElementById("tokenizeBtn").disabled = true;
  }