function onVisaCheckoutReady() {
    V.init( {
    //apikey: "O468JVE2SHZ3LTIBSI1U21AOQxeAF0UrtY4f0Ofciprw5SXVg",
    apikey: "P9XTECIYD4ARG6M24EII21DX5uJWczjcbgoYYYoCrfDbPT-QQ",
    //encryptionKey: "RG1G42TE45BFT5O0107D13nES-18QRx1UqTwYGQVtWaKsCaLc",
    paymentRequest: {
    currencyCode: "BRL",
    subtotal: "10.00"
    }
    });
    V.on("payment.success", function(payment) {
    document.write("payment.success: \n" + JSON.stringify(payment));
    });
    V.on("payment.cancel", function(payment) {
    document.write("payment.cancel: \n" + JSON.stringify(payment));
    });
    V.on("payment.error", function(payment, error) {
    document.write("payment.error: \n" +
    JSON.stringify(payment) + "\n" +
    JSON.stringify(error));
    });
    }    