package models

import (
	"fmt"
	"time"

	log "github.com/micro/go-micro/v2/logger"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/superryanguo/lightning/basic/config"
	"github.com/superryanguo/lightning/utils"
)

const (
	ORDER_STATUS_WAIT_ACCEPT  = "WAIT_ACCEPT"
	ORDER_STATUS_WAIT_PAYMENT = "WAIT_PAYMENT"
	ORDER_STATUS_PAID         = "PAID"
	ORDER_STATUS_WAIT_COMMENT = "WAIT_COMMENT"
	ORDER_STATUS_COMPLETE     = "COMPLETE"
	ORDER_STATUS_CANCELED     = "CONCELED"
	ORDER_STATUS_REJECTED     = "REJECTED"
)

var HOME_PAGE_MAX_HOUSES int = 5

var HOUSE_LIST_PAGE_CAPACITY int = 2

type User struct {
	//gorm.Model, we only need the p-key, so it's necessary to add the gorm.model
	Uid           string `gorm:"primary_key;size:256" json:"user_id"`
	Name          string `gorm:"size:50"  json:"name"`
	Password_hash string `gorm:"size:256" json:"password"`
	Email         string `gorm:"size:50;unique"  json:"email"`
	Real_name     string `gorm:"size:32" json:"real_name"`
	Id_card       string `gorm:"size:20" json:"id_card"`
	Avatar_url    string `gorm:"size:256" json:"avatar_url"`
	//one to many, `gorm:"foreignKey:UserID"`
	Houses []*House      `json:"houses"`
	Orders []*OrderHouse `json:"orders"`
}

type House struct {
	//gorm.Model
	ID              int           `json:"house_id"`
	UserID          uint          `json:"user_id"`
	AreaID          uint          `json:"area_id"`
	Title           string        `gorm:"size:64" json:"title"`
	Price           int           `gorm:"default:0" json:"price"`
	Address         string        `gorm:"size:512;default:''" json:"address"`
	Room_count      int           `gorm:"default:1" json:"room_count"`
	Acreage         int           `gorm:"default:0" json:"acreage"`
	Unit            string        `gorm:"size:32;default:''" json:"unit"`
	Capacity        int           `gorm:"default:1" json:"capacity"`
	Beds            string        `gorm:"size:64;default:''" json:"beds"`
	Deposit         int           `gorm:"default:0" json:"deposit"`
	Min_days        int           `gorm:"default:1" json:"min_days"`
	Max_days        int           `gorm:"default:0" json:"max_days"`
	Order_count     int           `gorm:"default:0" json:"order_count"`
	Index_image_url string        `gorm:"size:256;default:''" json:"index_image_url"`
	Facilities      []*Facility   `gorm:"many2many:house_facilities;" json:"facilities"`
	Images          []*HouseImage `json:"img_urls"`
	Orders          []*OrderHouse `json:"orders"`
	Ctime           time.Time     `gorm:"autoCreateTime" json:"ctime"`
}

type OrderHouse struct {
	ID          int       `json:"order_id"`
	UserID      uint      `json:"user_id"`
	HouseID     uint      `json:"house_id"`
	Begin_date  time.Time `gorm:"type:time"`
	End_date    time.Time `gorm:"type:time"`
	Days        int
	House_price int
	Amount      int
	Status      string    `gorm:"default:'WAIT_ACCEPT'"`
	Comment     string    `gorm:"size:512"`
	Ctime       time.Time `gorm:"autoCreateTime" json:"ctime"`
	Credit      bool
}

type HouseImage struct {
	ID      int    `json:"house_image_id"`
	Url     string `gorm:"size:256" json:"url"`
	HouseID uint   `json:"house_id"`
}

type Area struct {
	//gorm.Model
	ID     int      `json:"area_id"`
	Name   string   `gorm:"size:32" json:"area_name"`
	Houses []*House `json:"houses"`
}

type Facility struct {
	ID     int      `json:"fid"`
	Name   string   `gorm:"size:32"`
	Houses []*House `gorm:"many2many:house_facilities"`
}

