package models

import (
	"time"

	log "github.com/micro/go-micro/v2/logger"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/superryanguo/lightning/utils"
)

type User struct {
	Uid           string        `orm:"pk;size(256)" json:"user_id"`
	Name          string        `orm:"size(50)"  json:"name"`
	Password_hash string        `orm:"size(256)" json:"password"`
	Email         string        `orm:"size(50);unique"  json:"email"`
	Real_name     string        `orm:"size(32)" json:"real_name"`
	Id_card       string        `orm:"size(20)" json:"id_card"`
	Avatar_url    string        `orm:"size(256)" json:"avatar_url"`
	Houses        []*House      `orm:"reverse(many)" json:"houses"`
	Orders        []*OrderHouse `orm:"reverse(many)" json:"orders"`
}

type House struct {
	Id              int           `json:"house_id"`
	User            *User         `orm:"rel(fk)" json:"user_id"`
	Area            *Area         `orm:"rel(fk)" json:"area_id"`
	Title           string        `orm:"size(64)" json:"title"`
	Price           int           `orm:"default(0)" json:"price"`
	Address         string        `orm:"size(512)" orm:"default("")" json:"address"`
	Room_count      int           `orm:"default(1)" json:"room_count"`
	Acreage         int           `orm:"default(0)" json:"acreage"`
	Unit            string        `orm:"size(32)" orm:"default("")" json:"unit"`
	Capacity        int           `orm:"default(1)" json:"capacity"`
	Beds            string        `orm:"size(64)" orm:"default("")" json:"beds"`
	Deposit         int           `orm:"default(0)" json:"deposit"`
	Min_days        int           `orm:"default(1)" json:"min_days"`
	Max_days        int           `orm:"default(0)" json:"max_days"`
	Order_count     int           `orm:"default(0)" json:"order_count"`
	Index_image_url string        `orm:"size(256)" orm:"default("")" json:"index_image_url"`
	Facilities      []*Facility   `orm:"reverse(many)" json:"facilities"`
	Images          []*HouseImage `orm:"reverse(many)" json:"img_urls"`
	Orders          []*OrderHouse `orm:"reverse(many)" json:"orders"`
	Ctime           time.Time     `orm:"auto_now_add;type(datetime)" json:"ctime"`
}

var HOME_PAGE_MAX_HOUSES int = 5

var HOUSE_LIST_PAGE_CAPACITY int = 2

func (this *House) To_house_info() interface{} {
	house_info := map[string]interface{}{
		"house_id":    this.Id,
		"title":       this.Title,
		"price":       this.Price,
		"area_name":   this.Area.Name,
		"img_url":     utils.AddDomain2Url(this.Index_image_url),
		"room_count":  this.Room_count,
		"order_count": this.Order_count,
		"address":     this.Address,
		"user_avatar": utils.AddDomain2Url(this.User.Avatar_url),
		"ctime":       this.Ctime.Format("2006-01-02 15:04:05"),
	}

	return house_info
}

func (this *House) To_one_house_desc() interface{} {
	house_desc := map[string]interface{}{
		"hid":         this.Id,
		"user_id":     this.User.Uid,
		"user_name":   this.User.Name,
		"user_avatar": utils.AddDomain2Url(this.User.Avatar_url),
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

	img_urls := []string{}
	for _, img_url := range this.Images {
		img_urls = append(img_urls, utils.AddDomain2Url(img_url.Url))
	}
	house_desc["img_urls"] = img_urls

	facilities := []int{}
	for _, facility := range this.Facilities {
		facilities = append(facilities, facility.Id)
	}
	house_desc["facilities"] = facilities

	//评论信息

	comments := []interface{}{}
	orders := []OrderHouse{}
	o := orm.NewOrm()
	order_num, err := o.QueryTable("order_house").Filter("house__id", this.Id).Filter("status", ORDER_STATUS_COMPLETE).OrderBy("-ctime").Limit(10).All(&orders)
	if err != nil {
		log.Error("select orders comments error, err =", err, "house id = ", this.Id)
	}
	for i := 0; i < int(order_num); i++ {
		o.LoadRelated(&orders[i], "User")
		var username string
		if orders[i].User.Name == "" {
			username = "匿名用户"
		} else {
			username = orders[i].User.Name
		}

		comment := map[string]string{
			"comment":   orders[i].Comment,
			"user_name": username,
			"ctime":     orders[i].Ctime.Format("2006-01-02 15:04:05"),
		}
		comments = append(comments, comment)
	}
	house_desc["comments"] = comments

	return house_desc
}

type Area struct {
	Id     int      `json:"aid"`
	Name   string   `orm:"size(32)" json:"aname"`
	Houses []*House `orm:"reverse(many)" json:"houses"`
}

type Facility struct {
	Id     int      `json:"fid"`
	Name   string   `orm:"size(32)"`
	Houses []*House `orm:"rel(m2m)"`
}

type HouseImage struct {
	Id    int    `json:"house_image_id"`
	Url   string `orm:"size(256)" json:"url"`
	House *House `orm:"rel(fk)" json:"house_id"`
}

const (
	ORDER_STATUS_WAIT_ACCEPT  = "WAIT_ACCEPT"  //待接单
	ORDER_STATUS_WAIT_PAYMENT = "WAIT_PAYMENT" //待支付
	ORDER_STATUS_PAID         = "PAID"         //已支付
	ORDER_STATUS_WAIT_COMMENT = "WAIT_COMMENT" //待评价
	ORDER_STATUS_COMPLETE     = "COMPLETE"     //已完成
	ORDER_STATUS_CANCELED     = "CONCELED"     //已取消
	ORDER_STATUS_REJECTED     = "REJECTED"     //已拒单
)

type OrderHouse struct {
	Id          int       `json:"order_id"`
	User        *User     `orm:"rel(fk)" json:"user_id"`
	House       *House    `orm:"rel(fk)" json:"house_id"`
	Begin_date  time.Time `orm:"type(datetime)"`
	End_date    time.Time `orm:"type(datetime)"`
	Days        int
	House_price int
	Amount      int
	Status      string    `orm:"default(WAIT_ACCEPT)"`
	Comment     string    `orm:"size(512)"`
	Ctime       time.Time `orm:"auto_now;type(datetime)" json:"ctime"`
	Credit      bool
}

func (this *OrderHouse) To_order_info() interface{} {
	order_info := map[string]interface{}{
		"order_id":   this.Id,
		"title":      this.House.Title,
		"img_url":    utils.AddDomain2Url(this.House.Index_image_url),
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

func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)

	//connect database   ( 默认参数 ，mysql数据库 ，"数据库的用户名 ：数据库密码@tcp("+数据库地址+":"+数据库端口+")/库名？格式",默认参数）
	orm.RegisterDataBase("default", "mysql", "root:"+utils.G_mysql_passwd+"@tcp("+utils.G_mysql_addr+":"+utils.G_mysql_port+")/"+utils.G_mysql_dbname+"?charset=utf8", 30)

	//create tables
	orm.RegisterModel(new(User), new(House), new(Area), new(Facility), new(HouseImage), new(OrderHouse))

	orm.RunSyncdb("default", false, true)
}
