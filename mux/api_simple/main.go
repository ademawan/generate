package main

import (
	user_controller "api_simple/delivery/controllers/user"
	route "api_simple/delivery/routes"
	user_usecase "api_simple/delivery/usecase/user"
	redis_repository "api_simple/repository/redis"

	"api_simple/middleware"
	"os"

	"context"
	"crypto/tls"
	"log"
	"net/http"
	"time"

	redisV9 "github.com/redis/go-redis/v9"

	"github.com/gorilla/mux"
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
	redisRepo := redis_repository.NewRedisRepository(redisClientDB1)

	userUsecase := user_usecase.NewUserUsecase(redisRepo)

	userController := user_controller.NewUserController(userUsecase)
	customMiddleware := middleware.NewCustomMiddleware(redisClientDB1)
	r := mux.NewRouter()

	route.RegisterPath(r, userController, customMiddleware)

	server := &http.Server{Addr: ":" + "8092", Handler: r, TLSConfig: &tls.Config{InsecureSkipVerify: true}}
	log.Printf("starting service on port %s", "8092")
	server.ListenAndServe()
}

const (
	MaxIdleConns       int  = 100
	MaxIdleConnections int  = 100
	RequestTimeout     int  = 30
	SSL                bool = true
)

func createHTTPClient() *http.Client {
	tr := &http.Transport{
		TLSClientConfig:     &tls.Config{InsecureSkipVerify: SSL},
		MaxIdleConns:        MaxIdleConns,
		MaxIdleConnsPerHost: MaxIdleConnections,
	}
	client := &http.Client{
		Transport: tr,
		Timeout:   time.Duration(RequestTimeout) * time.Second,
	}

	return client
}
