package util

import (
	"crypto/md5"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/alwindoss/morse"
	"strings"
)

func EncryptMorse(texto string) string {

	h := morse.NewHacker()
	morseCode, err := h.Encode(strings.NewReader(texto))
	if err != nil {
		return ""
	}
	return string(morseCode)
}

func EncryptMd5(texto string) string {

	hash := md5.Sum([]byte(texto))
	return hex.EncodeToString(hash[:])
}

func EncryptSha512(texto string) string {

	hash := []byte(texto)
	dig := sha512.Sum512(hash)
	for i := 1; i < 5000; i++ {
		dig = sha512.Sum512(append(dig[:], hash[:]...))
	}
	newPasswd := base64.StdEncoding.EncodeToString(dig[:])
	return fmt.Sprintf("%s", newPasswd)
}

func EncryptMurcielago(texto string) string {
	replacer := strings.NewReplacer("M", "0", "U", "1", "R", "2", "C", "3", "I", "4", "E", "5", "L","6",
		"A","7","G","8","O","9","m", "0", "u", "1", "r", "2", "c", "3", "i", "4", "i", "5", "i","6",
		"a","7","g","8","o","9")
	return replacer.Replace(texto)

}
