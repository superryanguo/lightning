package basic

import (
	"github.com/superryanguo/lightning/basic/config"
	"github.com/superryanguo/lightning/basic/db"
	"github.com/superryanguo/lightning/basic/redis"
)

func Init() {
	config.Init()
	db.Init()
	redis.Init()
}
