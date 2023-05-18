package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	redis "github.com/redis/go-redis/v9"
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
	for i := 0; i < 100000; i++ {
		fmt.Println(i)
		redisClient, err := NewRedis()
		if err != nil {
			panic(err)
		}
		defer redisClient.Close()

		ress, err := redisClient.Keys(context.Background(), "*sessions*").Result()
		if err != nil {
			panic(err)
		}
		redisClient.Del(context.Background(), "*sessions*")
		fmt.Println(ress[0])
		for _, val := range ress {
			fmt.Println(val)
		}
		fmt.Println("success connect radis")

		redisRepo := NewAuthRepository(redisClient)
		fmt.Println(key)
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
		res, err := redisRepo.GetInt(key)
		if err != nil {
			panic(err)
		}
		fmt.Println(res)

		// r, err := Del(key, redisClient)
		// if err != nil {
		// 	panic(err)
		// }
		// fmt.Println("R:", r)

		res2, err := redisRepo.GetInt(key)
		if err != nil {
			fmt.Println("ERROR2:")
			panic(err)
		}
		fmt.Println(res2)

	}

	time.Sleep(time.Duration(999) * time.Second)

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

type AuthRepository struct {
	redis *redis.Client
}

func NewAuthRepository(redis *redis.Client) *AuthRepository {
	return &AuthRepository{redis: redis}
}

func (r *AuthRepository) Set(key string, value interface{}) error {
	ctx := context.Background()

	ttl, err := strconv.Atoi("999")
	if err != nil {
		return err
	}
	err = r.redis.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()

	if err != nil {
		fmt.Println("err redis", err)
		return err
	}

	return nil
}
func (r *AuthRepository) Get(key string) (int64, error) {
	ctx := context.Background()

	dataRedis, err := r.redis.Get(ctx, key).Int64()
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

	ctx := context.Background()
	txf := func(tx *redis.Tx) error {
		// Get the current value or zero.
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return err
		}

		// Actual operation (local in optimistic lock).
		n++
		if n > 2 {
			return errors.New("tommany request")
		}

		// Operation is commited only if the watched keys remain unchanged.
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n, 0)
			return nil
		})
		return err
	}
	// time.Sleep(time.Duration(5) * time.Second)
	// Retry if the key has been changed.
	for i := 0; i < maxRetries; i++ {
		fmt.Println(i)

		err := rdb.Watch(ctx, txf, key)
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
func Del(key string, redis *redis.Client) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	countDelete, err := redis.Del(ctx, key).Result()
	if err != nil {

		log.Println("error get inqury redis", err)
		return countDelete, err
	}

	return countDelete, nil
}
func (r *AuthRepository) GetInt(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	dataRedis, err := r.redis.Get(ctx, key).Int64()
	if err != nil {

		log.Println("error get inqury redis", err)
		return dataRedis, err
	}

	return dataRedis, nil
}
