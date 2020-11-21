package basic

import (
	"github.com/superryanguo/lightning/basic/cache"
	"github.com/superryanguo/lightning/basic/config"
	"github.com/superryanguo/lightning/basic/db"
	"github.com/superryanguo/lightning/basic/rediser"
)

func Init() {
	config.Init()
	rediser.Init()
	db.Init()
	cache.Init()
}
