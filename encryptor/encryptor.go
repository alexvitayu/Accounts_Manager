package encryptor

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"os"
)

type Encryptor struct {
	Key string
}

func NewEncryptor() *Encryptor {
	key := os.Getenv("KEY")
	if key == "" {
		panic("не передан параметр Key в переменные окружения")
	}
	return &Encryptor{
		Key: key,
	}
}

func (enc *Encryptor) Encrypt(plainStr []byte) []byte {
	block, err := aes.NewCipher([]byte(enc.Key))
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

func (enc *Encryptor) Decrypt(encryptedStr []byte) []byte {
	return []byte{}
}
