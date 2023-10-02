package repository

import (
	"barafiri-platform-service/internal/core/helper"
	"barafiri-platform-service/internal/ports"
	"github.com/gomodule/redigo/redis"
)

type ConfigurationInfra struct {
	Collection *redis.Pool
}

func NewConfiguration(Collection *redis.Pool) *ConfigurationInfra {
	return &ConfigurationInfra{Collection}
}

var _ ports.ConfigurationRepository = &ConfigurationInfra{}

func (r *ConfigurationInfra) SetConfiguration(key string, configuration interface{}) (interface{}, error) {
	helper.LogEvent("INFO", "Persisting platform configuration on redis hash with key: "+key)

	conn := r.Collection.Get()
	reply, err := conn.Do("HSET", redis.Args{}.Add(key).AddFlat(configuration)...)
	if err != nil {
		helper.LogEvent("ERROR", helper.ErrorMessage(helper.RedisSetupError, err.Error()))
		return nil, err
	}
	helper.LogEvent("INFO", "Persisting platform configurations on redis hash completed successfully...")

	return reply, nil

}
