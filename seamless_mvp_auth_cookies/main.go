package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/mergermarket/go-pkcs7"
)

var (
	ssoTokenString = "040de6b0464f148845b79786a4ea82be-5b2174c29296c0cce0b7e4c7089e5a2988a5b4137853968c9c85a4c922574ec0e8fa607a99a92d85ea74243503ee5ed2c007707ae42e4d1d007fd55f417c2798483143f362ff72617dc0669e97e1d80cc41aea14b16c2d9add03c768d5f998cf"
	ssoTsel        = "04d1c2a0a63bf7cd643e5a2b266cf333-3f31d60f8a19faf204c57533ad4822d9b9add3f699491965dd4976fd119355f1c0fded5f8e29fa1d66ce42ed8d6d41cde05923bee938b775b15d3ddfc29d61bfee5a3bae576dfd53f5c3de20fd0d6d303f6ae3bef3defd2765571fa49cb0be2986e0c95869ae8d7b1fef68e88945b0908cee0088a3d547880625c19246a7ca8618cb4443fb87c355efed116c8c014d20590f9900c25ab7efa334212115a3014bc9f970419bdf6b5c58ac60031b50f806ca3b7342967d5e1223c529b00eac0a270ab9798d02cdd3e246659d2af5d6ee050941ef16ed556841c26378c08840b668"
)

func main() {
	// resDecrypt, err := Decrypt(ssoTokenString)
	// if err != nil {
	// 	fmt.Println("ERROR PROSES DECRYPT :", err.Error())
	// }
	// fmt.Println(fmt.Sprintf("Encrypted : %v", ssoTokenString))
	// fmt.Println()
	// fmt.Println(fmt.Sprintf("Decrypted : %v", resDecrypt))

	resDecrypt, err := MyTselDecrypt(ssoTsel)
	if err != nil {
		fmt.Println("ERROR PROSES DECRYPT :", err.Error())
	}
	fmt.Println(fmt.Sprintf("Encrypted : %v", ssoTokenString))
	fmt.Println()
	fmt.Println(fmt.Sprintf("Decrypted : %v", resDecrypt))

	// resEncrypt,_:=Encrypt(ssoTokenString)
}

func Pad(buf []byte, size int) ([]byte, error) {
	bufLen := len(buf)
	padLen := size - bufLen%size
	padded := make([]byte, bufLen+padLen)
	copy(padded, buf)
	for i := 0; i < padLen; i++ {
		padded[bufLen+i] = byte(padLen)
	}
	return padded, nil
}

func Unpad(padded []byte, size int) ([]byte, error) {
	if size < 1 {
		return nil, errors.New("Block size looks wrong")
	}

	if len(padded)%size != 0 {
		return nil, errors.New("Data isn't aligned to size")
	}

	if len(padded) == 0 {
		return nil, errors.New("Data is empty")
	}

	paddingLength := int(padded[len(padded)-1])
	for _, el := range padded[len(padded)-paddingLength:] {
		if el != byte(paddingLength) {
			err := fmt.Sprintf("Padding had malformed entries. Have '%x', expected '%x'", paddingLength, el)
			return nil, errors.New(err)
		}
	}

	return padded[:len(padded)-paddingLength], nil
}

// Encrypt encrypts plain text string into cipher text string
func Encrypt(unencrypted string) (string, error) {
	CIPHERKEY := "OMtT8Wsf8Wy6WH2X2BdD2CFIx9WWCE9c"
	key := []byte(CIPHERKEY)
	plainText := []byte(unencrypted)
	plainText, err := Pad(plainText, aes.BlockSize)
	if err != nil {
		return "", fmt.Errorf(`plainText: "%s" has error`, plainText)
	}
	if len(plainText)%aes.BlockSize != 0 {
		err := fmt.Errorf(`plainText: "%s" has the wrong block size`, plainText)
		return "", err
	}

	block, err := aes.NewCipher(key)
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

// Decrypt decrypts cipher text string into plain text string
func Decrypt(encrypted string) (string, error) {
	//production
	// CIPHERKEY := "PMtA8Wsf8Zy6WA2X2BdD2CFIx9WXUE8k"

	//preproduction
	CIPHERKEY := "OMtT8Wsf8Wy6WH2X2BdD2CFIx9WWCE9c"
	key := []byte(CIPHERKEY)
	cipherText, _ := hex.DecodeString(encrypted)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println("1", err.Error())
		return "", err
	}

	if len(cipherText) < aes.BlockSize {
		fmt.Println("2", err.Error())
		return "", errors.New("cipherText too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		fmt.Println("3", "cipherText is not a multiple of the block size")
		return "", errors.New("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, err = pkcs7.Unpad(cipherText, aes.BlockSize)
	fmt.Println(err)
	return string(cipherText), nil
}

var (
	cipherNewCBCDecrypter = cipher.NewCBCDecrypter
	aesNewCipher          = aes.NewCipher
	hexDecodeString       = hex.DecodeString
)

func MyTselDecrypt(encrypted string) (string, error) {

	//BdjFKfZJ5lKG9kqaDYrNwNhTilrqOECB
	tselKey := os.Getenv("GDJVesGYZJNEVUU7UhgGEhwBa8fv5nmk")
	key := []byte(tselKey)
	concat := strings.Replace(encrypted, "-", "", -1)
	cipherText, _ := hexDecodeString(concat)
	block, err := aesNewCipher(key)
	if err != nil {
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		return "", errors.New("len ciphertext < aes.BlockSize")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		return "", errors.New("len ciphertext invalid")
	}
	mode := cipherNewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)
	cipherText, _ = pkcs7.Unpad(cipherText, aes.BlockSize)
	chiperString := fmt.Sprintf("%s", cipherText)
	encrypted = chiperString

	return encrypted, nil
}

func UnpadTsel(data []byte, blockSize uint) ([]byte, error) {
	if blockSize < 1 {
		return nil, errors.New("Block size looks wrong")
	}
	if uint(len(data))%blockSize != 0 {
		return nil, errors.New("Data isn't aligned to blockSize")
	}
	if len(data) == 0 {
		return nil, errors.New("Data is empty")
	}
	paddingLength := int(data[len(data)-1])
	for _, el := range data[len(data)-paddingLength:] {
		if el != byte(paddingLength) {
			err := fmt.Sprintf("Padding had malformed entries. Have '%x', expected '%x'", paddingLength, el)
			return nil, errors.New(err)
		}
	}
	return data[:len(data)-paddingLength], nil
}
