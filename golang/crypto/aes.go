package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"os"
)

func getKeyBytes(key string) []byte {
	keyBytes := []byte(key)
	switch l := len(keyBytes); {
	case l < 16:
		keyBytes = append(keyBytes, make([]byte, 16-l)...)
	case l > 16:
		keyBytes = keyBytes[:16]
	}
	return keyBytes
}

func encrypt(key string, origData []byte) ([]byte, error) {
	keyBytes := getKeyBytes(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keyBytes[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func decrpt(key string, crypted []byte) ([]byte, error) {
	keyBytes := getKeyBytes(key)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyBytes[:blockSize])
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func Encrypt(key string, val string) (string, error) {
	origData := []byte(val)
	crypted, err := encrypt(key, origData)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(crypted), nil
}

func Decrypt(key string, val string) (string, error) {
	crypted, err := base64.URLEncoding.DecodeString(val)
	if err != nil {
		return "", err
	}
	origData, err := decrpt(key, crypted)
	if err != nil {
		return "", err
	}
	return string(origData), nil
}

func main() {

	argc := len(os.Args)
	if argc != 4 {
		os.Stdout.WriteString("usage: AES (-e|-d) <key> <content>")
		return
	}

	switch os.Args[1] {
	case "-e":
		ret, err := Encrypt(os.Args[2], os.Args[3])
		if err != nil {
			os.Stderr.WriteString(err.Error())
			os.Exit(1)
		}
		println(ret)
	case "-d":
		ret, err := Decrypt(os.Args[2], os.Args[3])
		if err != nil {
			os.Stderr.WriteString(err.Error())
			os.Exit(1)
		}
		println(ret)
	default:
		os.Stdout.WriteString("usage: AES (-e|-d) <key> <content>")
	}
}
