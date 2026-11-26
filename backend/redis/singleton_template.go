package redis

import (
	"errors"
	"sync"

	redigo "github.com/gomodule/redigo/redis"
)

var once sync.Once
var server *SingletonTemplate

// GetInstance 单例redis服务
func GetInstance(redisIps []string, config *Config) *SingletonTemplate {
	once.Do(func() {
		//make(map[string]redis.Pool), config
		server = &SingletonTemplate{
			pools:  make(map[string]redigo.Pool),
			config: config,
		}
		server.Reset(redisIps)
	})

	return server
}

// SingletonTemplate  Singleton Redis Template
type SingletonTemplate struct {
	pools  map[string]redigo.Pool
	config *Config
}

// Reset Reset
func (s *SingletonTemplate) Reset(redisIps []string) {

	for _, value := range redisIps {
		if _, ok := s.pools[value]; ok {
			continue
		} else {
			pool := newPool(value, s.config.Password, s.config.RedisMaxIdleConnection, 0,
				s.config.RedisIdleTimeout, s.config.RedisConnectTimeout, s.config.RedisReadTimeout, s.config.RedisWriteTimeout)
			s.pools[value] = pool
		}
	}

	for k, p := range s.pools {
		exist := false
		for _, value := range redisIps {
			if k == value {
				exist = true
			}
		}
		if !exist {
			p.Close()
			delete(s.pools, k)
		}
	}

}

func (s *SingletonTemplate) getPool() *redigo.Pool {

	for _, value := range s.pools {
		return &value
	}
	return nil
}

// mget mget
func (s *SingletonTemplate) mget(keys ...string) ([]string, error) {

	return nil, errors.New("not support")
}
