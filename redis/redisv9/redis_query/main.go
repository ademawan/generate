package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	redis "github.com/redis/go-redis/v9"
)

type UserInfo struct {
	MerchantID  int    `json:"merchant_id"`
	RoleID      string `json:"role_id"`
	Permissions string `json:"permissions"`
	UUID        string `json:"uuid"`
	IsMerchant  bool   `json:"is_merchant"`
	CreatedAt   string `json:"created_at"`
	AccessToken string `json:"access_token"`
}

// userInfo["merchant_id"] = merchantID
// userInfo["role_id"] = roleID
// userInfo["permissions"] = permissionsString
// userInfo["uuid"] = userData.Uuid
// userInfo["is_merchant"] = userData.IsMerchant
// userInfo["created_at"] = createdAt

// merchants/sessions/{mercahant_id}/{uuid}/{device_id} userinfo
// merchants/refresh/{mercahant_id}/{uuid}/{device_id} userinfo

var (
	getEnv      = os.Getenv
	redisClient *redis.Client

	maxIdleConns       int  = 100
	maxIdleConnections int  = 100
	requestTimeout     int  = 30
	ssl                bool = true
	RedisClientConn         = RedisClient
)

func main() {
	// hashName := "folder_fmc"
	// prefixName := "sso_token"
	key := "merchants:sessions:11:uuidxxxxxx:deviceid01"
	data := UserInfo{}
	data.MerchantID = 11
	data.RoleID = "MP003"
	data.Permissions = "SOS01|SOS02|XOS01"
	data.UUID = "uuid1xxxxxx"
	data.IsMerchant = true
	data.AccessToken = "eyJ0eXAiOiJKV1QiLCJraWQiOiJ3VTNpZklJYUxPVUFSZVJCL0ZHNmVNMVAxUU09IiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiIzMzNmM2I5Yi1jNzI2LTQzYzAtOTQ1ZS03NTQ1YTcyZWY3ZWQiLCJjdHMiOiJPQVVUSDJfU1RBVEVMRVNTX0dSQU5UIiwiYXV0aF9sZXZlbCI6MCwiYXVkaXRUcmFja2luZ0lkIjoiNmYwYzk0YzgtZDMwOC00MzBhLTkyYzAtZTM1NDhlYWFmZWUxLTQxOTYyIiwic3VibmFtZSI6IjMzM2YzYjliLWM3MjYtNDNjMC05NDVlLTc1NDVhNzJlZjdlZCIsImlzcyI6Imh0dHBzOi8vYW06NDQzL2FtL29hdXRoMi90c2VsL3dlYy93ZWIiLCJ0b2tlbk5hbWUiOiJhY2Nlc3NfdG9rZW4iLCJ0b2tlbl90eXBlIjoiQmVhcmVyIiwiYXV0aEdyYW50SWQiOiJ4dHRVT09MUnZpd0Zrc05PR2ZGSkpZSE5YME0iLCJub25jZSI6InRydWUiLCJhdWQiOiJiMzkzNjE4NDM2ZTUxMWVjOGQzZDAyNDJhYzEzMDAwMyIsIm5iZiI6MTY4NDcyNTg5NywiZ3JhbnRfdHlwZSI6ImF1dGhvcml6YXRpb25fY29kZSIsInNjb3BlIjpbIm9wZW5pZCIsInByb2ZpbGUiXSwiYXV0aF90aW1lIjoxNjg0NzI1ODk2LCJyZWFsbSI6Ii90c2VsL3dlYy93ZWIiLCJleHAiOjE2ODQ3MjY3OTcsImlhdCI6MTY4NDcyNTg5NywiZXhwaXJlc19pbiI6OTAwLCJqdGkiOiJuTTMwRy1uVnVTUXJENTFEcGlLS1JuR2UxOHcifQ.j8ts7EG071Zk3m_xIltMhfEEk8TRAe_H0t588EXiE2MzQGCgreO8YvTceAogT_ENMQj9qMOJrSe4A2uP2JB5xroyEdqLSzQIJ8vKuSiJQQ-xNxs8ywAw5OwjzxnhxKofDyr7rraCFN4XFwzGRYuBGLjfokrY8nBhgITLEpRB16gzfslSO9WbvdoQiyuR9OYBqhU2zNahSxzEg3aKtQXuwzCrYx-PzX3ck9wx_DGzNoDs0qdwBU_X0skdcpxp3YjouSraF0vSM4WlXT9qY16qXlE-0ofbNH1TaXiqYBzomWRnexcv31B4Bs2KcVTK3A_lqy6xN0bJ2Bp8_RAel8NH1Q"

	redisClient, err := NewRedis()
	if err != nil {
		panic(err)
	}
	defer redisClient.Close()

	redisRepo := NewRedisRepository(redisClient)
	fmt.Println(key)

	redisRepo.SetString(key+"A3", data, 60)
	res, err := redisRepo.GetString(key + "A3")
	if err != nil {
		if err == redis.Nil {
			fmt.Println(fmt.Sprintf("ERROR GET %v", err.Error()))

		}
		return
	}
	fmt.Println(string(res))
	data2 := UserInfo{}
	err = json.Unmarshal([]byte(res), &data2)
	if err != nil {
		panic(err)
	}
	_, err = redisRepo.GetKeys(`merchants:sessions:11*`)
	if err != nil {
		panic(err)
	}

}

