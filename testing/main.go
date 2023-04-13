package main

import (
	"crypto/tls"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

var mySigningKey = []byte("wectelkomselcom")

type FormatRequestMigrationInfo struct {
	Msisdn string `json:"msisdn"`
	Token  string `json:"token_otp"`
	AuthId string `json:"authId"`
}

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func main() {
	fmt.Println("time:", time.Duration(100)*time.Second)
	t := time.Duration(20) * time.Second
	fmt.Println("time:", t)
	str := "([]Golan-g@%Progra ms#"
	str = regexp.MustCompile(`[^a-zA-Z0-9 ]+`).ReplaceAllString(str, "")
	fmt.Println(str)
	text := "jaga pulsa/jaga tagihan"
	lang := "en"
	if lang == "id" {
		text = strings.Replace(text, "&", "dan", -1)
		text = strings.Replace(text, "/", " atau ", -1)

	} else if lang == "en" {
		text = strings.Replace(text, "&", "and", -1)
		text = strings.Replace(text, "/", " or ", -1)
	}
	fmt.Println(text)

	host := "103.13.207.248"
	port := "5432"
	user := "postgres"
	password := "X2023RoG@1"
	dbname := "rogerdev_project20230107_db"

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	result, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Tidak Konek DB Errornya : %s", err)
	}
	fmt.Println("connecr")
	defer result.Close()
	fmt.Println(result)
	res := &Response{}
	tc := &FormatRequestMigrationInfo{}
	// tc.Token = "kldksdksd"
	res.Message = "ok"
	// res.Data = make(map[string]interface{})
	if tc.Token != "" {
		res.Data = tc
	}
	fmt.Println(res)
	file, _ := json.MarshalIndent(res, "", " ")

	_ = ioutil.WriteFile("test.json", file, 0644)

	// xToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJjdXN0X3R5cGUiOiJQcmVwYWlkIiwibXNpc2RuIjoiNjI4MTIyMTk5NzQ1NiIsInVzZXIiOiJ3ZWMgdGVsa29tc2VsIn0.MKwlXwFjJCa7cL3ZiiQR5wjsfgbhSntmDxIp2j2zZ0M"
	// point, err := GetPoint(xToken)
	// fmt.Println(point)

	checked := CheckIsNew("2023-02-07T05:57:56Z")
	fmt.Println(checked)

	// cekOrderDate, status := CheckOrderDate("2023-02-04 23:44:00", 60)
	// fmt.Println(cekOrderDate, status)

	//=====================testing pointer================

}
func ValidateJWT(accessToken string) (jwt.MapClaims, bool) {

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {

			return nil, errors.New("not valid")
		}

		return mySigningKey, nil
	})

	if err != nil {

		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		str, _ := json.Marshal(claims)
		fmt.Println(string(str))
		return claims, true

	} else {

		return nil, false
	}

}
func CreateHTTPClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 100,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(30) * time.Second,
	}

	return client
}

type PointProfileResponse struct {
	CustomerName string `json:"customer_name"`
	Poin         int    `json:"poin"`
	Tier         string `json:"tier"`
	Msisdn       string `json:"msisdn"`
}

func GetPoint(xToken string) (*PointProfileResponse, error) {
	// start := time.Now()
	point := new(PointProfileResponse)
	client := CreateHTTPClient()
	// claims, validateToken := helper.HelperValidateJWT(accessToken)

	// ch, err := sr.mb.Channel()
	baseUrl := "https://43.255.196.44/api/poins/detail/profile"
	///merchant-panel/store-page-service/{merchantslug}
	req, err := http.NewRequest("GET", baseUrl, nil)
	fmt.Println(baseUrl)
	if err != nil {

		return point, err
	}
	fmt.Println(xToken)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-TOKEN", xToken)

	res, err := client.Do(req)
	if err != nil {

		return point, err
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {

		return point, err
	}

	io.Copy(ioutil.Discard, res.Body)
	defer res.Body.Close()

	json.Unmarshal([]byte(body), &point)
	if res.Status == "404" {

		return point, errors.New("Poin [Profile Is Not Found")
	}
	// helper1.HttpService("info", "POST",	 string(payloadJSON), http.StatusOK, "Success", time.Since(start), Url, string(body), nil)
	return point, err

}
func CheckIsNew(createdAt string) bool {
	var status bool
	layoutFormat := "2006-01-02T15:04:05Z07:00"
	timeCreate, _ := time.Parse(layoutFormat, createdAt)
	fmt.Println(timeCreate.Local())
	timeNow := time.Now().Format(time.RFC3339)
	timeNowNew, _ := time.Parse(layoutFormat, timeNow)
	fmt.Println(timeNowNew)
	diff := timeNowNew.Sub(timeCreate)
	if diff.Hours() > float64(24) {
		fmt.Println("diff:", diff.Hours())

		return status
	}
	fmt.Println("diff:", diff.Hours())
	status = true
	return status
}
func CheckOrderDate(orderDate string, second int) (int, bool) {
	fmt.Println("second", second)
	var status bool
	layoutFormat := "2006-01-02 15:04:05"
	timeCreate, _ := time.Parse(layoutFormat, orderDate)
	timeNow := time.Now().String()
	fmt.Println(timeNow[:19])
	timeNowModify, _ := time.Parse(layoutFormat, timeNow[:19])

	diff := timeNowModify.Sub(timeCreate)
	if int(diff.Seconds()) > int(time.Second)*second {
		status = true
		return int(diff.Seconds()), status
	}

	return int(diff.Seconds()), status
}

func TimeTransaction(montMinus int) (string, string) {

	format := "02-January-2006"

	timeNow := time.Now().AddDate(0, -montMinus, 0)

	times := timeNow.Format(format)

	sp := strings.Split(times, "-")

	join := fmt.Sprintf("%s %s", sp[1], sp[2])
	return join, sp[0]

}
