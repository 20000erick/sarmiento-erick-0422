package decorator_encrypt

import (
	"configurations/sarmiento-erick-0422/encrypt_text/util"
	"fmt"
)

type batEncryptor struct {
	encrypt IEncryptDecorator
}

func NewBatEncryptor(encrypt IEncryptDecorator) IEncryptDecorator {
	return &batEncryptor{encrypt: encrypt}
}

func (m *batEncryptor) EncryptDecorator(TextToEncrypt string, user string) {

	if m.encrypt != nil {
		m.encrypt.EncryptDecorator(TextToEncrypt, user)
	}
	fmt.Println(util.Concatena("Para el usuario ", user, " la encriptación del texto ", "en método Murcielago es: ", util.EncryptMurcielago(TextToEncrypt)))
}
