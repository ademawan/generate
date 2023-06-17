package redis

type RedisRepository interface {
	Set(key string, value interface{}, ttlDuration string) error
	Get(key string) (map[string]string, error)
	GetInt(key string) (int64, error)
	Increment(key string, max int) error
	Del(key string) (int64, error)
}
