package repository

import (
	"barafiri-platform-service/internal/core/helper"
	"barafiri-platform-service/internal/ports"
	"github.com/gomodule/redigo/redis"
	"time"
)

type RedisRepositories struct {
	Configuration ports.ConfigurationRepository
}

func ConnectToRedis(redisHost string, redisPort string) (RedisRepositories, error) {
	helper.LogEvent("INFO", "Establishing redis connection with given credentials...")
	collection := &redis.Pool{
		MaxIdle:     5,
		IdleTimeout: 60 * time.Second,
		Dial:        func() (redis.Conn, error) { return redis.Dial("tcp", redisHost+":"+redisPort) },
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, "Unable to connect to Redis"))
				return nil
			}
			_, err := c.Do("PING")
			helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, ""))
			return err
		},
	}
	repo := RedisRepositories{
		Configuration: NewConfiguration(collection),
	}
	return repo, nil
}
