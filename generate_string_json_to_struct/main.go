package main

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	MerchantID  int      `json:"merchant_id"`
	RoleID      string   `json:"role_id"`
	Permissions []string `json:"permissions"`
	UUID        string   `json:"uuid"`
	IsMerchant  bool     `json:"is_merchant"`
	CreatedAt   string   `json:"created_at"`
	AccessToken string   `json:"access_token"`
	DeviceID    string   `json:"device_id"`
}

func main() {
	data := &UserInfo{}

	stringJsonData := `{"merchant_id":20}`
	json.Unmarshal([]byte(stringJsonData), &data)
	fmt.Println(data)
}
