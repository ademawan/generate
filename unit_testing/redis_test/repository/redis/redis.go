package repository

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	redisV9 "github.com/redis/go-redis/v9"

	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
)

var (
	scope = "INTERNAL|REPOSITORY|REDIS|AUTH|"
)

type ResponseUsecaseSubs struct {
	CustName   string `json:"cust_name"`
	CustMsisdn string `json:"cust_msisdn"`
	CustType   string `json:"cust_type"`
	XToken     string `json:"x_token"`
	Email      string `json:"email"`
}

type AuthRepository struct {
	redis   *redis.Client
	redisV9 *redisV9.Client
}

func NewAuthRepository(redis *redis.Client, redisV9 *redisV9.Client) *AuthRepository {
	return &AuthRepository{redis: redis, redisV9: redisV9}
}

func (r *AuthRepository) InsertSeamless(custParam string, request *ResponseUsecaseSubs) {
	envKey := os.Getenv("REDIS_KEY_AUTH")

	key := envKey + ":" + custParam
	ttl, _ := strconv.Atoi(os.Getenv("REDIS_TTL_AUTH"))
	var m = make(map[string]interface{})
	m["msisdn"] = request.CustMsisdn
	m["email"] = request.Email
	m["cust_name"] = request.CustName
	m["x_token"] = request.XToken
	m["cust_type"] = request.CustType

	err := r.redis.HMSet(key, m).Err()
	if err != nil {

		log.Println("error set redis", err)

	}
	err = r.redis.Do("EXPIRE", key, ttl).Err()

	if err != nil {
		log.Println("error set expire", err)

	}

}

func (r *AuthRepository) GetSeamless(request string) *ResponseUsecaseSubs {
	envKey := os.Getenv("REDIS_KEY_AUTH")

	key := envKey + ":" + request

	dataRedis, err := r.redis.HGetAll(key).Result()
	if err != nil {

		log.Println("error get inqury redis", err)
		return nil
	}

	if len(dataRedis) != 0 {

		response := &ResponseUsecaseSubs{
			CustMsisdn: dataRedis["msisdn"],
			Email:      dataRedis["email"],
			CustName:   dataRedis["cust_name"],
			XToken:     dataRedis["x_token"],
			CustType:   dataRedis["cust_type"],
		}
		return response
	}

	log.Println("error get inqury redis", err)
	return nil
}
func (r *AuthRepository) Set(key string, value interface{}) error {

	ttl, err := strconv.Atoi(os.Getenv("TTL_DURATION_REDIS"))
	if err != nil {
		logrus.Errorf("error submit fmc msisdn %s", err)
		return err
	}
	err = r.redis.Set(key, value, time.Duration(ttl)*time.Second).Err()

	if err != nil {
		fmt.Println("err redis", err)
		return err
	}

	return nil
}
func (r *AuthRepository) HMSet(key string, ttlDuration int, useTTL bool, value ...string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := r.redisV9.HMSet(ctx, key, value).Err()

	if err != nil {

		return err
	}
	if useTTL {
		err = r.redis.Do("EXPIRE", key, ttlDuration).Err()

		if err != nil {

			return err
		}
	}

	return nil
}
func (r *AuthRepository) HGetAll(key string) (map[string]string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	res := r.redisV9.HGetAll(ctx, key)
	if len(res.Val()) == 0 {
		fmt.Println("redis len(res.Val()) == 0 ")
	}
	return res.Val(), nil
}
func (r *AuthRepository) Get(key string) (map[string]string, error) {

	dataRedis, err := r.redis.HGetAll(key).Result()
	if err != nil {

		log.Println("error get inqury redis", err)
		return dataRedis, err
	}

	return dataRedis, nil
}
