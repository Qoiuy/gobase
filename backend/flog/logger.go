package flog

import (
	"log"
)

type LoggerConfig struct {
	LoggerAddress   string
	LoggerNamespace string
	LoggerPassword  string
	LoggerPrefix    string
}

func InitLogger(cfg LoggerConfig, debug bool) {
	prefix := "go-base-daemon"
	// 要么都填，要么都不填
	if !((len(cfg.LoggerAddress) == 0 && len(cfg.LoggerNamespace) == 0) ||
		(len(cfg.LoggerAddress) > 0 && len(cfg.LoggerNamespace) > 0)) {
		log.Fatalf("wrong logger parameter: missing LoggerAddress(%s) or LoggerNamespace(%s)", cfg.LoggerAddress, cfg.LoggerNamespace)
	}

	if cfg.LoggerPrefix != "" {
		prefix = cfg.LoggerPrefix
	}
	SetPrefix(prefix)
	Infof("begin_serving %s ...", prefix)

	// debug flag
	if debug {
		Info("Set debug model")
		SetLevel(DEBUG)
	}
}
