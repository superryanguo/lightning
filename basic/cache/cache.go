package cache

import (
	"fmt"
	"sync"
	"time"

	"github.com/go-redis/redis"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/superryanguo/lightning/basic/rediser"
)

var (
	rc          *redis.Client
	m           sync.RWMutex
	ExpiredDate = 3600 * 24 * 30 * time.Second
)

func SaveToCache(key string, val []byte) (err error) {
	if rc == nil {
		log.Debug("redis.client un-init")
	}
	log.Debug("SaveCache key=", key, " val=", string(val))

	if err = rc.Set(key, val, ExpiredDate).Err(); err != nil {
		return fmt.Errorf("[saveToCache] err:" + err.Error())
	}
	return
}

func DelFromCache(key string) (err error) {
	if rc == nil {
		log.Debug("redis.client un-init")
	}
	if err = rc.Del(key).Err(); err != nil {
		return fmt.Errorf("[delFromCache] err:" + err.Error())
	}
	return
}

func GetFromCache(key string) (string, error) {
	if rc == nil {
		log.Debug("redis.client un-init")
	}
	val, err := rc.Get(key).Result()
	if err != nil {
		return "", fmt.Errorf("[getFromCache]find no %s", err)
	}

	return val, nil
}

func Init() {
	m.Lock()
	defer m.Unlock()

	rc = rediser.GetRedis()
}
