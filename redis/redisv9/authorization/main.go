package main

import (
	"context"
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
	hashName := "folder_fmc"
	prefixName := "sso_token"
	key := hashName + ":" + prefixName + "abc"
	redisClient, err := NewRedis()
	if err != nil {
		panic(err)
	}
	defer redisClient.Close()

	redisRepo := NewRedisRepository(redisClient)
	fmt.Println(key)
	items := ExampleUser{"jane", 22}
	Del(key, redisClient)
	err = redisRepo.HandlerHSet(key+"AA", items)
	if err != nil {
		panic(err)
	}
	err = redisRepo.HandlerHSet(key+"BB", ExampleUser{Name: "Mawan", Age: 55})
	if err != nil {
		panic(err)
	}
	fmt.Println("success HSet")
	res2 := redisRepo.HandlerHGetAll(key)
	fmt.Println(len(res2["age"]))
	fmt.Println(res2)

	// }

	// time.Sleep(time.Duration(999) * time.Second)

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

func (r *RedisRepository) Set(key string, value interface{}) error {
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
func (r *RedisRepository) Get(key string) (int64, error) {
	ctx := context.Background()

	dataRedis, err := r.redis.Get(ctx, key).Int64()
	if err != nil {

		log.Println("error get inqury redis", err)
		return dataRedis, err
	}
	return dataRedis, nil
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
func (r *RedisRepository) GetInt(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	dataRedis, err := r.redis.Get(ctx, key).Int64()
	if err != nil {

		log.Println("error get inqury redis", err)
		return dataRedis, err
	}

	return dataRedis, nil
}

type ExampleUser struct {
	Name string `redis:"name"`
	Age  int    `redis:"age"`
}

func (r *RedisRepository) HandlerHSet(key string, val ExampleUser) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := r.redis.HSet(ctx, key, val).Err()
	if err != nil {
		return err
	}

	return nil
}
func (r *RedisRepository) HandlerHGetAll(key string) map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res := r.redis.HGetAll(ctx, key).Val()

	return res
}
