package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encryptor struct {
	key string
}

func NewEncryptor() *Encryptor {
	key := os.Getenv("KEY")
	return &Encryptor{
		key: key,
	}
}

func (enc *Encryptor) Encrypt(plainStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.key))
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	return aesGCM.Seal(nonce, nonce, plainStr, nil)
}

func (enc *Encryptor) Decrypt(ecryptedStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.key))
	if err != nil {
		panic(err.Error())
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGCM.NonceSize()
	//nonce := ecryptedStr[:nonceSize]
	//cipherText := ecryptedStr[nonceSize:]
	nonce, cipherText := ecryptedStr[:nonceSize], ecryptedStr[nonceSize:]
	plainText, err := aesGCM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return plainText
}
