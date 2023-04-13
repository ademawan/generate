package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"wec-auth/helper/exceptions"
)

var (
	cipherKey = "Xs5AL7oKTo83R2022"
)

func main() {

	test := []byte(cipherKey)
	by := make([]byte, 32)
	copy(by, test)
	encryptedEmail, err := EncryptEmail("rogertest51@email.ghostinspector.com|1", by)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(encryptedEmail)
	email, err := DecryptEmail(encryptedEmail)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(email)
}

// DecryptEmail function
func DecryptEmail(encrypted string) (decrypted string, err error) {

	test := []byte(cipherKey)
	cipherKeyByte := make([]byte, 32)
	copy(cipherKeyByte, test)

	cipherText, err := hex.DecodeString(encrypted)
	if err != nil {
		err = exceptions.ErrValidDecrypt
		return
	}
	block, err := aes.NewCipher(cipherKeyByte)
	if err != nil {
		err = exceptions.ErrValidDecrypt
		return
	}

	if len(cipherText) < aes.BlockSize {
		err = exceptions.ErrValidDecrypt
		return
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		err = errors.New("invalid data")
		return
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = Unpad(cipherText, aes.BlockSize)
	decrypted = fmt.Sprintf("%s", cipherText)

	return
}

// EncryptUserInfo function
func EncryptUserInfo(unEncrypted string, cipherKey []byte) (string, error) {
	// key := []byte(cipherKey)
	plainText := []byte(unEncrypted)
	plainText, err := Pad(plainText, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf(`plainText: "%s" has error`, plainText)
	}
	if len(plainText)%aes.BlockSize != 0 {
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return "", err
	}

	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return fmt.Sprintf("%x", cipherText), nil
}

// Pad function
func Pad(data []byte, blockSize uint) ([]byte, error) {
	neededBytes := blockSize - (uint(len(data)) % blockSize)
	return append(data, bytes.Repeat([]byte{byte(neededBytes)}, int(neededBytes))...), nil

}

// Unpad function
func Unpad(data []byte, blockSize uint) ([]byte, error) {
	if blockSize < 1 {
		return nil, fmt.Errorf("Block size looks wrong")
	}

	if uint(len(data))%blockSize != 0 {
		return nil, fmt.Errorf("Data isn't aligned to blockSize")
	}

	if len(data) == 0 {
		return nil, fmt.Errorf("Data is empty")
	}

	paddingLength := int(data[len(data)-1])
	for _, el := range data[len(data)-paddingLength:] {
		if el != byte(paddingLength) {
			return nil, fmt.Errorf("Padding had malformed entries. Have '%x', expected '%x'", paddingLength, el)
		}
	}

	return data[:len(data)-paddingLength], nil
}

func EncryptEmail(unEncrypted string, cipherKey []byte) (string, error) {
	// key := []byte(cipherKey)
	plainText := []byte(unEncrypted)
	plainText, err := Pad(plainText, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf(`plainText: "%s" has error`, plainText)
	}
	if len(plainText)%aes.BlockSize != 0 {
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return "", err
	}

	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(plainText))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(cipherText[aes.BlockSize:], plainText)

	return fmt.Sprintf("%x", cipherText), nil
}
