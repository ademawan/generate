package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/forgoer/openssl"
	"github.com/golang-jwt/jwt"
)

const (
	blockSize = 16
)

var (
	errPosition = "usecase.TransformProcess"
	dataTest    = ""
)

type TransformingInformation struct {
	Data          string
	TransformType string
}

func main() {
	transformingInformation := &TransformingInformation{}
	transformingInformation.Data = "3d6154b13b275978eaaaf3f2ab97e2b1-9315e364b24339effa802a16b658440b41dc38f34806adbd20b9db7c65b47a3b2d4214892e6df8ac1caf026dbf138c36b69e3794a44a21d28c141480155c106f2c7c8679af4174a34ffa4961af5fe517d097eee9df463372bc5a3f3d383a8002"
	transformingInformation.TransformType = "A"
	// TranformProcess(transformingInformation)
	// fmt.Println(transformingInformation.Data)
	// test := []byte("Xs5AL7oKTo83R2022")
	// by := make([]byte, 32)
	// copy(by, test)
	// encryptData := "e840ed03f761fe7f00d2e0d90f5fca2666e9d697b60bbca668d1eeb49a8ba904d9fd6f239a690bfc4de52b09689a744eca57bf5d9d3a26fcc1170c05389e95497843348a2de8d45dd7f7cb3cee0431ed651368e691ff98318cb4d514690f17f4a260b87c25b74f7bdc9252d6b0270f42109c23a030d1c29fdfedfede66695f87c008f8809c528b4a674963b6048f8ee1bf4d0d53a075c39c52fba16d238f2136"
	// ress, errt := Decrypt(encryptData, by)
	// if errt != nil {
	// 	fmt.Println(errt.Error())
	// }
	// fmt.Println(ress)
	// claims, valid := ValidateJwt("eyJ0eXAiOiJKV1QiLCJraWQiOiJ3VTNpZklJYUxPVUFSZVJCL0ZHNmVNMVAxUU09IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI2OTMwYjEwOC00NDU4LTRmNDMtYTA4Zi0xNWRhMzZhY2ZlZjEiLCJjdHMiOiJPQVVUSDJfR1JBTlRfU0VUIiwiYXV0aF9sZXZlbCI6MCwiYXVkaXRUcmFja2luZ0lkIjoiYTJlZWQ4NjUtZWI5Zi00MGEzLWI2MTktZmFhMWQ2MWUyZjAzLTc3MTk3MCIsImlzcyI6Imh0dHBzOi8vY2lhbWFtcHJlcGRhcHAuY2lhbS50ZWxrb21zZWwuY29tOjEwMDAzL29wZW5hbS9vYXV0aDIvdHNlbC93ZWMvd2ViIiwidG9rZW5OYW1lIjoiYWNjZXNzX3Rva2VuIiwidG9rZW5fdHlwZSI6IkJlYXJlciIsImF1dGhHcmFudElkIjoibG1qLUNFbjdlQ1FvQVFuTFQtNElnYU9pQ2p3LjQweTN5U0VMZl8xZGlqdERZcUI3d29NbHRaVSIsIm5vbmNlIjoidHJ1ZSIsImF1ZCI6ImIzOTM2MTg0MzZlNTExZWM4ZDNkMDI0MmFjMTMwMDAzIiwibmJmIjoxNjc0NzkxNzM5LCJncmFudF90eXBlIjoiYXV0aG9yaXphdGlvbl9jb2RlIiwic2NvcGUiOlsib3BlbmlkIiwicHJvZmlsZSJdLCJhdXRoX3RpbWUiOjE2NzQ3OTE3MzksInJlYWxtIjoiL3RzZWwvd2VjL3dlYiIsImV4cCI6MTY3NDc5MjYzOSwiaWF0IjoxNjc0NzkxNzM5LCJleHBpcmVzX2luIjo5MDAsImp0aSI6Imxtai1DRW43ZUNRb0FRbkxULTRJZ2FPaUNqdy53Y2hVUy1nRUc2dkZfTXByQU50dUJHNzRaYWcifQ.n-zCNJWGYqcBZ7eqiozDuTlzHsrsQmomYUkcUAcOdsZEjm-wdMwsaA6G98gRNRdd6386sTcO1ljJMgREJ57NBin_d3NC1eMmYvqwr4FlWsqNWzwengTbPx1LWNzx5LiAN3P3nAqRJ3x-um6OwOr5SaOON0hQulOe2UvdUaF4e6dr2L2ZD8nvEN8MAnvC-AuO3mK4idxYZCoQgNo-fpujGBBy7_-LA7L7d2Nn1d91kJGhaBu7p01dtF9G98zm2YV9sYQvNDbKnLe9_1_BpB0MRHU_zJlN9IYOYqAMNb2E0BwwXCJp8UFl5lqc3bD1Qs5dbwzQ9TbRrF4zKKrzAK2UQg")
	// fmt.Println("CLAIM:", claims)
	// fmt.Println("valid:", valid)

	// data, err := DecryptDigiSubs("3d6154b13b275978eaaaf3f2ab97e2b1-9315e364b24339effa802a16b658440b41dc38f34806adbd20b9db7c65b47a3b2d4214892e6df8ac1caf026dbf138c36b69e3794a44a21d28c141480155c106f2c7c8679af4174a34ffa4961af5fe517d097eee9df463372bc5a3f3d383a8002")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println("tes", data)
	err3 := TselDecrypt(transformingInformation, "BdjFKfZJ5lKG9kqaDYrNwNhTilrqOECB", 2)
	if err3 != nil {
		fmt.Println(err3.Error())
	}
}

