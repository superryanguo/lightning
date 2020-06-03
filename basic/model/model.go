package model

//TODO: make the model name to be redis?!

import (
	"fmt"
	"sync"
	"time"

	r "github.com/go-redis/redis"
	"github.com/superryanguo/lightning/basic/redis"
)

var (
	ca          *r.Client
	m           sync.RWMutex
	ExpiredDate = 3600 * 24 * 30 * time.Second
)

func SaveToCache(key string, val []byte) (err error) {

	if err = ca.Set(key, val, ExpiredDate).Err(); err != nil {
		return fmt.Errorf("[saveToCache] 保存到缓存发生错误，err:" + err.Error())
	}
	return
}

func DelFromCache(key string) (err error) {
	if err = ca.Del(key).Err(); err != nil {
		return fmt.Errorf("[delFromCache] 清空缓存发生错误，err:" + err.Error())
	}
	return
}

func GetFromCache(key string) (string, error) {
	val, err := ca.Get(key).Result()
	if err != nil {
		return "", fmt.Errorf("[getFromCache]不存在 %s", err)
	}

	return val, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()

	ca = redis.GetRedis()
}
