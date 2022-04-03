package configurations

import (
	"configurations/sarmiento-erick-0422/encrypt_text/decorator_encrypt"
	"configurations/sarmiento-erick-0422/encrypt_text/encryptors"
	"configurations/sarmiento-erick-0422/encrypt_text/factory_encryptor"
)

type configuration struct {
	FactoryEncryptor factory_encryptor.IFactoryEncryptor
}

// NewConfigurationEncrypt NewConfigurationValidator / Agregar nuevas implementaciones aqui
func NewConfigurationEncrypt() *configuration {
	sha512Encrypt := decorator_encrypt.NewSha512Encryptor(nil)
	shaAndMurcielagoEncrypt := decorator_encrypt.NewBatEncryptor(sha512Encrypt)

	md5Encrypt := decorator_encrypt.NewMd5Encryptor(nil)
	morseAndMd5Encrypt := decorator_encrypt.NewMorseEncryptor(md5Encrypt)

	adminSecurity1 := encryptors.NewAdminEncryptor(shaAndMurcielagoEncrypt)
	userSecurity1 := encryptors.NewUserEncryptor(morseAndMd5Encrypt)

	roles := map[string]encryptors.IEncryptText{
		"adminSecurity1": adminSecurity1,
		"userSecurity1":  userSecurity1,
	}

	factoryRoles := factory_encryptor.NewEncryptorFactory(roles)
	return &configuration{FactoryEncryptor: factoryRoles}
}
