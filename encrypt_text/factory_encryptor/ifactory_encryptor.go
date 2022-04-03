package factory_encryptor

import "configurations/sarmiento-erick-0422/encrypt_text/encryptors"

type IFactoryEncryptor interface {
	GetEncryptor(id string) encryptors.IEncryptText
}