func (this *House) To_house_info() interface{} {
	house_info := map[string]interface{}{
		"house_id":    this.ID,
		"title":       this.Title,
		"price":       this.Price,
		"area_name":   "this.Area.Name", //this.Area.Name,
		"img_url":     utils.AddDomain2Url(this.Index_image_url),
		"room_count":  this.Room_count,
		"order_count": this.Order_count,
		"address":     this.Address,
		"user_avatar": "NA url", //utils.AddDomain2Url(this.User.Avatar_url),
		"ctime":       this.Ctime.Format("2006-01-02 15:04:05"),
	}

	return house_info
}

func (this *OrderHouse) To_order_info() interface{} {
	order_info := map[string]interface{}{
		"order_id":   this.ID,
		"title":      "NA house title", //this.House.Title,
		"img_url":    "NA url",         //utils.AddDomain2Url(this.House.Index_image_url),
		"start_date": this.Begin_date.Format("2006-01-02 15:04:05"),
		"end_date":   this.End_date.Format("2006-01-02 15:04:05"),
		"ctime":      this.Ctime.Format("2006-01-02 15:04:05"),
		"days":       this.Days,
		"amount":     this.Amount,
		"status":     this.Status,
		"comment":    this.Comment,
		"credit":     this.Credit,
	}

	return order_info
}

func (this *House) To_one_house_desc() interface{} {
	house_desc := map[string]interface{}{
		"hid":         this.ID,
		"user_id":     "userid NA",       //this.User.Uid,
		"user_name":   "user name NA",    //this.User.Name,
		"user_avatar": "user avastar NA", //utils.AddDomain2Url(this.User.Avatar_url),
		"title":       this.Title,
		"price":       this.Price,
		"address":     this.Address,
		"room_count":  this.Room_count,
		"acreage":     this.Acreage,
		"unit":        this.Unit,
		"capacity":    this.Capacity,
		"beds":        this.Beds,
		"deposit":     this.Deposit,
		"min_days":    this.Min_days,
		"max_days":    this.Max_days,
	}

	//img_urls := []string{}
	//for _, img_url := range this.Images {
	//img_urls = append(img_urls, utils.AddDomain2Url(img_url.Url))
	//}
	//house_desc["img_urls"] = img_urls

	//facilities := []int{}
	//for _, facility := range this.Facilities {
	//facilities = append(facilities, facility.Id)
	//}
	//house_desc["facilities"] = facilities

	//comments := []interface{}{}
	//orders := []OrderHouse{}
	////o := orm.NewOrm()
	//order_num, err := o.QueryTable("order_house").Filter("house__id", this.Id).Filter("status", ORDER_STATUS_COMPLETE).OrderBy("-ctime").Limit(10).All(&orders)
	//if err != nil {
	//log.Error("select orders comments error, err =", err, "house id = ", this.Id)
	//}
	//for i := 0; i < int(order_num); i++ {
	//o.LoadRelated(&orders[i], "User")
	//var username string
	//if orders[i].User.Name == "" {
	//username = "匿名用户"
	//} else {
	//username = orders[i].User.Name
	//}

	//comment := map[string]string{
	//"comment":   orders[i].Comment,
	//"user_name": username,
	//"ctime":     orders[i].Ctime.Format("2006-01-02 15:04:05"),
	//}
	//comments = append(comments, comment)
	//}
	//house_desc["comments"] = comments

	return house_desc
}

func Init() {
	log.Info("Initing the models to create tables")
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetMysqlConfig().GetName(),
		config.GetMysqlConfig().GetPsw(),
		config.GetMysqlConfig().GetURL(),
		config.GetMysqlConfig().GetDbName())
	log.Debug("connect config=", config)
	db, err := gorm.Open("mysql", config)
	//db, err := gorm.Open("mysql", "root:yourpassword@tcp(127.0.0.1:3306)/testorm?charset=utf8mb4&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{}, &House{}, &OrderHouse{}, &Area{}, &Facility{}, &HouseImage{})

	log.Info("Database tables init done")
}
