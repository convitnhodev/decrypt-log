package decryptRepository

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	b64 "encoding/base64"
	"fmt"
	"os"
	"regexp"
)

func (text *ManagerString) GetSource() string {
	return text.Source
}

func (text *ManagerString) SetDes(input string) {
	text.Des = input
}

func (text *ManagerString) ParseLog() [][]int {

	r := regexp.MustCompile("<" + regexp.QuoteMeta("?") + ">.*?</" + regexp.QuoteMeta("?") + ">")
	matches := r.FindAllStringIndex(text.Source, -1)

	return matches
}

func (text *ManagerString) DecryptLogSnippet(commonIVSecret string, keyTextSecret string) string {

	var commonIV = []byte(commonIVSecret)

	// aes encryption string

	keyText := sha256.Sum256([]byte(keyTextSecret))

	// Create the aes encryption algorithm
	c, err := aes.NewCipher(keyText[:])
	if err != nil {
		fmt.Printf("Error: NewCipher(%d bytes) = %s", len(keyText), err)
		os.Exit(-1)
	}
	ciphertext, _ := b64.StdEncoding.DecodeString(text.Des)

	// Decrypt strings
	cfbdec := cipher.NewCFBDecrypter(c, commonIV)
	plaintextCopy := make([]byte, len(ciphertext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	return string(plaintextCopy)
}
