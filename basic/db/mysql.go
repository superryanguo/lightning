package db

import (
	"database/sql"
	"fmt"
	"time"

	log "github.com/micro/go-micro/v2/logger"
	"github.com/superryanguo/lightning/basic/config"
)

func initMysql() {
	var err error

	// 创建连接
	connect := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.GetMysqlConfig().GetName(),
		config.GetMysqlConfig().GetPsw(),
		config.GetMysqlConfig().GetURL(),
		config.GetMysqlConfig().GetDbName())
	log.Debug("connect sql=", connect)
	mysqlDB, err = sql.Open("mysql", connect)
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
	log.Info("Connecting the mysql database, PING...")
	if err = mysqlDB.Ping(); err != nil {
		log.Fatal(err)
	}
	log.Info("PONG... mysql connected successfully...")
}
