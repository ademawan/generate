package repository_test

import (
	repository "redis-test/repository/redis"
	"testing"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
	redisV9 "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func NewMock() *redis.Client {
	mr, _ := miniredis.Run()
	client := redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	return client
}
func NewMockRedisV9() *redisV9.Client {
	mr, _ := miniredis.Run()
	client := redisV9.NewClient(&redisV9.Options{
		Addr: mr.Addr(),
	})

	return client
}
func TestRedis_HGetAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := NewMock()
		clienRedisV9 := NewMockRedisV9()

		repo := repository.NewAuthRepository(client, clienRedisV9)
		_, err := repo.HGetAll("key")
		assert.NoError(t, err)

	})

}
