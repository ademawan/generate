package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

const (
	blockSize = 16
)

func main() {

	////3b4599332d162fd95a599941166249a9-11515fd0a1b8824ba58581c00637baeffc048ce3014e8e001a173259ee40ca145e72eadf35ea0d6a8da7a9a623098b97490351a5883b9419f2be11facf33c6d60ab5818cdada71f1ca2ff3ea2dab0774z
	c, _ := hex.DecodeString("11515fd0a1b8824ba58581c00637baeffc048ce3014e8e001a173259ee40ca145e72eadf35ea0d6a8da7a9a623098b97490351a5883b9419f2be11facf33c6d60ab5818cdada71f1ca2ff3ea2dab0774z")
	// fmt.Println(c)
	// d:=fmt.Printf("% x", c)
	ff := fmt.Sprintf("% x", c)
	fmt.Println(ff)

	// b := []byte("3b4599332d162fd95a599941166249a9")
	// fmt.Println(c)
	Decrypt()
}

func Decrypt() {
	key := []byte("GDJVesGYZJNEVUU7UhgGEhwBa8fv5nmk")
	cipherText, _ := hex.DecodeString("3b4599332d162fd95a599941166249a911515fd0a1b8824ba58581c00637baeffc048ce3014e8e001a173259ee40ca145e72eadf35ea0d6a8da7a9a623098b97490351a5883b9419f2be11facf33c6d60ab5818cdada71f1ca2ff3ea2dab0774z")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	if len(cipherText) < aes.BlockSize {
		panic("cipherText too short")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		panic("cipherText is not a multiple of the block size")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = Unpad(cipherText, aes.BlockSize)
	x := fmt.Sprintf("%s", cipherText)
	fmt.Println(x)

	en, _ := Encrypt("628121314356")
	z := fmt.Sprintf("%s", en)
	fmt.Println(z)

}
func Encrypt(unencrypted string) (string, error) {
	key := []byte("GDJVesGYZJNEVUU7UhgGEhwBa8fv5nmk")
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

func Pad(data []byte, blockSize uint) ([]byte, error) {
	neededBytes := blockSize - (uint(len(data)) % blockSize)
	return append(data, bytes.Repeat([]byte{byte(neededBytes)}, int(neededBytes))...), nil

}
