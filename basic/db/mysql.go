package db

import (
	"database/sql"
	"time"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/superryanguo/lightning/basic/config"
)

func initMysql() {
	var err error

	// 创建连接
	mysqlDB, err = sql.Open("mysql", config.GetMysqlConfig().GetURL())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	// 最大连接数
	mysqlDB.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	// 最大闲置数
	mysqlDB.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	//连接数据库闲置断线的问题
	mysqlDB.SetConnMaxLifetime(time.Second * time.Duration(config.GetMysqlConfig().GetConnMaxLifetime()))
	// 激活链接
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
}
