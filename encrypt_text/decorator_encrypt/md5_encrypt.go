package decorator_encrypt

import (
	"configurations/sarmiento-erick-0422/encrypt_text/util"
	"fmt"
)

type md5Encryptor struct {
	encrypt IEncryptDecorator
}

func NewMd5Encryptor(encrypt IEncryptDecorator) IEncryptDecorator {
	return &md5Encryptor{encrypt: encrypt}
}

func (m *md5Encryptor) EncryptDecorator(TextToEncrypt string, user string) {
	if m.encrypt != nil {
		m.encrypt.EncryptDecorator(TextToEncrypt, user)
	}

	fmt.Println(util.Concatena("Para el usuario ", user, " la encriptación del texto ", "en método MD5 es: ", util.EncryptMd5(TextToEncrypt)))
}
