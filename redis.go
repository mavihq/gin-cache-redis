package cache

import (
	"time"

	"github.com/go-redis/redis"
)

type Redis struct {
	Client *redis.Client
}

func NewRedis(options *redis.Options) *Redis {
	var _options *redis.Options
	if options != nil {
		_options = options
	} else {
		_options = &redis.Options{
			Addr:     "localhost:6379",
			Password: "", // no password set
			DB:       0,  // use default DB
		}
	}
	client := redis.NewClient(_options)
	return &Redis{client}
}

func (self *Redis) Get(key string) ([]byte, error) {
	return self.Client.Get(key).Bytes()
}

func (self *Redis) Set(key string, value string, expire time.Duration) error {
	return self.Client.Set(key, value, expire).Err()
}

func (self *Redis) Remove(key string) error {
	return self.Client.Del(key).Err()
}

func (self *Redis) Update(key string, value string, expire time.Duration) error {
	return self.Set(key, value, expire)
}
