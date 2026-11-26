package redis

import "time"

// NewDefaultServer 提供便捷的创建redis服务的接口
func NewDefaultServer(address, password string, redisDB int) *RedisServer {

	maxIdleConnection := 50
	idleTimeout := time.Duration(180) * time.Second
	timeout := time.Duration(10) * time.Second

	cfg := NewConfig(maxIdleConnection, idleTimeout, timeout, timeout, timeout, redisDB, password)

	normalServerTemplate := NewNormalTemplate(address, cfg)
	normalRedis := NewRedisServer()
	normalRedis.SetRedisTemplate(normalServerTemplate)

	return normalRedis
}

// Check 启动时检查是否已连接
func Check(server *RedisServer) error {
	if _, err := server.Ping(); err != nil {
		return err
	}
	return nil
}
