package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type (
	UserRequestCreate struct {
		Email     string  `json:"email"`
		Name      string  `json:"name"`
		Username  string  `json:"username"`
		Phone     string  `json:"phone"`
		Password  string  `json:"password"`
		Photo     string  `json:"photo"`
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	MerchantRequestCreate struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
		Password string `json:"password"`
		OtpToken string `json:"otp_token"`
	}
	ResponseSubscriberProfile struct {
		Profiles Profiles `json:"profiles"`
	}
	Profiles struct {
		CustomerType string `json:"customer_type"`
	}
)

func main() {
	users := []*UserRequestCreate{}
	user := UserRequestCreate{}

	user.Email = "ademawan@gmail.com"
	user.Name = "Ade Mawan"
	user.Username = "ademawan1122"
	user.Phone = "081221997456"
	user.Password = "xyz"
	user.Photo = "phototest"
	user.Latitude = -6.580752527500459
	user.Longitude = 106.72555696702857
	// users = append(users, user)

	file, _ := json.MarshalIndent(user, "", " ")

	_ = ioutil.WriteFile("file/test.go", file, 0644)
	data := UserRequestCreate{}
	GetJson("", "test", &data)
	fmt.Println(data)

	dataTest := ResponseSubscriberProfile{}
	dataTest.Profiles.CustomerType = "Telkomsel"
	file2, _ := json.MarshalIndent(dataTest, "", " ")

	_ = ioutil.WriteFile("file/datatest.json", file2, 0644)

	data2 := ResponseSubscriberProfile{}
	GetJson("", "datatest", &data2)
	fmt.Println(data2, data2.Profiles.CustomerType)

	file3, _ := json.MarshalIndent(users, "", " ")

	_ = ioutil.WriteFile("file/testpointer.json", file3, 0644)

	GetJson("", "datatest", &dataTest)
	fmt.Println(dataTest.Profiles.CustomerType)

	d, err := Queue(user, "name")
	if err != nil {
		fmt.Println(err.Error())
	}
	file4, _ := json.MarshalIndent(d, "", " ")

	_ = ioutil.WriteFile("file/d.json", file4, 0644)
}

func GetJson(baseDir, fileName string, data interface{}) {
	jsonFile, err := os.Open(baseDir + "file/" + fileName + ".json")
	if err != nil {

		fmt.Println(err.Error())
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err.Error())
	}

	json.Unmarshal(byteValue, &data)

}

func Queue(dataQueueMap interface{}, routingKey string) (*UserRequestCreate, error) {
	// ch, err := conn.Channel()
	// if err != nil {
	// 	failOnErrorQueue(err, "", routingKey, "Failed to open a channel")

	// }
	reqBody, _ := json.Marshal(dataQueueMap)
	byteBuff := bytes.NewBuffer(reqBody)

	userRequestCreate := UserRequestCreate{}

	err := json.Unmarshal([]byte(byteBuff.Bytes()), &userRequestCreate)
	if err != nil {
		fmt.Println(err.Error())
		return &userRequestCreate, err
	}

	if err != nil {
		failOnErrorQueue(err, "", routingKey, "Failed to open a channel")
		return &userRequestCreate, err
	}

	return &userRequestCreate, nil
}
func failOnErrorQueue(err error, exchange string, routingName string, msg string) {
	if err != nil {
		fmt.Println("error")
	}
}
