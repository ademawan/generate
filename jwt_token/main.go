package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

type ClaimData struct {
	UserId   string `json:"user_id"`
	UserRole string `json:"user_role"`
}
type User struct {
	UserId    string
	UserNama  string
	UserEmail string
	UserRole  string
}

const JWT_SECRET = "secret"

func main() {

	user := User{
		UserId:    "ID001",
		UserNama:  "Ade Mawan",
		UserEmail: "ademawan1210@gmail.com",
		UserRole:  "1",
	}

	token, err := GenerateToken(user)
	token = "Bearer " + token
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(token)
	valid, t, claims, err := CheckToken(token)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(valid)
	fmt.Println(t)
	fmt.Println(claims)

	if claims["user_id"] != nil {
		fmt.Println(claims["user_id"].(string))
	}

	fmt.Println(generateTokenPair())
}
func GenerateToken(u User) (string, error) {
	if u.UserId == "" {
		return "cannot Generate token", errors.New("user_id == null")
	}

	codes := jwt.MapClaims{
		"user_id":   u.UserId,
		"user_role": u.UserRole,
		// "email":    u.Email,
		// "password": u.Password,
		"exp":  time.Now().Add(time.Hour * 1).Unix(),
		"auth": true,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, codes)
	// fmt.Println(token)
	return token.SignedString([]byte(JWT_SECRET))
}

// func ExtractTokenUserUid(token string) string {
// 	fmt.Println("TOKEN DI PROSES")
// 	user := e.Get("user").(*jwt.Token) //convert to jwt token from interface
// 	fmt.Println("USER", user)
// 	if user.Valid {
// 		codes := user.Claims.(jwt.MapClaims)
// 		id := codes["user_id"].(string)
// 		return id
// 	}
// 	return ""
// }

// func ExtractRoles(e echo.Context) string {
// 	var id string
// 	user := e.Get("user").(*jwt.Token) //convert to jwt token from interface
// 	if user.Valid {
// 		codes := user.Claims.(jwt.MapClaims)
// 		id = codes["roles"].(string)
// 		return id
// 	}
// 	return id
// }

func CheckToken(token string) (b bool, t *jwt.Token, m jwt.MapClaims, err error) {
	// skey, _ := web.AppConfig.String("SECRET_KEY")

	var skey_db = JWT_SECRET
	//ambil secret_key dari database

	kv := strings.Split(token, " ")
	if len(kv) != 2 || kv[0] != "Bearer" {
		// err := errors.New("error, token cannot be loaded")
		return false, nil, nil, nil
	}

	claims := jwt.MapClaims{}

	t, err = jwt.ParseWithClaims(kv[1], claims, func(*jwt.Token) (interface{}, error) {

		return []byte(skey_db), nil
	})
	if err != nil {
		return false, nil, nil, errors.New("error parse token")
	}

	// cek db token
	var newClaim ClaimData

	bytes, err := json.Marshal(claims)
	if err != nil {
		return false, nil, nil, nil
	}
	json.Unmarshal(bytes, &newClaim)

	m = claims

	return true, t, m, nil
}

// with refresh token
func generateTokenPair() (map[string]string, error) {
	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	// This is the information which frontend can use
	// The backend can also decode the token and get admin etc.
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = 1
	claims["name"] = "Jon Doe"
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Minute * 15).Unix()

	// Generate encoded token and send it as response.
	// The signing string should be secret (a generated UUID works too)
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["sub"] = 1
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	rt, err := refreshToken.SignedString([]byte("secret"))
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  t,
		"refresh_token": rt,
	}, nil
}
