package encryptors

import "configurations/sarmiento-erick-0422/encrypt_text/decorator_encrypt"

type adminEncryptor struct {
	decoratorEncrypt decorator_encrypt.IEncryptDecorator
}

func NewAdminEncryptor(decoratorEncrypt decorator_encrypt.IEncryptDecorator) IEncryptText {
	return &adminEncryptor{decoratorEncrypt: decoratorEncrypt}
}

func (a *adminEncryptor) EncryptText(text string, user string) {
	a.decoratorEncrypt.EncryptDecorator(text, user)
}
