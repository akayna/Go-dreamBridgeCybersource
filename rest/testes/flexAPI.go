package main

import (
	"Go-dreamBridgeCybersource/rest/commons"
	"Go-dreamBridgeCybersource/rest/flexAPI"
	"Go-dreamBridgeUtils/jsonfile"
	"fmt"
	"log"
)

func main() {
	var credentials commons.Credentials

	err := jsonfile.ReadJSONFile2("/home/rafaelsonhador/Documents/Credenciais Cybersource/", "vileve.json", &credentials)

	if err != nil {
		log.Println("Erro ao ler credenciais.")
		log.Println("Erro: ", err)

		return
	}

	generatedKey, msg, err := flexAPI.GenerateKey(&credentials.CyberSourceCredential, nil)

	if err != nil {
		log.Println("main - Error generating key.")
		log.Println(err)
		return
	}

	fmt.Println(msg)
	fmt.Printf("Key: %+v\n", generatedKey)
	fmt.Printf("KeyID: %+v\n", *generatedKey.KeyID)

}
