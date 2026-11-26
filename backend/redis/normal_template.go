package redis

import (
	redigo "github.com/gomodule/redigo/redis"
)

// NormalTemplate NormalTemplate
type NormalTemplate struct {
	pool redigo.Pool
}

// NewNormalTemplate NewNormalTemplate
func NewNormalTemplate(url string, config *Config) *NormalTemplate {
	s := &NormalTemplate{}
	s.pool = newPool(url, config.Password, config.RedisMaxIdleConnection, config.RedisDatabase,
		config.RedisIdleTimeout, config.RedisConnectTimeout, config.RedisReadTimeout, config.RedisWriteTimeout)
	return s
}

func (s *NormalTemplate) getPool() *redigo.Pool {
	return &s.pool
}

// mget mget
func (s *NormalTemplate) mget(keys ...string) ([]string, error) {
	pool := s.getPool()

	args := []interface{}{}
	for _, key := range keys {
		args = append(args, key)
	}

	rv, err := do(pool, "MGET", args...)
	iRet, err := redigo.Values(rv, err)
	if err != nil {
		return nil, err
	}
	strBuf := make([]string, 0, len(iRet))
	for _, value := range iRet {
		if value != nil {
			strBuf = append(strBuf, string(Uint8s2Bytes(value.([]uint8))))
		} else {
			strBuf = append(strBuf, "")
		}

	}
	return strBuf, nil
}

// Uint8s2Bytes uint8 -> byte
func Uint8s2Bytes(data []uint8) []byte {
	buf := make([]byte, 0, len(data))
	for _, v := range data {
		buf = append(buf, byte(v))
	}
	return buf
}
