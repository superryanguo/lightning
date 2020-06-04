package basic

import (
	"github.com/superryanguo/lightning/basic/cache"
	"github.com/superryanguo/lightning/basic/config"
	"github.com/superryanguo/lightning/basic/db"
	"github.com/superryanguo/lightning/basic/redis"
)

func Init() {
	config.Init()
	redis.Init()
	db.Init()
	cache.Init()
}
