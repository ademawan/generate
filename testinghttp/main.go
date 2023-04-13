package main

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

type FormatRequestMigrationInfo struct {
	Msisdn string `json:"msisdn"`
	Token  string `json:"token_otp"`
	AuthId string `json:"authId"`
}
type PointProfileResponse struct {
	CustomerName string `json:"customer_name"`
	Poin         int    `json:"poin"`
	Tier         string `json:"tier"`
	Msisdn       string `json:"msisdn"`
}

func main() {
	datas, err := GetPoint()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(datas)
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

type Data struct {
	Code    int
	Message string
	Data    interface{}
}

func GetPoint() (*Data, error) {
	// start := time.Now()
	point := new(Data)
	client := CreateHTTPClient()
	// claims, validateToken := helper.HelperValidateJWT(accessToken)
	requestBody, _ := json.Marshal(
		map[string]interface{}{
			"phone": "081221997499",
		},
	)

	// ch, err := sr.mb.Channel()
	baseUrl := "https://103.13.207.248/v1/auth/request/otp"
	///merchant-panel/store-page-service/{merchantslug}

	req, err := http.NewRequest("POST", baseUrl, bytes.NewReader(requestBody))
	fmt.Println(baseUrl)
	if err != nil {

		return point, err
	}
	req.Header.Set("Content-Type", "application/json")
	// req.Header.Set("X-TOKEN", xToken)

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