// RedisClient function
func RedisClient(r *redis.Client) (*redis.Client, error) {
	ctx := context.Background()
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
func NewRedis() (*redis.Client, error) {

	// 	//
	var (
		REDIS      = "localhost:6379"
		REDIS_PASS = ""
	)
	// REDIS_TTL=86400
	// REDIS_TTL_REQ_OTP_SECOND=60
	// REDIS_TTL_VALIDATE_OTP_SECOND=300
	// MAX_ATTEMP_REQ_OTP_PER_MSISDN=5
	// MAX_ATTEMP_VALIDATE_OTP_PER_AUTHID=5
	redisConn := redis.NewClient(&redis.Options{
		Addr:     REDIS,
		Password: REDIS_PASS,
		DB:       1,
	})
	redisClient, err := RedisClientConn(redisConn)
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}

type RedisRepository struct {
	redis *redis.Client
}

func NewRedisRepository(redis *redis.Client) *RedisRepository {
	return &RedisRepository{redis: redis}
}

func (r *RedisRepository) GetKeys(prefixKey string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	var cursor uint64
	var keys []string
	for {

		var err error
		keys, cursor, err = r.redis.Scan(ctx, cursor, prefixKey, 0).Result()
		if err != nil {
			return []string{}, err
		}

		for _, key := range keys {
			fmt.Println("key", key)
		}

		if cursor == 0 { // no more keys
			break
		}
	}
	return keys, nil
}

func (r *RedisRepository) GetString(key string) (string, error) {

	contextTimeOut, _ := strconv.Atoi("20")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(contextTimeOut)*time.Second)
	defer cancel()
	dataRedis, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		fmt.Println("ERROR GetString:", err.Error())

		return dataRedis, err
	}

	return dataRedis, nil
}

func (r *RedisRepository) SetString(key string, value interface{}, duration int) error {

	contextTimeOut, _ := strconv.Atoi("20")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(contextTimeOut)*time.Second)
	defer cancel()
	stringValue, _ := json.Marshal(value)
	err := r.redis.Set(ctx, key, stringValue, time.Duration(duration)*time.Second).Err()

	if err != nil {
		fmt.Println("ERROR SetString:", err.Error())
		return err
	}

	return nil
}
func (r *RedisRepository) Del(key string) (int64, error) {
	contextTimeOut, _ := strconv.Atoi("20")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(contextTimeOut)*time.Second)
	defer cancel()

	countDelete, err := r.redis.Del(ctx, key).Result()
	if err != nil {

		fmt.Println("ERROR Del:", err.Error())

		return countDelete, err
	}

	return countDelete, nil
}
