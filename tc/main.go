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

type TCStruct struct {
	Name string
	Age  int
}
type ChildStruct struct {
}

var (
	cipherNewCBCDecrypter = cipher.NewCBCDecrypter
	helperMyTselDecrypt   = MyTselDecrypt
	helperMyTselDecrypt2  = MyTselDecrypt
	aesNewCipher          = aes.NewCipher
	hexDecodeString       = hex.DecodeString
)

func main() {

	tcStruct := []*TCStruct{}

	var tcStruct2 []*TCStruct
	tcStruct3 := []*TCStruct{}
	tcStruct3 = append(tcStruct3, &TCStruct{})
	fmt.Println(tcStruct, tcStruct2, tcStruct3)
	//==============================================
	var custparam = "014c5c93ac762d638512d56f25ef5e64-d1d1a6ba77be129364ac7152243f5215e1006c400dee0e880888f81e2e10c3ac37045d35cefaca9912825cef94dd83564805e37bf55e145c0fd79b1d538926206239f4ffcdc21ad6ac47ee68a3cf309e4f5a8e0e2c48665ea7df09db85dad097"
	decrypt1, err := helperMyTselDecrypt(custparam)
	if err != nil {
		fmt.Println("error 0")

		panic(err)
	}

	fmt.Println(fmt.Sprintf("%v", true))

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
	tselKey := "GDJVesGYZJNEVUU7UhgGEhwBa8fv5nmk"
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
