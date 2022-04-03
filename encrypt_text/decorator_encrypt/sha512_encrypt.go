package decorator_encrypt

import (
	"configurations/sarmiento-erick-0422/encrypt_text/util"
	"fmt"
)

type sha512Encryptor struct {
	encrypt IEncryptDecorator
}

func NewSha512Encryptor(encrypt IEncryptDecorator) IEncryptDecorator {
	return &sha512Encryptor{encrypt: encrypt}
}

func (m *sha512Encryptor) EncryptDecorator(TextToEncrypt string, user string) {
	if m.encrypt != nil {
		m.encrypt.EncryptDecorator(TextToEncrypt, user)
	}

	fmt.Println(util.Concatena("Para el usuario ", user, " la encriptación del texto ", "en método SHA512 es: ", util.EncryptSha512(TextToEncrypt)))
}
