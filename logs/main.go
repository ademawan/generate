package main

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Alamat struct {
	Name     string
	NoTelpon string
}

func main() {
	// trxID := GenerateTransactionId("081221997456")
	// fmt.Println(trxID)
	// test := GenerateTransactionId("081221997456")
	// fmt.Println(test)

	var objek = Alamat{}
	objek.Name = "Test"

	x := createSignatureMytSelPreprod()
	fmt.Println("APYKEY:", x)
	res := hashs(x)
	fmt.Println("x-signature-key", res)
	tc := fmt.Sprintf("%v", &objek)
	fmt.Println(tc)
	Coba(&objek)
	// hash := ValidationHash()
	// hash2 := hashs()

	// fmt.Println(x)
	// fmt.Println(hash)
	// hash2 := md5.Sum([]byte(res))
	// xSignature := hex.EncodeToString(hash2[:])
	// fmt.Println(xSignature)

	// trxIDNew := GenerateTransactionIdNew("081221997456")
	// fmt.Println(trxIDNew)

	TimeTesting()

	GeneratePmbnk()
	fmt.Println(time.Stamp)
	// fmt.Println(GenerateTransactionId("081221997456"))
}
func Coba(a *Alamat) {
	fmt.Println(&a.Name)
}

func GenerateTransactionId(appId, msisdn string) string {
	msisdn_client := msisdn
	loc, _ := time.LoadLocation("Asia/Jakarta")

	timeStamp := time.Now().In(loc).Format("060102150405.000")
	newTimeStamp := strings.Replace(timeStamp, ".", "", 1)
	lastDigit := msisdn_client[len(msisdn_client)-5:]
	changeableDigit := "0"

	generate := appId + newTimeStamp + lastDigit + changeableDigit
	return generate
}
func GenerateTransactionIdNew(msisdn string) string {
	msisdn_client := msisdn
	loc, _ := time.LoadLocation("Asia/Jakarta")
	appId := "OM"
	timeStamp := time.Now().In(loc).Unix()
	newTimeStamp := strconv.Itoa(int(timeStamp))
	lastDigit := msisdn_client[len(msisdn_client)-5:]
	changeableDigit := "0"

	generate := appId + newTimeStamp + lastDigit + changeableDigit

	return generate
}
func createSignatureMytSelPreprod() string {
	// preprod   // API KEY             SECRET
	// createSignature := "eedfevw8q7rym4zhdsa64cma" + "CSkjSURNZA" + strconv.FormatInt(time.Now().Unix(), 10)
	// prod   //
	//4wj7qbrzr8yv57qjby4c9wg7
	createSignature := "MyTselMytselXS4LT" + "Xs3cr3t202211" + strconv.FormatInt(time.Now().Unix(), 10)
	hash := md5.Sum([]byte(createSignature))
	xSignature := hex.EncodeToString(hash[:])

	return xSignature
}
func hashs(apiKey string) string {

	signatureSecret := "S4LTXS3XRET2022@"
	fmt.Println("SecretKey:", signatureSecret)
	secret := signatureSecret
	// apikey adalah "secretkey" yang di generate di endpoint transforming information
	apiKey = "a5685710fdefd048cce666bf9e1b79d1542790094b4f79142752738746a06aac"
	currentTime := time.Now()
	date := currentTime.Format("2006/1/2")
	fmt.Println(date)
	var str string = apiKey + secret + date
	hasher := md5.New()
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}
func ValidationHash() string {

	// 	KEYHASH=awjd9810u29h9u1j2iue1
	// SECRETHASH=zgRjvdgbed

	api_key := os.Getenv("awjd9810u29h9u1j2iue1")
	secret := os.Getenv("zgRjvdgbed")
	currentTime, _ := TimeIn(time.Now(), "Asia/Jakarta")
	date := currentTime.Format("2006/1/2")

	var str string = api_key + secret + date
	hasher := md5.New()
	hasher.Write([]byte(str))
	return hex.EncodeToString(hasher.Sum(nil))
}

func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
		t = t.In(loc)
	}
	return t, err
}

