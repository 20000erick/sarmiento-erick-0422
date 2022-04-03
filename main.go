package main

import (
	"configurations/sarmiento-erick-0422/encrypt_text/configurations"
	"configurations/sarmiento-erick-0422/kit/constants"
	"fmt"
)

func main() {

	user := "adminSecurity1"

	GetEncryptedKeys(user, "Hola holA")
}

func GetEncryptedKeys(User string, TextToEncrypt string) {
	loadConfiguration := configurations.NewConfigurationEncrypt().FactoryEncryptor.GetEncryptor(User)
	if loadConfiguration != nil {
		loadConfiguration.EncryptText(TextToEncrypt, User)
	} else {
		fmt.Println(constants.UserUndefined)
	}

}
