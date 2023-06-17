package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	domain "merchant_panel_force_logout/domain/auth_merchant_panel"
	domain_redis "merchant_panel_force_logout/domain/redis"

	"merchant_panel_force_logout/helper"
	"merchant_panel_force_logout/helper/exceptions"
	"merchant_panel_force_logout/redis_revosotory"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	redisV9 "github.com/redis/go-redis/v9"
)

var (
	RedisClientConnV7 = RedisClientV7
)

// RedisClient function
func RedisClientV7(r *redisV9.Client) (*redisV9.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	pong, err := r.Ping(ctx).Result()

	if err != nil {
		return nil, err
	}

	if pong != "" {
		log.Println(pong, " from redis")
		return r, nil
	}

	return nil, err
}
func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("error load env|" + err.Error())
		panic(err.Error())
	}

	redisConnDB01 := redisV9.NewClient(&redisV9.Options{
		Addr:     os.Getenv("REDIS"),
		Password: os.Getenv("REDIS_PASS"),
		DB:       1,
	})
	redisClientDB1, err := RedisClientConnV7(redisConnDB01)
	if err != nil {
		panic(err.Error())
	}
	defer redisClientDB1.Close()
	redisRepo := redis_revosotory.NewRedisRepository(redisClientDB1)
	authMerchantController := NewAuthMerchantPanel(redisRepo)
	controller := NewAuthController(redisRepo, authMerchantController)

	r := mux.NewRouter()
	port := os.Getenv("SERVER_PORT")

	resetPasswordRouter := r.PathPrefix("/auth/reset/password").Subrouter()

	resetPasswordRouter.HandleFunc("", controller.RequestRessetPassword).Methods("POST")
	resetPasswordRouter.HandleFunc("/validate", controller.ValidateRequestReset).Methods("POST")
	resetPasswordRouter.HandleFunc("/submit", controller.SubmitNwwValue).Methods("POST")
	http.ListenAndServe(":"+port, r)

}

type AuthController struct {
	redis        *redis_revosotory.RedisRepository
	authMerchant *AuthMerchantPanel
}

func NewAuthController(redisRepo *redis_revosotory.RedisRepository, authMerchant *AuthMerchantPanel) *AuthController {

	return &AuthController{redisRepo, authMerchant}

}

func (a *AuthController) RequestRessetPassword(w http.ResponseWriter, r *http.Request) {
	uuid := r.Header.Get("uuid")
	email := r.Header.Get("email")

	res := &domain.LoginBasic{}
	res.StatusCode = 200
	res.Token = "res_TOKEN_from_request_reset"

	key := "auth:merchant_panel:password:request:" + res.Token
	duration, _ := strconv.Atoi(os.Getenv("TTL_DURATION_REDIS_REQUEST"))

	err := a.redis.SetString(key, domain_redis.RedisValue{Increment: 1, Email: email, UUID: uuid}, duration)
	if err != nil {
		fmt.Println("ERROR :", err.Error())
	}
	err = a.CetakGetString(key)
	if err != nil {
		fmt.Println("ERRORR CETAK STRING :", err.Error())
	}
	helper.RespondWithJSON(w, http.StatusOK, "SUCCESS")
}
func (a *AuthController) ValidateRequestReset(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")
	redisValue := &domain_redis.RedisValue{}
	keyRequest := "auth:merchant_panel:password:request:" + token
	res := &domain.LoginBasic{}
	res.StatusCode = 200
	res.Token = "res_TOKEN_from_validate_reset"

	resRedis, err := a.redis.GetString(keyRequest)
	if err != nil {
		helper.RespondWithJSON(w, http.StatusBadRequest, err.Error())

		return
	}
	err = json.Unmarshal([]byte(resRedis), &redisValue)
	if err != nil {
		data := &helper.HTTPLogData{
			Level: "error",
			// Service: event + "JSON_UNMARSHAL",
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("ERROR : %v | UUID %v", err.Error(), redisValue.UUID),
			Err:     err,
		}
		helper.HTTPLog(data)
		helper.RespondWithJSON(w, http.StatusBadRequest, err.Error())

		return
	}
	key := "auth:merchant_panel:password:validate:" + res.Token
	duration, _ := strconv.Atoi(os.Getenv("TTL_DURATION_REDIS_REQUEST"))

	err = a.redis.SetString(key, domain_redis.RedisValue{Increment: 1, Email: redisValue.Email, UUID: redisValue.UUID}, duration)
	if err != nil {
		fmt.Println("ERROR :", err.Error())
	}
	err = a.CetakGetString(key)
	if err != nil {
		fmt.Println("ERRORR CETAK STRING :", err.Error())
	}
	helper.RespondWithJSON(w, http.StatusOK, "Success")
}

