package storage

import "github.com/go-redis/redis"

// Storage exposes a basic Key/Value Storage Interface
type Storage interface {
	Get(key string) (string, error)
	Set(key string, value interface{}) error
	Exists(key string) (bool, error)
}

// RedisStorage Storage Interface backed by Redis
type RedisStorage struct {
	client *redis.Client
}

// Get from redis
func (r *RedisStorage) Get(key string) (string, error) {
	return r.client.Get(key).Result()
}

// Set in redis
func (r *RedisStorage) Set(key string, value interface{}) error {
	return r.client.Set(key, value, 0).Err()
}

// Exists checks if the specified key is in use
// Returns false when an error occurs.
func (r *RedisStorage) Exists(key string) (bool, error) {
	exists, err := r.client.Exists(key).Result()
	if err != nil {
		return false, err
	}

	return exists == 1, nil
}

// String RedisStorage
func (r RedisStorage) String() string {
	return "RedisStorage"
}

// New returns a new Storage Instance
func New(address string) *RedisStorage {
	var storage = &RedisStorage{}
	storage.client = redis.NewClient(&redis.Options{
		Addr: address,
	})

	return storage
}


