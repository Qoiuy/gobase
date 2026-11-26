package redis

import (
	"encoding/json"
	"errors"
	"reflect"
)

var D1CacheExpire int64 = 60 * 60 * 24 // 一天过期
var M5CacheExpire int64 = 60 * 5       // 5分钟 过期

//var M10CacheExpire int64 = 60 * 10       // 10分钟 过期

var M30CacheExpire int64 = 60 * 30 // 30分钟 过期

// Get4C redis Get方法 给 Cache使用  主要用于 缓存层
func (s *RedisServer) Get4C(key string, mysql func() (string, error)) (*string, error) {
	return s.GetCache(key, D1CacheExpire, mysql)
}

// GetCache get data from redis. if failed, then get data from mysql and write it into redis
func (s *RedisServer) GetCache(key string, expireTime int64, mysql func() (string, error)) (*string, error) {
	if expireTime < 0 {
		expireTime = D1CacheExpire
	}

	defaultValue := "-"

	redisJson, err := s.Get(key)
	if err != nil {
		return nil, err
	}
	if redisJson != "" {
		if redisJson == defaultValue {
			return nil, nil
		}
		return &redisJson, nil
	}
	res, err := mysql()
	if err != nil {
		return nil, err
	}
	if res == "" {
		if err = s.SetEx(key, defaultValue, M5CacheExpire); err != nil {
			return nil, err
		}
		return nil, nil
	}

	// REDIS SET
	if err = s.SetEx(key, res, expireTime); err != nil {
		return nil, err
	}

	return &res, nil
}

// GetCache 从redis获取数据，若获取失败，执行原有函数。data必须是指针，且非nil 处理资源不存在的情况，请勿使用
func (s *RedisServer) GetCacheWithOrigin(key string, expireTime int64, getDataOriginFunc func() (interface{}, error), data interface{}) error {
	if expireTime < 0 {
		expireTime = D1CacheExpire
	}

	rv := reflect.ValueOf(data)
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return errors.New("data 必须是指针类型且不是nil")
	}

	defaultValue := "-"

	redisJson, err := s.Get(key)
	if err != nil {
		return err
	}
	if redisJson != "" {
		if redisJson == defaultValue {
			return nil
		}
		return json.Unmarshal([]byte(redisJson), data)
	}
	res, err := getDataOriginFunc()
	if err != nil {
		return err
	}

	// 将数据写入原有data
	resValue := reflect.ValueOf(res)
	if resValue.Kind() == reflect.Ptr {
		rv.Elem().Set(resValue.Elem())
	} else {
		rv.Elem().Set(resValue)
	}

	go func() {
		bytes, _ := json.Marshal(res)
		_ = s.SetEx(key, string(bytes), expireTime)
	}()

	return nil
}
