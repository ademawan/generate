package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	redis "github.com/go-redis/redis"
)

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
	//get folder_fmc:sso_tokenabc
	hashName := "folder_fmc"
	prefixName := "sso_token"
	key := hashName + ":" + prefixName + "abc"
	var (
		REDIS_DB0      = "localhost:6379"
		REDIS_PASS_DB0 = ""
		DB_DB0         = 0
		REDIS_DB1      = "localhost:6379"
		REDIS_PASS_DB1 = ""
		DB_DB1         = 1
	)

	redisClient, err := NewRedis(&redis.Options{
		Addr:     REDIS_DB0,
		Password: REDIS_PASS_DB0,
		DB:       DB_DB0,
	})
	if err != nil {
		panic(err)
	}
	defer redisClient.Close()
	redisClientDB1, err := NewRedis(&redis.Options{
		Addr:     REDIS_DB1,
		Password: REDIS_PASS_DB1,
		DB:       DB_DB1,
	})
	if err != nil {
		panic(err)
	}
	defer redisClientDB1.Close()

	fmt.Println("success connect radis")

	redisRepo := NewAuthRepository(redisClient)
	fmt.Println(key)

	redisRepo2 := NewAuthRepository(redisClientDB1)

	// var data = struct {
	// 	Nama string `json:"nama"`
	// }{
	// 	Nama: "saya",
	// }

	// d, err := json.Marshal(data)
	// if err != nil {
	// 	panic(err)
	// }

	// err = redisRepo.Set(key, 1)
	// if err != nil {
	// 	panic(err)
	// }
	err = increment(key, redisClient)
	if err != nil {
		fmt.Println(err)
	}
	err = increment(key, redisClientDB1)
	if err != nil {
		fmt.Println(err)
	}
	res, err := redisRepo.Get(key)
	if err != nil {
		panic(err)
	}
	fmt.Println("db0:", res)
	res2, err := redisRepo2.Get(key)
	if err != nil {
		panic(err)
	}
	fmt.Println("db01:", res2)

}

// RedisClient function
func RedisClient(r *redis.Client) (*redis.Client, error) {

	pong, err := r.Ping().Result()

	if err != nil {
		return nil, err
	}

	if pong != "" {
		log.Println(pong, " from redis")
		return r, nil
	}

	return nil, err
}
func NewRedis(option *redis.Options) (*redis.Client, error) {

	// 	//

	// REDIS_TTL=86400
	// REDIS_TTL_REQ_OTP_SECOND=60
	// REDIS_TTL_VALIDATE_OTP_SECOND=300
	// MAX_ATTEMP_REQ_OTP_PER_MSISDN=5
	// MAX_ATTEMP_VALIDATE_OTP_PER_AUTHID=5
	redisConn := redis.NewClient(option)
	redisClient, err := RedisClientConn(redisConn)
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}

type AuthRepository struct {
	redis *redis.Client
}

func NewAuthRepository(redis *redis.Client) *AuthRepository {
	return &AuthRepository{redis: redis}
}

func (r *AuthRepository) Set(key string, value interface{}) error {

	ttl, err := strconv.Atoi("999")
	if err != nil {
		return err
	}
	err = r.redis.Set(key, value, time.Duration(ttl)*time.Second).Err()

	if err != nil {
		fmt.Println("err redis", err)
		return err
	}

	return nil
}
func (r *AuthRepository) Get(key string) (int64, error) {

	dataRedis, err := r.redis.Get(key).Int64()
	if err != nil {

		log.Println("error get inqury redis", err)
		return dataRedis, err
	}
	return dataRedis, nil
}

// Increment transactionally increments the key using GET and SET commands.
func increment(key string, rdb *redis.Client) error {
	// Transactional function.
	const maxRetries = 1000

	txf := func(tx *redis.Tx) error {
		// Get the current value or zero.
		n, err := tx.Get(key).Int()
		if err != nil && err != redis.Nil {
			return err
		}

		// Actual operation (local in optimistic lock).
		n++
		if n > 3 {
			return errors.New("tommany request")
		}

		// Operation is commited only if the watched keys remain unchanged.
		_, err = tx.TxPipelined(func(pipe redis.Pipeliner) error {
			pipe.Set(key, n, 0)
			return nil
		})
		return err
	}
	time.Sleep(time.Duration(10) * time.Second)
	// Retry if the key has been changed.
	for i := 0; i < maxRetries; i++ {
		fmt.Println(i)

		err := rdb.Watch(txf, key)
		if err == nil {
			// Success.
			return nil
		}
		if err == redis.TxFailedErr {
			// Optimistic lock lost. Retry.
			continue
		}
		// Return any other error.
		return err
	}

	return errors.New("increment reached maximum number of retries")
}
