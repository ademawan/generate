package redis_revosotory

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	domain_reset_password "merchant_panel_force_logout/domain/reset_password"
	"merchant_panel_force_logout/helper"
	"merchant_panel_force_logout/helper/exceptions"

	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisRepository struct {
	redis *redis.Client
}

var (
	scope = "INTERNAL|DELIVERY|REDIS|REDIS_REPOSITORY|"
)

func NewRedisRepository(redis *redis.Client) *RedisRepository {
	return &RedisRepository{redis: redis}
}

func (r *RedisRepository) Set(key string, value interface{}, ttlDuration string) error {
	event := scope + "Set|"
	contextTimeOut, _ := strconv.Atoi(os.Getenv("REDIS_CONTEXT_TIMEOUT"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(contextTimeOut)*time.Second)
	defer cancel()
	ttl, err := strconv.Atoi(ttlDuration)
	if err != nil {
		data := &helper.HTTPLogData{
			Level:   "error",
			Service: event + "STRCONV.ATOI",
			Status:  http.StatusInternalServerError,
			Message: "failed convert ttlduration",
			Err:     err,
		}
		helper.HTTPLog(data)
		return err
	}
	err = r.redis.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()

	if err != nil {
		data := &helper.HTTPLogData{
			Level:   "error",
			Service: event + "Redis Set",
			Status:  http.StatusInternalServerError,
			Message: "failed Set to redis",
			Err:     err,
		}
		helper.HTTPLog(data)
		return err
	}

	return nil
}
func (r *RedisRepository) SetString(key string, value interface{}, duration int) error {
	event := scope + "SetString|"

	contextTimeOut, _ := strconv.Atoi(os.Getenv("REDIS_CONTEXT_TIMEOUT"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(contextTimeOut)*time.Second)
	defer cancel()
	stringValue, _ := json.Marshal(value)
	err := r.redis.Set(ctx, key, stringValue, time.Duration(duration)*time.Second).Err()

	if err != nil {
		data := &helper.HTTPLogData{
			Level:   "error",
			Service: event + "Redis Set",
			Status:  http.StatusInternalServerError,
			Message: "failed Set to redis",
			Err:     err,
		}
		helper.HTTPLog(data)
		return err
	}

	return nil
}
func (r *RedisRepository) Get(key string) (map[string]string, error) {
	event := scope + "Get|"

	contextTimeOut, _ := strconv.Atoi(os.Getenv("REDIS_CONTEXT_TIMEOUT"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(contextTimeOut)*time.Second)
	defer cancel()

	dataRedis, err := r.redis.HGetAll(ctx, key).Result()
	if err != nil {
		data := &helper.HTTPLogData{
			Level:   "error",
			Service: event + "Redis HGetAll",
			Status:  http.StatusInternalServerError,
			Message: "failed HGetAll redis",
			Err:     err,
		}
		helper.HTTPLog(data)
		return dataRedis, err
	}

	return dataRedis, nil
}

func (r *RedisRepository) GetString(key string) (string, error) {
	event := scope + "GetString|"

	contextTimeOut, _ := strconv.Atoi(os.Getenv("REDIS_CONTEXT_TIMEOUT"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(contextTimeOut)*time.Second)
	defer cancel()
	dataRedis, err := r.redis.Get(ctx, key).Result()
	if err != nil {
		data := &helper.HTTPLogData{
			Level:   "error",
			Service: event + "Redis Get",
			Status:  http.StatusInternalServerError,
			Message: "failed Get redis",
			Err:     err,
		}
		helper.HTTPLog(data)
		return dataRedis, err
	}

	return dataRedis, nil
}

func (r *RedisRepository) GetInt(key string) (int64, error) {
	event := scope + "GetInt|"

	contextTimeOut, _ := strconv.Atoi(os.Getenv("REDIS_CONTEXT_TIMEOUT"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(contextTimeOut)*time.Second)
	defer cancel()

	dataRedis, err := r.redis.Get(ctx, key).Int64()
	if err != nil {

		data := &helper.HTTPLogData{
			Level:   "error",
			Service: event + "Redis Get",
			Status:  http.StatusInternalServerError,
			Message: "failed Get redis",
			Err:     err,
		}
		helper.HTTPLog(data)
		return dataRedis, err
	}

	return dataRedis, nil
}
func (r *RedisRepository) GetKeys(prefixKey string) ([]string, error) {
	event := scope + "GetKeys|"

	contextTimeOut, _ := strconv.Atoi(os.Getenv("REDIS_CONTEXT_TIMEOUT"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(contextTimeOut)*time.Second)
	defer cancel()
	var cursor uint64
	var keys []string
	intCheck := 0
	for {
		var err error
		keys, cursor, err = r.redis.Scan(ctx, cursor, prefixKey+"*", 0).Result()
		if err != nil {
			data := &helper.HTTPLogData{
				Level:   "error",
				Service: event + "Redis Scan",
				Status:  http.StatusInternalServerError,
				Message: "failed Scan Keys redis",
				Err:     err,
			}
			helper.HTTPLog(data)
			return keys, err
		}
		// fmt.Println(keys)
		fmt.Println(intCheck, prefixKey, cursor)

		for _, key := range keys {
			fmt.Println("key", key)
		}

		if cursor == 0 { // no more keys
			break
		}
		intCheck++
	}
	return keys, nil
}
func (r *RedisRepository) Increment(key string, max int) error {
	event := scope + "Increment|"

	// Transactional function.
	const maxRetries = 1000

	contextTimeOut, _ := strconv.Atoi(os.Getenv("REDIS_CONTEXT_TIMEOUT"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(contextTimeOut)*time.Second)
	defer cancel()
	txf := func(tx *redis.Tx) error {
		// Get the current value or zero.
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			data := &helper.HTTPLogData{
				Level:   "error",
				Service: event + "Redis Get Int",
				Status:  http.StatusInternalServerError,
				Message: "failed Get redis",
				Err:     err,
			}
			helper.HTTPLog(data)
			return exceptions.ErrSystem
		}

		// Actual operation (local in optimistic lock).
		n++
		if n > max {
			data := &helper.HTTPLogData{
				Level:   "error",
				Service: event + "Redis Validate",
				Status:  http.StatusInternalServerError,
				Message: "failed redis validate",
				Err:     exceptions.ErrUnauthorized,
			}
			helper.HTTPLog(data)
			return exceptions.ErrUnauthorized
		}

		// Operation is commited only if the watched keys remain unchanged.
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			pipe.Set(ctx, key, n, 0)
			return nil
		})
		data := &helper.HTTPLogData{
			Level:   "error",
			Service: event + "Redis tx Pipelined",
			Status:  http.StatusInternalServerError,
			Message: "failed tx Pipelined Set",
			Err:     err,
		}
		helper.HTTPLog(data)

		return err
	}
	// time.Sleep(time.Duration(5) * time.Second)
	// Retry if the key has been changed.
	for i := 0; i < maxRetries; i++ {
		fmt.Println(i)

		err := r.redis.Watch(ctx, txf, key)
		if err == nil {
			// Success.
			return nil
		}
		if err == redis.TxFailedErr {
			// Optimistic lock lost. Retry.
			continue
		}
		// Return any other error.
		data := &helper.HTTPLogData{
			Level:   "error",
			Service: event + "Redis Watch",
			Status:  http.StatusInternalServerError,
			Message: "failed Watch redis",
			Err:     err,
		}
		helper.HTTPLog(data)
		return err
	}
	err := errors.New("increment reached maximum number of retries")
	data := &helper.HTTPLogData{
		Level:   "error",
		Service: event + "Redis Watch",
		Status:  http.StatusInternalServerError,
		Message: "failed Watch redis",
		Err:     err,
	}
	helper.HTTPLog(data)
	return err
}

func (r *RedisRepository) IncrementWithObject(key string, max int) (*domain_reset_password.RedisValue, error) {
	event := scope + "IncrementWithObject|"

	// Transactional function.
	const maxRetries = 1000
	redisValue := &domain_reset_password.RedisValue{}

	contextTimeOut, _ := strconv.Atoi(os.Getenv("REDIS_CONTEXT_TIMEOUT"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(contextTimeOut)*time.Second)
	defer cancel()
	txf := func(tx *redis.Tx) error {
		// Get the current value or zero.
		dataRedisString, err := tx.Get(ctx, key).Result()
		if err != nil && err != redis.Nil {
			data := &helper.HTTPLogData{
				Level:   "error",
				Service: event + "Redis Get String",
				Status:  http.StatusInternalServerError,
				Message: "failed Get redis",
				Err:     err,
			}
			helper.HTTPLog(data)
			return exceptions.ErrSystem
		}

		err = json.Unmarshal([]byte(dataRedisString), &redisValue)
		if err != nil {
			data := &helper.HTTPLogData{
				Level:   "error",
				Service: event + "JSON_UNMARSHAL",
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("ERROR : %v | UUID %v", err.Error(), redisValue.UUID),
				Err:     err,
			}
			helper.HTTPLog(data)
			return exceptions.ErrSystem
		}
		redisValue.Increment++
		// Actual operation (local in optimistic lock).
		if redisValue.Increment > max {
			data := &helper.HTTPLogData{
				Level:   "error",
				Service: event + "Redis Validate",
				Status:  http.StatusInternalServerError,
				Message: "failed redis validate",
				Err:     exceptions.ErrUnauthorized,
			}
			helper.HTTPLog(data)
			return exceptions.ErrUnauthorized
		}

		// Operation is commited only if the watched keys remain unchanged.
		_, err = tx.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
			stringValue, _ := json.Marshal(redisValue)
			pipe.Set(ctx, key, stringValue, 0)
			return nil
		})
		data := &helper.HTTPLogData{
			Level:   "error",
			Service: event + "Redis tx Pipelined",
			Status:  http.StatusInternalServerError,
			Message: "failed tx Pipelined Set",
			Err:     err,
		}
		helper.HTTPLog(data)

		return err
	}
	// time.Sleep(time.Duration(5) * time.Second)
	// Retry if the key has been changed.
	for i := 0; i < maxRetries; i++ {
		fmt.Println(i)

		err := r.redis.Watch(ctx, txf, key)
		if err == nil {
			// Success.
			return redisValue, nil
		}
		if err == redis.TxFailedErr {
			// Optimistic lock lost. Retry.
			continue
		}
		// Return any other error.
		data := &helper.HTTPLogData{
			Level:   "error",
			Service: event + "Redis Watch",
			Status:  http.StatusInternalServerError,
			Message: "failed Watch redis",
			Err:     err,
		}
		helper.HTTPLog(data)
		return nil, err
	}
	err := errors.New("increment reached maximum number of retries")
	data := &helper.HTTPLogData{
		Level:   "error",
		Service: event + "Redis Watch",
		Status:  http.StatusInternalServerError,
		Message: "failed Watch redis",
		Err:     err,
	}
	helper.HTTPLog(data)
	return nil, err
}
func (r *RedisRepository) Del(key string) (int64, error) {
	event := scope + "Del|"
	contextTimeOut, _ := strconv.Atoi(os.Getenv("REDIS_CONTEXT_TIMEOUT"))
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(contextTimeOut)*time.Second)
	defer cancel()

	countDelete, err := r.redis.Del(ctx, key).Result()
	if err != nil {

		data := &helper.HTTPLogData{
			Level:   "error",
			Service: event + "Redis Del",
			Status:  http.StatusInternalServerError,
			Message: "failed Delete redis data",
			Err:     err,
		}
		helper.HTTPLog(data)
		return countDelete, exceptions.ErrSystem
	}

	return countDelete, nil
}
