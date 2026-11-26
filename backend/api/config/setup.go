package config

import (
	"log"
	"os"

	"github.com/spf13/viper"
)

// defaultConfigFile 默认配置文件，相对于工作目录
const defaultConfigFile = "config.yaml"

// Setup 从 config.yaml 读取启动参数
func Setup() {
	setupFromFile()
}

func setupFromFile() {
	configFile := defaultConfigFile
	if p := os.Getenv("CONFIG_FILE"); p != "" {
		configFile = p
	}

	viper.SetConfigFile(configFile)

	if err := viper.ReadInConfig(); err != nil {
		log.Panicf("读取配置文件 %s 失败: %s（可复制 config.example.yaml 后填写）", configFile, err)
	}

	if err := viper.Unmarshal(Opts); err != nil {
		log.Panicf("解析配置文件 %s 失败: %s", configFile, err)
	}
}
