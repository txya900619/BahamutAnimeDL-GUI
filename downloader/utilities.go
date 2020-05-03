package downloader

import (
	"crypto/aes"
	"crypto/cipher"
	"math/rand"
	"time"
)

func randomHash() string {
	rand.Seed(time.Now().UnixNano())
	dictionary := "1234567890qwertyuiopasdfghjklzxcvbnm"
	hash := ""
	for i := 0; i < 12; i++ {
		hash = hash + string(dictionary[rand.Intn(len(dictionary))])
	}
	return hash
}

func decryptAES128(encrypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	mode := cipher.NewCBCDecrypter(block, key[:aes.BlockSize])
	decryptedData := make([]byte, len(encrypted))
	mode.CryptBlocks(decryptedData, encrypted)
	decryptedData = pkc5Unpadding(decryptedData)

	return decryptedData, nil
}

func pkc5Unpadding(decryptedData []byte) []byte {
	dataLength := len(decryptedData)
	unPadding := int(decryptedData[dataLength-1])
	return decryptedData[:(dataLength - unPadding)]
}