func TranformProcess(ar *TransformingInformation) (string, error) {
	//example if get config from DB
	//eyJ0eXAiOiJKV1QiLCJraWQiOiJ3VTNpZklJYUxPVUFSZVJCL0ZHNmVNMVAxUU09IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiI3YzVjOTZiZS0zMGFjLTQ5NzctYjkyYy02MjY0NjMyYjIwNjkiLCJjdHMiOiJPQVVUSDJfR1JBTlRfU0VUIiwiYXV0aF9sZXZlbCI6MCwiYXVkaXRUcmFja2luZ0lkIjoiNWFkOTIyNDEtZWJiNy00OWVmLTg2MzYtOGE2YTViYTliMjRhLTc1MjE1NiIsImlzcyI6Imh0dHBzOi8vY2lhbWFtcHJlcGRhcHAuY2lhbS50ZWxrb21zZWwuY29tOjEwMDAzL29wZW5hbS9vYXV0aDIvdHNlbC93ZWMvd2ViIiwidG9rZW5OYW1lIjoiYWNjZXNzX3Rva2VuIiwidG9rZW5fdHlwZSI6IkJlYXJlciIsImF1dGhHcmFudElkIjoibk1qWm53SFlIYzlCYVJ2REJEQzUxejNuZVpvLlBRY3hXRzVvOGNtaERfUEcxUEFyNWZmUTlWWSIsImF1ZCI6ImIzOTM2MTg0MzZlNTExZWM4ZDNkMDI0MmFjMTMwMDAzIiwibmJmIjoxNjc0Nzg4NDA2LCJncmFudF90eXBlIjoicmVmcmVzaF90b2tlbiIsInNjb3BlIjpbIm9wZW5pZCIsInByb2ZpbGUiXSwiYXV0aF90aW1lIjoxNjc0Nzg3NzY2LCJyZWFsbSI6Ii90c2VsL3dlYy93ZWIiLCJleHAiOjE2NzQ3ODkzMDYsImlhdCI6MTY3NDc4ODQwNiwiZXhwaXJlc19pbiI6OTAwLCJqdGkiOiJuTWpabndIWUhjOUJhUnZEQkRDNTF6M25lWm8uS2tnMzdnbjYxaUViVGc4b2EwUWQ2WGJNT3lNIn0.b2KXl7xmtaErT4xixLu_rQ9HmOpqOhFqhDTY3KF4MvBsf4F1SagEuttAF8Ycb-kP-sUHhNHHgcIMSiHFnRS0RNFI_Lx8q5oXTftVEQ7R9PNkQAVq42MxoAGrPtSgC5YaaL4ctOEFRzj0eFY-hjGfsEVIB8b8VhPFwW_XacZdl6bUgA5jNlTcHEhIupsB1JyER12DZpxjQEujcMVhlVMA_LOYKNFR-IwETn7u4J1IASg2n1p-T02aVG_TBLBtiIt4aDh2tkw6T03h8LKADLFYyQnZVLUuSivQSr5iq86tN86ghRkw1_PWdeDYGIPOfToL5sOxldSqWndLgn-_vgF8Yg|e840ed03f761fe7f00d2e0d90f5fca2666e9d697b60bbca668d1eeb49a8ba904d9fd6f239a690bfc4de52b09689a744eca57bf5d9d3a26fcc1170c05389e95497843348a2de8d45dd7f7cb3cee0431ed651368e691ff98318cb4d514690f17f4a260b87c25b74f7bdc9252d6b0270f42109c23a030d1c29fdfedfede66695f87c008f8809c528b4a674963b6048f8ee1bf4d0d53a075c39c52fba16d238f2136
	if ar.TransformType == "A" {
		tselKey := "GDJVesGYZJNEVUU7UhgGEhwBa8fv5nmk"

		err := TselDecrypt(ar, tselKey, 2)
		if err != nil {
			errPosition += "|TselDecrypt"
			return errPosition, err
		}

		return "", nil
	}

	return "", nil
}
func TselDecrypt(ar *TransformingInformation, tselKey string, x int) error {

	key := []byte(tselKey)
	concat := strings.Replace(ar.Data, "-", "", -1)
	// fmt.Println(concat)
	// split := strings.Split(concat, "|")
	// fmt.Println(split)
	// if x == 1 {
	// 	// concat = split[1]
	// 	dataTest = split[0]
	// 	fmt.Println(dataTest)
	// 	// fmt.Println(concat[:32])
	// 	// datas, _ := hex.DecodeString(dataTest)
	// 	// datax := fmt.Sprintf("%s", datas)
	// 	// fmt.Println("hallo", string(datax))
	// }
	// if x == 1 {
	// 	concat = dataTest
	// }
	cipherText, _ := hex.DecodeString(concat)
	// if x == 1 {
	// 	fmt.Println(concat[:32])
	// 	datax := fmt.Sprintf("%s", cipherText)
	// 	fmt.Println("hallo", datax)
	// 	fmt.Println()

	// }
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}

	if len(cipherText) < aes.BlockSize {
		return errors.New("len ciphertext < aes.BlockSize")
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	if len(cipherText)%aes.BlockSize != 0 {
		return errors.New("len ciphertext invalid")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	cipherText, _ = Unpad(cipherText, aes.BlockSize)
	chiperString := fmt.Sprintf("%s", cipherText)
	ar.Data = chiperString
	fmt.Println(ar.Data)

	if x == 1 {
		return nil
	}
	x -= 1

	TselDecrypt(ar, tselKey, x)

	return nil
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
func Pad(data []byte, blockSize uint) ([]byte, error) {
	neededBytes := blockSize - (uint(len(data)) % blockSize)
	return append(data, bytes.Repeat([]byte{byte(neededBytes)}, int(neededBytes))...), nil

}
func Decrypt(encrypted string, cipherKey []byte) (decrypted string, err error) {

	cipherText, _ := hex.DecodeString(encrypted)

	block, err := aes.NewCipher(cipherKey)
	if err != nil {
		return
	}

	if len(cipherText) < aes.BlockSize {
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

var js = `{
	"kty": "RSA",
	"kid": "wU3ifIIaLOUAReRB/FG6eM1P1QM=",
	"use": "sig",
	"x5t": "5eOfy1Nn2MMIKVRRkq0OgFAw348",
	"x5c": [
		"MIIDdzCCAl+gAwIBAgIES3eb+zANBgkqhkiG9w0BAQsFADBsMRAwDgYDVQQGEwdVbmtub3duMRAwDgYDVQQIEwdVbmtub3duMRAwDgYDVQQHEwdVbmtub3duMRAwDgYDVQQKEwdVbmtub3duMRAwDgYDVQQLEwdVbmtub3duMRAwDgYDVQQDEwdVbmtub3duMB4XDTE2MDUyNDEzNDEzN1oXDTI2MDUyMjEzNDEzN1owbDEQMA4GA1UEBhMHVW5rbm93bjEQMA4GA1UECBMHVW5rbm93bjEQMA4GA1UEBxMHVW5rbm93bjEQMA4GA1UEChMHVW5rbm93bjEQMA4GA1UECxMHVW5rbm93bjEQMA4GA1UEAxMHVW5rbm93bjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEBANdIhkOZeSHagT9ZecG+QQwWaUsi7OMv1JvpBr/7HtAZEZMDGWrxg/zao6vMd/nyjSOOZ1OxOwjgIfII5+iwl37oOexEH4tIDoCoToVXC5iqiBFz5qnmoLzJ3bF1iMupPFjz8Ac0pDeTwyygVyhv19QcFbzhPdu+p68epSatwoDW5ohIoaLzbf+oOaQsYkmqyJNrmht091XuoVCazNFt+UJqqzTPay95Wj4F7Qrs+LCSTd6xp0Kv9uWG1GsFvS9TE1W6isVosjeVm16FlIPLaNQ4aEJ18w8piDIRWuOTUy4cbXR/Qg6a11l1gWls6PJiBXrOciOACVuGUoNTzztlCUkCAwEAAaMhMB8wHQYDVR0OBBYEFMm4/1hF4WEPYS5gMXRmmH0gs6XjMA0GCSqGSIb3DQEBCwUAA4IBAQDVH/Md9lCQWxbSbie5lPdPLB72F4831glHlaqms7kzAM6IhRjXmd0QTYq3Ey1J88KSDf8A0HUZefhudnFaHmtxFv0SF5VdMUY14bJ9UsxJ5f4oP4CVh57fHK0w+EaKGGIw6TQEkL5L/+5QZZAywKgPz67A3o+uk45aKpF3GaNWjGRWEPqcGkyQ0sIC2o7FUTV+MV1KHDRuBgreRCEpqMoY5XGXe/IJc1EJLFDnsjIOQU1rrUzfM+WP/DigEQTPpkKWHJpouP+LLrGRj2ziYVbBDveP8KtHvLFsnexA/TidjOOxChKSLT9LYFyQqsvUyCagBb4aLs009kbW6inN8zA6"
	],
	"n": "10iGQ5l5IdqBP1l5wb5BDBZpSyLs4y_Um-kGv_se0BkRkwMZavGD_Nqjq8x3-fKNI45nU7E7COAh8gjn6LCXfug57EQfi0gOgKhOhVcLmKqIEXPmqeagvMndsXWIy6k8WPPwBzSkN5PDLKBXKG_X1BwVvOE9276nrx6lJq3CgNbmiEihovNt_6g5pCxiSarIk2uaG3T3Ve6hUJrM0W35QmqrNM9rL3laPgXtCuz4sJJN3rGnQq_25YbUawW9L1MTVbqKxWiyN5WbXoWUg8to1DhoQnXzDymIMhFa45NTLhxtdH9CDprXWXWBaWzo8mIFes5yI4AJW4ZSg1PPO2UJSQ",
	"e": "AQAB",
	"alg": "RS256"
}`

func ValidateJwt(tokenString string) (string, bool) {
	jwk := map[string]string{}
	json.Unmarshal([]byte(js), &jwk)

	if jwk["kty"] != "RSA" {
		fmt.Println("error")
		return "", false
	}

	// decode the base64 bytes for n
	nb, err := base64.RawURLEncoding.DecodeString(jwk["n"])
	if err != nil {
		fmt.Println("error", err.Error())
		return "", false
	}

	e := 0
	// The default exponent is usually 65537, so just compare the
	// base64 for [1,0,1] or [0,1,0,1]
	if jwk["e"] == "AQAB" || jwk["e"] == "AAEAAQ" {
		e = 65537
	} else {
		// need to decode "e" as a big-endian int
		fmt.Println("error", err.Error())
		return "", false
	}

	pk := &rsa.PublicKey{
		N: new(big.Int).SetBytes(nb),
		E: e,
	}

	der, err := x509.MarshalPKIXPublicKey(pk)
	if err != nil {
		fmt.Println("error", err.Error())
		return "", false
	}

	block := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: der,
	}

	var out bytes.Buffer
	pem.Encode(&out, block)

	block, _ = pem.Decode([]byte(out.String()))
	if block == nil {
		fmt.Println("error", err.Error())
		return "", false
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println("error", err.Error())
		return "", false
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return pub, nil
	})

	// ... error handling
	if err != nil {
		fmt.Println("error", err.Error())
	}

	if claims["sub"] != nil {
		return claims["sub"].(string), token.Valid
	}
	if claims["msisdn"] != nil {
		return claims["msisdn"].(string), true
	}

	fmt.Println("token jwt", claims["msisdn"])
	return "", false
}
func DecryptDigiSubs(input string) (string, error) {
	keytsel := []byte("BdjFKfZJ5lKG9kqaDYrNwNhTilrqOECB")
	a, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		fmt.Println(err.Error())
		return "", errors.New("error")
	}
	h := md5.New()
	h.Write(keytsel)
	key := []byte(string(h.Sum(nil)))
	dst, err := openssl.AesECBDecrypt(a, key, openssl.PKCS7_PADDING)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("ERROR")
	}
	d := string(dst)
	if len(d) >= 1 {
		if d[0] == '"' && d[len(d)-1] == '"' {
			return (d[1 : len(d)-1]), nil
		}
	}
	return (d[1 : len(d)-1]), nil
}

func DecryptTsel(ar *TransformingInformation, key string) (string, error) {
	err := TselDecrypt(ar, key, 2)
	if err != nil {
		return "", err
	}
	return ar.Data, nil

	//mapping_whitelist, err := file.GetFile("/app/file/whitelistdigisubs.json")
	//if err != nil {
	//    return "", err
	//}
	//var mapping_digi_subs map[string]interface{}
	//_ = json.Unmarshal(mapping_whitelist, &mapping_digi_subs)
	//
	//if _, ok := mapping_digi_subs[ar.Data]; ok {
	//    return ar.Data, nil
	//} else {
	//    return "", response.ErrUnAuthorized
	//}

}
