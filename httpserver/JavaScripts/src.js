var initData = {
    "srciTransactionId": "55490ddf-eb53-4962-8806-2da8094cad1c",
    "srciDpaId": "DPAID",
    "srcInitiatorId": "IPW4W1DX5D4N5E8LMAC6115la44Nu9zYob_bgrg5EUMksbgT4",
    "dpaData": {
        "srcDpaId": "",
        "dpaPresentationName": "Disney Online",
        "dpaUri": "http://www.disneyonline.com",
        "dpaThreeDsPreference": "UNKNOWN"
    },
        "dpaTransactionOptions": {
            "dpaLocale": "US",
            "dpaAcceptedBillingCountries": [
                "US",
                "CA"
            ],
            "dpaAcceptedShippingCountries": [
                "US",
                "CA"
            ],
            "dpaBillingPreference": "FULL",
            "dpaShippingPreference": "FULL",
            "consumerNameRequested": true,
            "consumerEmailAddressRequested": true,
            "consumerPhoneNumberRequested": true,
            "paymentOptions": {
            "dpaDynamicDataTtlMinutes": 2,
            "dynamicDataType": "TAVV",
            "dpaPanRequested": false
        },
            "reviewAction": "continue",
            "transactionType": "PURCHASE",
            "orderType": "REAUTHORIZATION",
            "payloadTypeIndicator": "FULL",
            "transactionAmount": {
            "transactionAmount": "99.95",
            "transactionCurrencyCode": "USD"
        },
            "merchantOrderId": "28b1b61b-bbec-4637-b78f-33babc3b5187",
            "merchantCategoryCode": "3000",
            "merchantCountryCode": "US",
            "threeDsInputData": {
            "requestorId": "requestorId",
            "acquirerId": "acquirerId",
            "acquirerMid": "acquirerMid"
        }
    }
};

async function initiSRC(initParams, cb) {
    const promiseData = await
        vSrc.init(initParams).then(function(response) {
        return response;
    })
    .catch(function(error) {
        return error;
    });
    
    cb(promiseData);
}
    
function callInitSRC() {
    initiSRC(initData, function(result) {
        console.log(result);
        if (!Object.keys(result).length) {
            console.log("init() successful");
            //Next -> call isRecognized()
            callIsRecognized();
        } else {
            console.log("init() failure");
            // Some error in init, validate input or retry
        }
    });
}