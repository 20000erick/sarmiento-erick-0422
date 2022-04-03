package factory_encryptor

import (
	"configurations/sarmiento-erick-0422/encrypt_text/encryptors"
)

type encryptorFactory struct {
	encryptors map[string]encryptors.IEncryptText
}

func NewEncryptorFactory(encryptors map[string]encryptors.IEncryptText) IFactoryEncryptor {
	return &encryptorFactory{encryptors: encryptors}
}

func (v *encryptorFactory) GetEncryptor(id string) encryptors.IEncryptText {
	_, exist := v.encryptors[id]
	if exist {
		return v.encryptors[id]
	}
	return nil
}
