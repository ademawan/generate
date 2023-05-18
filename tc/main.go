package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TransformingInformation struct {
	Data string
}

var (
	cipherNewCBCDecrypter = cipher.NewCBCDecrypter
	helperMyTselDecrypt   = MyTselDecrypt
	helperMyTselDecrypt2  = MyTselDecrypt
	aesNewCipher          = aes.NewCipher
	hexDecodeString       = hex.DecodeString
)

func main() {
	var custparam = "935f9c16004b65257da3a0577099642f-dcae5a277dc09bc5f3f6f4e3f66ba641c72d5a4982886d47a31c380f975f55f945ca559302f53070a1e4185fb0543dc1fabe1a19db8cc2e43bb103deaf8d9650cd39d4f96df5e8edb5ad44421c27e1c65e3ff46d3a5144c115d98fba744bfa6d"
	decrypt1, err := helperMyTselDecrypt(custparam)
	if err != nil {
		fmt.Println("error 0")

		panic(err)
	}

	splt := strings.Split(decrypt1, "|")
	str := splt[1]
	cut := str[:len(str)-3]
	tmp1, _ := time.Parse("060102150405", cut)
	tmp := tmp1.Local()
	fmt.Println("timestamp", str)
	ssoexpired, _ := strconv.Atoi("9999")
	tmp2 := tmp.Add(time.Minute * time.Duration(ssoexpired))

	if tmp2.Unix() < time.Now().Unix() {
		fmt.Println("error 1")
		panic(err)
	}
	msisdn, err := helperMyTselDecrypt2(splt[0])
	if err != nil {
		fmt.Println("error 2")

		panic(err)
	}
	fmt.Println("h")
	msisdn2 := strings.ReplaceAll(msisdn, " ", "")
	fmt.Println("MSISDN:", msisdn2)

}

func MyTselDecrypt(encrypted string) (string, error) {
	tselKey := "BdjFKfZJ5lKG9kqaDYrNwNhTilrqOECB"
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
	// cipherText, _ = Unpad(cipherText, aes.BlockSize)
	cipherText, err = Unpad(cipherText, aes.BlockSize)
	if err != nil {
		return "", err
	}

	chiperString := fmt.Sprintf("%s", cipherText)
	encrypted = chiperString

	return encrypted, nil
}

func Unpad(data []byte, blockSize uint) ([]byte, error) {
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
