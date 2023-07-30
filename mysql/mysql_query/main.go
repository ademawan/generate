package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

	IsOwner bool `json:"is_owner"`
}

func main() {
	var usersInfo = []UserInfo{}

	dbURL := "root:root@tcp(localhost:3308)/tcc?parseTime=true"
	db, err := sql.Open("mysql", dbURL)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	cipher := "s4lTc1ph3RtelK0ms3L@"
	query := `select merchant_id as merchantID,role_id as roleID from merchant_user where user_email=AES_ENCRYPT(?,'%[1]s')`
	query = fmt.Sprintf(query, cipher)
	rows1, err := db.Query(query, "ademawan@gmail.com")
	if err != nil {
		fmt.Println("ERROR 1" + err.Error())
	}
	defer rows1.Close()
	for rows1.Next() {
		fmt.Println("s")
		userInfo := UserInfo{}

		var merchantID int
		var roleID string
		var isMerchant = true

		err := rows1.Scan(&merchantID, &roleID)
		if err != nil {
			fmt.Println(err.Error())
		}

		userInfo.MerchantID = merchantID
		userInfo.RoleID = roleID
		userInfo.IsMerchant = isMerchant

		usersInfo = append(usersInfo, userInfo)

	}
	fmt.Println(usersInfo)

}

// func (aum AuthBasicMysql) GetUserInfoByUUID(uuid string) ([]domain_auth_merchant.UserInfo, error) {
// 	var usersInfo = []domain_auth_merchant.UserInfo{}
// 	cipher := os.Getenv("CHIPER_MYSQL")

// 	// query := helper.GetQueryUserInfoByUUID()
// 	queryGetUuid := `SELECT AES_DECRYPT(email,'%[1]s') FROM users  where uuid=?`
// 	queryGetUuid = fmt.Sprintf(queryGetUuid, cipher)
// 	var email string
// 	err := aum.db.QueryRow(queryGetUuid, uuid).Scan(&email)
// 	if err != nil {
// 		if err != sql.ErrNoRows {
// 			helper.StringLog("error", "REPOSITORY|USERS|GetUserInfo ERROR QueryRow get merchant user info "+err.Error())
// 			return usersInfo, exceptions.ErrSystem
// 		} else {
// 			return usersInfo, exceptions.ErrNotFound
// 		}
// 	}
// 	query := `select merchant_id as merchantID,role_id as roleID from merchant_user where user_email=AES_ENCRYPT(?,'%[1]s')`
// 	query = fmt.Sprintf(query, cipher)

// 	rows1, err := aum.db.Query(query, email)
// 	if err != nil {
// 		helper.StringLog("error", "REPOSITORY|USERS|GetUserInfo ERROR QueryRow get merchant user info "+err.Error())
// 		return usersInfo, exceptions.ErrSystem
// 	}
// 	defer rows1.Close()
// 	if rows1.Next() {
// 		userInfo := domain_auth_merchant.UserInfo{}

// 		var merchantID int
// 		var roleID string
// 		var isMerchant = true

// 		rows1.Scan(&merchantID, &roleID)
// 		if roleID == "1" {
// 			userInfo.IsOwner = true
// 			userInfo.Permissions = append(userInfo.Permissions, "OWNER")
// 		} else if roleID == "0" || roleID == "" {
// 			userInfo.Permissions = append(userInfo.Permissions, "USER")
// 		} else {
// 			permissions := []string{}
// 			query = `SELECT permission_id FROM roles_permissions WHERE role_id =?`
// 			rows2, err := aum.db.Query(query, roleID)
// 			if err != nil {
// 				helper.StringLog("error", "REPOSITORY|USERS|GetUserInfo ERROR Get Permission "+err.Error())
// 				return usersInfo, exceptions.ErrSystem
// 			}
// 			defer rows2.Close()
// 			for rows2.Next() {
// 				var permission string

// 				err = rows2.Scan(&permission)

// 				if err != nil {
// 					helper.StringLog("error", "REPOSITORY|USERS|GetUserInfo ERROR getPermission|Scan|"+err.Error())
// 					return usersInfo, exceptions.ErrSystem
// 				}
// 				permissions = append(permissions, permission)
// 			}
// 			userInfo.Permissions = permissions
// 		}
// 		userInfo.MerchantID = merchantID
// 		userInfo.RoleID = roleID
// 		userInfo.IsMerchant = isMerchant
// 		usersInfo = append(usersInfo, userInfo)

// 	}

// 	// userInfo["permissions"] = permissionsString
// 	// userInfo["authorize"] = authorizeString
// 	// userInfo["invoice"] = invoiceString
// 	// defer helper.Logs(&helper.Log{
// 	// 	Event:      "INTERNAL|DELIVERY|REPOSITORY|MYSQL|GetUserInfoByUUID",
// 	// 	StatusCode: http.StatusOK,
// 	// 	Request:    fmt.Sprintf("DATA_LENGTH:%v", len(usersInfo)),
// 	// 	URL:        "<nil>",
// 	// 	Response:   usersInfo,
// 	// }, "info")

// 	return usersInfo, nil
// }
