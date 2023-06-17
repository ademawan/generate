package redis_revosotory

import domain_reset_password "merchant_panel_force_logout/domain/reset_password"

type Redis interface {
	Set(key string, value interface{}, ttlDuration string) error
	SetString(key string, value interface{}, duration int) error
	Get(key string) (map[string]string, error)
	GetString(key string) (string, error)
	GetInt(key string) (int64, error)
	GetKeys(prefixKey string) ([]string, error)
	Increment(key string, max int) error
	IncrementWithObject(key string, max int) (*domain_reset_password.RedisValue, error)
	Del(key string) (int64, error)
}
