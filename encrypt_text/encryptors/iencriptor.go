package encryptors

type IEncryptText interface {
	EncryptText(text string, user string)
}
