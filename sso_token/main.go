package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"strings"
	"time"
)

var (
	cipherNewCBCDecrypter = cipher.NewCBCDecrypter
	aesNewCipher          = aes.NewCipher
	hexDecodeString       = hex.DecodeString
)

func main() {
	custparam := "014c5c93ac762d638512d56f25ef5e64-d1d1a6ba77be129364ac7152243f5215e1006c400dee0e880888f81e2e10c3ac37045d35cefaca9912825cef94dd83564805e37bf55e145c0fd79b1d538926206239f4ffcdc21ad6ac47ee68a3cf309e4f5a8e0e2c48665ea7df09db85dad097"
	decrypt1, err := MyTselDecrypt(custparam)
	if err != nil {
		fmt.Println("error x")
	}
	fmt.Println("first decrypt", decrypt1)
	splt := strings.Split(decrypt1, "|")
	str := splt[1]
	cut := str[:len(str)-3]
	tmp1, _ := time.Parse("060102150405", cut)
	tmp := tmp1.Local()
	fmt.Println("timestamp", str)
	ssoexpired := 60
	tmp2 := tmp.Add(time.Minute * time.Duration(ssoexpired))

	duration := time.Since(tmp2)
	fmt.Println("TIME SINCE :", int(math.Abs(duration.Seconds())), tmp)
	fmt.Println("TIME COMPARE", tmp2.Unix(), time.Now().Unix())
	fmt.Println("TMP2", tmp2, int(tmp2.Unix())-(int(time.Now().Unix())+25200))
	fmt.Println("xxxx ", time.Duration(ssoexpired)*time.Minute)
	fmt.Println("xxxx ", time.Duration(int(tmp2.Unix()-time.Now().Unix())-1)*time.Second)
	x := time.Now().Add(time.Minute * time.Duration(300))
	fmt.Println(x)
	fmt.Println("ZZZ ", int(tmp2.Unix()-x.Unix()))

	if tmp2.Unix() < time.Now().Unix() {
		fmt.Println("sso token expired")
	}
	msisdn, err := MyTselDecrypt(splt[0])
	if err != nil {
		fmt.Println("error")
	}
	msisdn2 := strings.ReplaceAll(msisdn, " ", "")
	fmt.Println("index 0", splt[0], "msisdn", msisdn2+"msisdn")

	if err != nil {
		fmt.Println("error 2")
	}
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
	cipherText, err = Unpad(cipherText, aes.BlockSize)
	if err != nil {
		return "", err
	}
	chiperString := fmt.Sprintf("%s", cipherText)
	encrypted = chiperString
	// if x == 1 {
	// 	fmt.Println("1")
	// 	return nil
	// }
	// x -= 1
	// fmt.Println(x)
	// //fmt.Println(chiperString)
	// TselDecrypt(ar, tselKey, x)
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
