package models

import (
	"fmt"
	"time"

	log "github.com/micro/go-micro/v2/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/superryanguo/lightning/basic/config"
)

type User struct {
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

type Area struct {
	Id     int      `json:"aid"`
	Name   string   `gorm:"size:32" json:"aname"`
	Houses []*House `gorm:"many2many:area_houses" json:"houses"`
}

type Facility struct {
	Id     int      `json:"fid"`
	Name   string   `gorm:"size:32"`
	Houses []*House `gorm:"many2many:facility_houses"`
}

type HouseImage struct {
	Id    int    `json:"house_image_id"`
	Url   string `gorm:"size:256" json:"url"`
	House *House `gorm:"ForeignKey:Id" json:"house_id"`
}
type House struct {
	Id              int           `json:"house_id"`
	User            *User         `gorm:"rel(fk)" json:"user_id"`
	Area            *Area         `gorm:"rel(fk)" json:"area_id"`
	Title           string        `gorm:"size:64" json:"title"`
	Price           int           `gorm:"default:0" json:"price"`
	Address         string        `gorm:"size:512;default:''" json:"address"`
	Room_count      int           `gorm:"default:1" json:"room_count"`
	Acreage         int           `gorm:"default:0" json:"acreage"`
	Unit            string        `gorm:"size:32;default:''" json:"unit"`
	Capacity        int           `gorm:"default:1" json:"capacity"`
	Beds            string        `gorm:"size:64;default:''" json:"beds"`
	Deposit         int           `gorm:"default(0)" json:"deposit"`
	Min_days        int           `gorm:"default(1)" json:"min_days"`
	Max_days        int           `gorm:"default(0)" json:"max_days"`
	Order_count     int           `gorm:"default(0)" json:"order_count"`
	Index_image_url string        `gorm:"size:256;default:''" json:"index_image_url"`
	Facilities      []*Facility   `gorm:"many2many:house_facilities;" json:"facilities"`
	Images          []*HouseImage `gorm:"many2many:house_images" json:"img_urls"`
	Orders          []*OrderHouse `gorm:"many2many:house_orders" json:"orders"`
	Ctime           time.Time     `gorm:"auto_now_add;type(datetime)" json:"ctime"`
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
	db.AutoMigrate(&User{}, &House{})

}