func (a *AuthController) SubmitNwwValue(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("token")

	key := "auth:merchant_panel:password:validate:" + token
	fmt.Println(token, "|", key)
	resRedis, err := a.redis.GetString(key)
	if err != nil {
		if err == redisV9.Nil {
			data := &helper.HTTPLogData{
				Level: "error",
				// Service: event + "JSON_UNMARSHAL",
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("ERROR : %v | Token %v", err.Error(), token),
				Err:     err,
			}
			helper.HTTPLog(data)
			helper.RespondWithJSON(w, http.StatusUnauthorized, err.Error()+"|HERE")

			return
		}
		data := &helper.HTTPLogData{
			Level: "error",
			// Service: event + "JSON_UNMARSHAL",
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("ERROR : %v | Token %v", err.Error(), token),
			Err:     err,
		}
		helper.HTTPLog(data)
		helper.RespondWithJSON(w, http.StatusBadRequest, err.Error())

		return
	}
	_, err = a.redis.Del(key)
	if err != nil {
		data := &helper.HTTPLogData{
			Level: "error",
			// Service: event + "JSON_UNMARSHAL",
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("ERROR : %v | Token %v", err.Error(), token),
			Err:     err,
		}
		helper.HTTPLog(data)
		helper.RespondWithJSON(w, http.StatusBadRequest, err.Error())

		return
	}

	redisValue := &domain_redis.RedisValue{}
	err = json.Unmarshal([]byte(resRedis), &redisValue)
	if err != nil {
		data := &helper.HTTPLogData{
			Level: "error",
			// Service: event + "JSON_UNMARSHAL",
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("ERROR : %v | Token %v", err.Error(), token),
			Err:     err,
		}
		helper.HTTPLog(data)
		helper.RespondWithJSON(w, http.StatusBadRequest, err.Error())

		return
	}
	sessionKey := os.Getenv("MERCHANT_SESSION_KEY") + "294" + redisValue.UUID
	duration, _ := strconv.Atoi(os.Getenv("TTL_DURATION_REDIS_REQUEST"))

	a.redis.SetString(sessionKey+":001", redisValue, duration)
	a.redis.SetString(sessionKey+":002", redisValue, duration)
	a.redis.SetString(sessionKey+":003", redisValue, duration)
	a.CetakGetString(sessionKey + ":001")
	a.CetakGetString(sessionKey + ":002")
	a.CetakGetString(sessionKey + ":003")
	status, err := a.authMerchant.ForceLogout(sessionKey)
	if err != nil {
		fmt.Println("ERROR FORCE LOGOUT | ERROR :", err.Error())
	}
	fmt.Println(status)

	a.CetakGetString(sessionKey + ":001")
	a.CetakGetString(sessionKey + ":002")
	a.CetakGetString(sessionKey + ":003")
	helper.RespondWithJSON(w, http.StatusOK, "SUCCESS SUBMIT")
}
func (a *AuthController) CetakGetString(key string) error {
	resRedis, err := a.redis.GetString(key)
	if err != nil {
		if err == redisV9.Nil {

			return err
		}
		return err
	}

	redisValue := &domain_redis.RedisValue{}
	err = json.Unmarshal([]byte(resRedis), &redisValue)
	if err != nil {

		return err
	}
	fmt.Println(redisValue, "|KEY :", key)
	return nil
}

