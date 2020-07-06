package models

import (
	"fmt"

	log "github.com/micro/go-micro/v2/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/superryanguo/lightning/basic/config"
)

type User struct {
	ggorm.Model
	Uid           string        `gorm:"primary_key;size:256" json:"user_id"`
	Name          string        `gorm:"size:50"  json:"name"`
	Password_hash string        `gorm:"size:256" json:"password"`
	Email         string        `gorm:"size:50;unique"  json:"email"`
	Real_name     string        `gorm:"size:32" json:"real_name"`
	Id_card       string        `gorm:"size:20" json:"id_card"`
	Avatar_url    string        `gorm:"size:256" json:"avatar_url"`
	Houses        []*House      `gorm:"many2many:user_houses;" json:"houses"`
	Orders        []*OrderHouse `gorm:"many2many:user_orders;" json:"orders"`
}

func Init() {
	log.Info("Initing the models to create tables")
	c := config.GetMysqlConfig()
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetMysqlConfig().GetName(),
		config.GetMysqlConfig().GetPsw(),
		config.GetMysqlConfig().GetURL(),
		config.GetMysqlConfig().GetDbName())
	db, err := gorm.Open("mysql", config)
	//db, err := gorm.Open("mysql", "root:yourpassword@tcp(127.0.0.1:3306)/testorm?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})

}