func XSignatureKeyValidate() string {
	// return false
	SIGNATURE_KEY := "awjd9810u29h9u1j2iue1"
	SIGNATURE_SECRET := "zgRjvdgbed"
	api_key := SIGNATURE_KEY
	secret := SIGNATURE_SECRET
	currentTime := time.Now()
	date := currentTime.Format("2006/1/2")

	var str string = api_key + secret + date
	hasher := md5.New()
	hasher.Write([]byte(str))

	return hex.EncodeToString(hasher.Sum(nil))
}

func TimeTesting() {
	layoutFormat := "2006-01-02 15:04:05"
	value := "2015-09-02 08:04:00"
	date, _ := time.Parse(layoutFormat, value)
	fmt.Println(date.String())
	fmt.Println(date)
	t1 := time.Now()
	t2 := t1.Add(time.Minute * 60)
	diff := t2.Sub(t1)
	fmt.Println(diff.Hours())
}

type Team struct {
	TeamName string
	Field    map[string]interface{}
}
type Field struct {
	FieldId    string
	FieldName  string
	FieldValue string
}

func GeneratePmbnk() {

	// 	1. use key-id for test : 'API-Key: b6e1231e-3f15-48f1-aed6-d72986ea1ce9
	// 2. use $username (clientId) = 73e843cd-0b28-4f98-a262-b5dbdf9ee549
	// 3. use $password (clientSecret) = 2d45705c-e992-4bd8-b7fe-7d2d1dc6c4c1

	apiKey := "b6e1231e-3f15-48f1-aed6-d72986ea1ce9"
	username := "73e843cd-0b28-4f98-a262-b5dbdf9ee549"
	password := "2d45705c-e992-4bd8-b7fe-7d2d1dc6c4c1"
	clientStaticKey := "WD628b40f7f480f6445279eeb9b76003"
	tm := time.Now()
	fmt.Printf("time1 %v\n", tm)
	r := tm.Format(time.RFC3339)
	fmt.Printf("time1 %v\n", r)
	Autorization := "Basic " + GenerateBasicAuth(username+":"+password)
	temp := "2017-12-09T03:52:01.000+07:00"

	OAUTHTimestamp := r[:19] + temp[19:]
	// fmt.Println(temp)
	// fmt.Println(OAUTHTimestamp)
	// fmt.Println(Autorization)

	OAUTHSignature := ComputeHmac256(apiKey+":"+OAUTHTimestamp+":"+"grant_type=client_credentials", clientStaticKey)
	APIKey := apiKey

	fmt.Println("Authorization ", Autorization)
	fmt.Println("OAUTH-Signature ", OAUTHSignature)
	fmt.Println("OAUTH-Timestamp ", OAUTHTimestamp)
	fmt.Println("API-Key ", APIKey)

	// curl -X POST -k https://api.pbdevtest.com/apiservice/InquiryServices/BalanceInfo_V2/inq -H 'authorizatin: Bearer e2jex8v6CwAhJSBVfXFveXtykKC4o01fDOZ3Q3Qp5vglylwk1NONWQ' -H 'cachecontrol: no-cache' -H 'content-type: application/json' -H 'organizationname: Permata Bank' -H 'permata-signature: YSHxWxEc2rohVDCI8f/H1S3oNBn7l1wf6hNcAscAdb4=' -H 'permata-timestamp: 2017-11-07T10:22:57' -d '{"BalInqRq":{"MsgRqHdr" {"RequestTimestamp":"2017-07-21T14:32:01+07:00","CustRefID":"0878987654321"},"InqInfo":{"AccountNumber":"701075323"}}}'

	// t := time.Now()
	// fmt.Println(t.Format("20060102150405"))

	// layoutFormat := "20060102150405"
	// value := "21222100805788"
	// date, _ := time.Parse(layoutFormat, value)
	// fmt.Println(value, " sss \t->", date.String())

	signature := ComputeHmac256("b6e1231e-3f15-48f1-aed6-d72986ea1ce9:"+OAUTHTimestamp+":grant_type=client_credentials", clientStaticKey)

	fmt.Println("SIGNATURE", signature)
}

func ComputeHmac256(message string, secret string) string {

	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
func GenerateBasicAuth(basicAuth string) string {
	return base64.StdEncoding.EncodeToString([]byte(basicAuth))
}
