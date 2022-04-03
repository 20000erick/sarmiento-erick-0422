package encryptors

import "configurations/sarmiento-erick-0422/encrypt_text/decorator_encrypt"

type userEncryptor struct {
	decoratorEncrypt decorator_encrypt.IEncryptDecorator
}

func NewUserEncryptor(decoratorEncrypt decorator_encrypt.IEncryptDecorator) IEncryptText {
	return &userEncryptor{decoratorEncrypt: decoratorEncrypt}
}

func (a *userEncryptor) EncryptText(text string, user string) {
	a.decoratorEncrypt.EncryptDecorator(text, user)
}
