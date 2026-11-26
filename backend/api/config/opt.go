package config

var Opts = &Options{}

// Options 启动参数
type Options struct {
	RedisConfig              `mapstructure:",squash"`
	LoggerConfig             `mapstructure:",squash"`
	MysqlConfig              `mapstructure:",squash"`
	Debug                    bool   `mapstructure:"debug"`
	BaseURL                  string `mapstructure:"base_url"`
	Port                     int    `mapstructure:"port"`
	AccessControlAllowOrigin string `mapstructure:"access_control_allow_origin"`
	Brokers                  string `mapstructure:"brokers"`
}

// RedisConfig redis配置项
type RedisConfig struct {
	RedisAddress  string `mapstructure:"redis_address"`
	RedisPassword string `mapstructure:"redis_password"`
	RedisDB       int    `mapstructure:"redis_db"`
}

// LoggerConfig 日志服务器配置项
type LoggerConfig struct {
	LoggerAddress   string `mapstructure:"logger_address"`
	LoggerPassword  string `mapstructure:"logger_password"`
	LoggerNamespace string `mapstructure:"logger_namespace"`
	LoggerPrefix    string `mapstructure:"logger_prefix"`
}

// MysqlConfig MySQL配置项
type MysqlConfig struct {
	MysqlAddress    string `mapstructure:"mysql_address"`
	MysqlUser       string `mapstructure:"mysql_user"`
	MysqlPassword   string `mapstructure:"mysql_password"`
	MysqlDatabase   string `mapstructure:"mysql_database"`
	MaxConnections  int    `mapstructure:"max_connections"`
	ConnMaxLifetime int    `mapstructure:"conn_max_life_time"`
	MaxIdleConns    int    `mapstructure:"max_idle_conns"`
	EnableLogging   int    `mapstructure:"enable_logging"`
}