type AuthMerchantPanel struct {
	redis *redis_revosotory.RedisRepository
}

var (
	scope = "INTERNAL|USECASE|AUTH_MERCHANT_PANEL|"
)

// NewAuthMerchantPanel initial function
func NewAuthMerchantPanel(redis *redis_revosotory.RedisRepository) *AuthMerchantPanel {
	return &AuthMerchantPanel{redis}
}

func (p *AuthMerchantPanel) ValidateAccessToken(payload *domain.AuthMerchantPanelPayload) (*domain.AuthMerchantPanelResponse, error) {
	event := scope + "VALIDATE_ACCESS_TOKEN|"
	//auth:merchant_panel:session:{email}:{device_id}:token_fdksakjfnwfewkfnwlkfnkafnwkfnkenfwekfnwafne
	sessionKey := os.Getenv("MERCHANT_SESSION_KEY")

	key := sessionKey + payload.MerchantID + ":" + payload.UUID + ":" + payload.DeviceID

	res, err := p.redis.GetString(key)
	if err != nil {
		if err == redisV9.Nil {
			return nil, exceptions.ErrUnauthorized
		}
		return nil, exceptions.ErrSystem
	}

	userInfo := &domain.UserInfo{}
	err = json.Unmarshal([]byte(res), &userInfo)
	if err != nil {
		data := &helper.HTTPLogData{
			Level:   "error",
			Service: event + "JSON_UNMARSHAL",
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("ERROR : %v | UUID %v", err.Error(), payload.UUID),
			Err:     err,
		}
		helper.HTTPLog(data)
		return nil, exceptions.ErrSystem
	}

	if userInfo.AccessToken != payload.AccessToken {
		return nil, exceptions.ErrUnauthorized
	}
	responseData := &domain.AuthMerchantPanelResponse{}
	responseData.UUID = userInfo.UUID
	responseData.MerchantID = userInfo.MerchantID
	responseData.DeviceID = userInfo.DeviceID
	responseData.RoleID = userInfo.RoleID
	responseData.Permissions = userInfo.Permissions

	return responseData, nil
}

func (p *AuthMerchantPanel) RevokeToken(payload *domain.AuthMerchantPanelPayload) (bool, error) {
	sessionKey := os.Getenv("MERCHANT_SESSION_KEY")

	key := sessionKey + payload.MerchantID + ":" + payload.UUID + ":" + payload.DeviceID

	_, err := p.redis.Del(key)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (p *AuthMerchantPanel) ForceLogout(key string) (bool, error) {
	event := scope + "FORCE_LOGOUT|"

	keys, err := p.redis.GetKeys(key)
	if err != nil {
		data := &helper.HTTPLogData{
			Level:   "error",
			Service: event + "GetKeys",
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("ERROR : %v |KEYS :%v  ", err.Error(), key),
			Err:     err,
		}
		helper.HTTPLog(data)
	}
	for _, val := range keys {
		_, err = p.redis.Del(key)
		if err != nil {
			data := &helper.HTTPLogData{
				Level:   "error",
				Service: event + "Del",
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("ERROR : %v |KEY :%v  ", err.Error(), val),
				Err:     err,
			}
			helper.HTTPLog(data)
		}
	}
	keysByte, _ := json.Marshal(keys)
	fmt.Println(fmt.Sprintf("KEYS :%v", string(keysByte)))
	data := &helper.HTTPLogData{
		Level:   "info",
		Service: event,
		Status:  http.StatusOK,
		Message: fmt.Sprintf("SUCCESS FORCE LOGOUT SESSION KEY :%v | KEYS :%v ", key, string(keysByte)),
		Err:     err,
	}
	helper.HTTPLog(data)

	return true, nil
}
