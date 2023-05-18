package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

func main() {
	fmt.Println(GetSignature())
}
func GetSignature() string {
	createSignature := "eedfevw8q7rym4zhdsa64cma" + "CSkjSURNZA" + strconv.FormatInt(time.Now().Unix(), 10)
	xSignature := GetMD5Hash(createSignature)
	return xSignature
}

func GetMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}
