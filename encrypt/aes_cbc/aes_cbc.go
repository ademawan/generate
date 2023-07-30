package aes_cbc

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

var (
	passphrase = "3scRLrpd17"
)

func AesCBC() {

	h := sha256.New()
	fmt.Println(h.Size())
	fmt.Println(h)
	count, _ := h.Write([]byte(passphrase))
	fmt.Println("Count:", count)
	key := h.Sum(nil)
	fmt.Printf("[]byte(passphrase):%x\n,EncodeToString:%s\n ByteKey:%x\n,Byte%v\n,Binner %08b\n", []byte(passphrase), hex.EncodeToString(key), key, key, key)

	hx := hex.EncodeToString([]byte(passphrase))
	fmt.Println(fmt.Sprintf("%s", hx))
	f := base64Encode([]byte(passphrase))

	fmt.Printf("%s", f)
}

func Encrypt() {

}
func decodeHex(input []byte) ([]byte, error) {
	db := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(db, input)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func base64Encode(input []byte) []byte {
	eb := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(eb, input)

	return eb
}
