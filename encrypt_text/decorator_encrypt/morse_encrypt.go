package decorator_encrypt

import (
	"configurations/sarmiento-erick-0422/encrypt_text/util"
	"fmt"
)

type morseEncryptor struct {
	encrypt IEncryptDecorator
}

func NewMorseEncryptor(encrypt IEncryptDecorator) IEncryptDecorator {
	return &morseEncryptor{encrypt: encrypt}
}

func (m *morseEncryptor) EncryptDecorator(TextToEncrypt string, user string) {

	if m.encrypt != nil {
		m.encrypt.EncryptDecorator(TextToEncrypt, user)
	}

	fmt.Println(util.Concatena("Para el usuario ", user, " la encriptación del texto ", "en método Morse es: ", util.EncryptMorse(TextToEncrypt)))
}
