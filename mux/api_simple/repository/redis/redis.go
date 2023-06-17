package redis

import (
	"api_simple/helper/exceptions"
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

type RedisRepository struct {
	redis *redis.Client
}

func NewRedisRepository(redis *redis.Client) *RedisRepository {
	return &RedisRepository{redis: redis}
}

func (r *RedisRepository) Set(key string, value interface{}, ttlDuration string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	ttl, err := strconv.Atoi(ttlDuration)
	if err != nil {
		logrus.Errorf("error submit fmc msisdn %s", err)
		return err
	}
	err = r.redis.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()

	if err != nil {
		fmt.Println("err redis", err)
		return err
	}

	return nil
}
func (r *RedisRepository) Get(key string) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	dataRedis, err := r.redis.HGetAll(ctx, key).Result()
	if err != nil {

		log.Println("error get inqury redis", err)
		return dataRedis, err
	}

	return dataRedis, nil
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
func (r *RedisRepository) Increment(key string, max int) error {
	// Transactional function.
	const maxRetries = 1000

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	txf := func(tx *redis.Tx) error {
		// Get the current value or zero.
		n, err := tx.Get(ctx, key).Int()
		if err != nil && err != redis.Nil {
			return exceptions.ErrSystem
		}

		// Actual operation (local in optimistic lock).
		n++
		if n > max {
			return exceptions.ErrUnauthorized
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
		return err
	}

	return errors.New("increment reached maximum number of retries")
}
func (r *RedisRepository) Del(key string) (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	countDelete, err := r.redis.Del(ctx, key).Result()
	if err != nil {

		log.Println("error get inqury redis", err)
		return countDelete, exceptions.ErrSystem
	}

	return countDelete, nil
}
