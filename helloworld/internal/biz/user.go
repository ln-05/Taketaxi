package biz

import (
	"github.com/google/wire"
	"time"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase)

type User struct {
	Id       int32     `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Mobile   string    `gorm:"column:mobile;type:char(11);default:NULL;" json:"mobile"`
	UserName string    `gorm:"column:user_name;type:varchar(20);default:NULL;" json:"user_name"`
	NickName string    `gorm:"column:nick_name;type:varchar(20);default:NULL;" json:"nick_name"`
	IdCard   string    `gorm:"column:id_card;type:varchar(18);default:NULL;" json:"id_card"`
	Sex      string    `gorm:"column:sex;type:varchar(5);default:NULL;" json:"sex"`
	Mileage  string    `gorm:"column:mileage;type:varchar(10);default:NULL;" json:"mileage"`
	CreateAt time.Time `gorm:"column:create_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"create_at"`
	UpdateAt time.Time `gorm:"column:update_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"update_at"`
	DeleteAt time.Time `gorm:"column:delete_at;type:datetime(3);default:NULL;" json:"delete_at"`
}

type Auth struct {
	Id          int32  `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	AppName     string `gorm:"column:app_name;type:varchar(20);default:NULL;" json:"app_name"`
	AppUnionid  string `gorm:"column:app_unionid;type:varchar(60);default:NULL;" json:"app_unionid"`
	AccessToken string `gorm:"column:access_token;type:varchar(60);default:NULL;" json:"access_token"`
	AppRemark   string `gorm:"column:app_remark;type:varchar(255);default:NULL;" json:"app_remark"`
}
type UserAuth struct {
	Id     int32 `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	UserId int32 `gorm:"column:user_id;type:int;default:NULL;" json:"user_id"`
	AuthId int32 `gorm:"column:auth_id;type:int;default:NULL;" json:"auth_id"`
}
type DriverAuth struct {
	Id           int32     `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	DriverId     int32     `gorm:"column:driver_id;type:int;default:NULL;" json:"driver_id"`
	IdCord       string    `gorm:"column:id_cord;type:varchar(20);default:NULL;" json:"id_cord"`
	VerifyStatus int8      `gorm:"column:verify_status;type:tinyint(1);comment:认证状态(0:未认证 1:已认证 2:拒绝)	;default:NULL;" json:"verify_status"` // 认证状态(0:未认证 1:已认证 2:拒绝)
	VerifyTime   time.Time `gorm:"column:verify_time;type:timestamp;default:NULL;" json:"verify_time"`
}
type Driver struct {
	Id           int32     `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	UserId       int32     `gorm:"column:user_id;type:int;default:NULL;" json:"user_id"`
	DriveStatus  int8      `gorm:"column:drive_status;type:tinyint(1);comment:"已激活" "未激活" "封号";default:NULL;" json:"drive_status"` // "已激活" "未激活" "封号"
	OnlineStatus int8      `gorm:"column:online_status;type:tinyint(1);comment:"在线" "离线";default:NULL;" json:"online_status"`      // "在线" "离线"
	WokeCode     string    `gorm:"column:woke_code;type:varchar(30);default:NULL;" json:"woke_code"`
	Level        int32     `gorm:"column:level;type:int;comment:等级;default:NULL;" json:"level"`               // 等级
	OrderTotal   int32     `gorm:"column:order_total;type:int;comment:订单总数;default:NULL;" json:"order_total"` // 订单总数
	CreateAt     time.Time `gorm:"column:create_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"create_at"`
	UpdateAt     time.Time `gorm:"column:update_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"update_at"`
}

type MapPoi struct {
	Id        uint64    `gorm:"column:id;type:bigint UNSIGNED;comment:主键ID;primaryKey;not null;" json:"id"`                          // 主键ID
	Name      string    `gorm:"column:name;type:varchar(255);comment:POI名称;not null;" json:"name"`                                   // POI名称
	Address   string    `gorm:"column:address;type:varchar(255);comment:详细地址;not null;" json:"address"`                              // 详细地址
	Province  string    `gorm:"column:province;type:varchar(50);comment:省份;not null;" json:"province"`                               // 省份
	City      string    `gorm:"column:city;type:varchar(50);comment:城市;not null;" json:"city"`                                       // 城市
	District  string    `gorm:"column:district;type:varchar(50);comment:区县;not null;" json:"district"`                               // 区县
	Longitude float64   `gorm:"column:longitude;type:decimal(10, 6);comment:经度;not null;default:0.000000;" json:"longitude"`         // 经度
	Latitude  float64   `gorm:"column:latitude;type:decimal(10, 6);comment:纬度;not null;default:0.000000;" json:"latitude"`           // 纬度
	Tel       string    `gorm:"column:tel;type:varchar(100);comment:联系电话;default:NULL;" json:"tel"`                                  // 联系电话
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;comment:创建时间;not null;default:CURRENT_TIMESTAMP;" json:"created_at"` // 创建时间
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;comment:更新时间;not null;default:CURRENT_TIMESTAMP;" json:"updated_at"` // 更新时间
}

func (u *User) TableName() string {
	return "user"
}
func (u *Driver) TableName() string {
	return "driver"
}
func (u *Auth) TableName() string {
	return "auth"
}
func (u *UserAuth) TableName() string {
	return "user_auth"
}
func (u *DriverAuth) TableName() string {
	return "driver_auth"
}
func (u *MapPoi) TableName() string {
	return "map_poi"
}
