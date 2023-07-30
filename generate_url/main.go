package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

type RequestCreditDetail struct {
	TransactionID    string `json:"transaction_id"`
	Channel          string `json:"channel"`
	PaymentCode      string `json:"payment_code"`
	OrganizationCode string `json:"organization_code"`
	PaymentMethod    string `json:"payment_method"`
}

func main() {
	base, _ := url.Parse("https://api.digitalcore.telkomsel.co.id/")

	base.Path += "preprod-omni/scrt/fulfillment/purchaseoffer/6281221997456"
	base.RawQuery = fmt.Sprintf("interface=%s&paymentmethod=%s&expiry_duration=%s&expiry_uom=%s&msisdn=%s&amount=%s&subscribe=%s", "Web", "OMNICHANNEL", "10", "hour", "6281221997456", "50000", "false")

	base.Query().Set("interface", "Web")
	base.Query().Add("paymentmethod", "OMNICHANNEL")
	base.Query().Add("expiry_duration", "10")
	base.Query().Add("expiry_uom", "hour")
	base.Query().Add("msisdn", "6281221997456")
	base.Query().Add("amount", "50000")
	base.Query().Add("subscribe", strconv.FormatBool(false))
	fmt.Println("BASE_URI:" + base.String())
	urlHasReplaceDinamicParam := SkipDinamicParam(base.Scheme+"://"+base.Host+base.Path, []int{}, []string{})
	fmt.Println("urlHasReplaceDinamicParam:" + urlHasReplaceDinamicParam)
	transaction_id := GenerateTransactionId()
	request := RequestCreditDetail{
		Channel:          os.Getenv("CHANNEL"),
		TransactionID:    transaction_id,
		PaymentCode:      "paymentCode",
		OrganizationCode: "6281221997456",
		PaymentMethod:    "OMNICHANNEL",
	}

	orderReq, err := json.Marshal(&request)

	if err != nil {
		fmt.Println("error")
	}
	fmt.Println(string(orderReq))

}

func GenerateTransactionId() string {
	passport, _ := rand.Prime(rand.Reader, 64)
	loc, _ := time.LoadLocation("Asia/Jakarta")
	appId := os.Getenv("APP_ID")
	timeStamp := time.Now().In(loc).Format("060102150405.000")
	newTimeStamp := strings.Replace(timeStamp, ".", "", 1)
	psprt := fmt.Sprintf("%d", passport)
	lastDigit := psprt[len(psprt)-5:]
	changeableDigit := "0"

	generate := appId + newTimeStamp + lastDigit + changeableDigit
	return generate
}
func SkipDinamicParam(uri string, index []int, value []string) string {
	uriSplit := strings.Split(uri, "/")
	for i, val := range uriSplit {
		x := strconv.Itoa(i)
		fmt.Println(x + ":" + val)

	}
	for i, val := range index {
		if i == val {
			uriSplit[i] = `{` + value[i] + `}`
		}
	}
	return strings.Join(uriSplit, "/")

}
