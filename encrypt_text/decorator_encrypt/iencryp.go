package decorator_encrypt

type IEncryptDecorator interface {
	EncryptDecorator(TextToEncrypt string, user string)
}
