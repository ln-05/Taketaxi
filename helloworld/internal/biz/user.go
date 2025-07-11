package biz

import (
	"time"

	"github.com/google/wire"
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

// 司机表
type LxhDriver struct {
	Id        int32   `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Name      string  `gorm:"column:name;type:varchar(20);comment:姓名;default:NULL;" json:"name"`                 // 姓名
	NickName  string  `gorm:"column:nick_name;type:varchar(20);comment:昵称;default:NULL;" json:"nick_name"`       // 昵称
	AllAcount float64 `gorm:"column:all_acount;type:decimal(10, 2);comment:总收益;default:0.00;" json:"all_acount"` // 总收益
	CarAge    int32   `gorm:"column:car_age;type:int;comment:车龄;default:NULL;" json:"car_age"`                   // 车龄
	Status    string  `gorm:"column:status;type:varchar(10);comment:是否开启接单;default:NULL;" json:"status"`         // 是否开启接单
	Mobile    string  `gorm:"column:mobile;type:char(11);comment:联系电话;default:NULL;" json:"mobile"`              // 联系电话
	FileId    int32   `gorm:"column:file_id;type:int;comment:头像文件ID;default:NULL;" json:"file_id"`               // 头像文件ID
}

// 司机资料审核表
type LxhDriverCheck struct {
	Id                   int32  `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	IdCardFileId         string `gorm:"column:id_card_file_id;type:varchar(100);comment:身份证;default:NULL;" json:"id_card_file_id"`                 // 身份证
	DriverLicenseFileId  string `gorm:"column:driver_license_file_id;type:varchar(100);comment:驾照;default:NULL;" json:"driver_license_file_id"`    // 驾照
	DrivingLicenseFileId string `gorm:"column:driving_license_file_id;type:varchar(100);comment:行驶证;default:NULL;" json:"driving_license_file_id"` // 行驶证
	AvatorFileId         string `gorm:"column:avator_file_id;type:varchar(100);comment:司机头像;default:NULL;" json:"avator_file_id"`                  // 司机头像
	CheckStatus          string `gorm:"column:check_status;type:varchar(5);comment:审核状态;default:NULL;" json:"check_status"`                        // 审核状态
	Remark               string `gorm:"column:remark;type:varchar(50);comment:备注;default:NULL;" json:"remark"`                                     // 备注
}

// 订单表
type LxhOrder struct {
	Id            int32      `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	OrderCode     string     `gorm:"column:order_code;type:varchar(50);comment:订单编号;default:NULL;" json:"order_code"`            // 订单编号
	Amount        float64    `gorm:"column:amount;type:decimal(10, 2);comment:总价;default:NULL;" json:"amount"`                   // 总价
	OrderStatus   string     `gorm:"column:order_status;type:varchar(10);comment:订单状态;default:NULL;" json:"order_status"`        // 订单状态
	PassengerId   int32      `gorm:"column:passenger_id;type:int;comment:乘客ID;default:NULL;" json:"passenger_id"`                // 乘客ID
	StartAddr     string     `gorm:"column:start_addr;type:varchar(50);comment:起始地;default:NULL;" json:"start_addr"`             // 起始地
	EndEnd        string     `gorm:"column:end_end;type:varchar(50);comment:目的地;default:NULL;" json:"end_end"`                   // 目的地
	Driver        int32      `gorm:"column:driver;type:int;comment:司机ID;default:NULL;" json:"driver"`                            // 司机ID
	StartTime     time.Time  `gorm:"column:start_time;type:datetime;comment:开始时间;default:NULL;" json:"start_time"`               // 开始时间
	EndTime       *time.Time `gorm:"column:end_time;type:datetime;comment:结束时间;default:NULL;" json:"end_time"`                   // 结束时间
	ConfirmBy     string     `gorm:"column:confirm_by;type:varchar(5);comment:取消方（乘客/司机）;default:NULL;" json:"confirm_by"`       // 取消方（乘客/司机）
	ConfirmPerson int32      `gorm:"column:confirm_person;type:int;comment:取消人;default:NULL;" json:"confirm_person"`             // 取消人
	ConfirmReason string     `gorm:"column:confirm_reason;type:varchar(50);comment:取消原因;default:NULL;" json:"confirm_reason"`    // 取消原因
	ConfirmRemark string     `gorm:"column:confirm_remark;type:varchar(50);comment:取消备注;default:NULL;" json:"confirm_remark"`    // 取消备注
	PayStatus     string     `gorm:"column:pay_status;type:varchar(20);comment:支付状态;default:NULL;" json:"pay_status"`            // 支付状态
	PayType       string     `gorm:"column:pay_type;type:varchar(20);comment:支付方式;default:NULL;" json:"pay_type"`                // 支付方式
	OrderType     string     `gorm:"column:order_type;type:varchar(20);comment:订单状态（快车订单/预约订单）;default:NULL;" json:"order_type"` // 订单状态（快车订单/预约订单）
}

// 订单扩展表
type LxhOrderDetail struct {
	Id        int32  `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	OrderCode string `gorm:"column:order_code;type:varchar(50);comment:订单编号;default:NULL;" json:"order_code"`         // 订单编号
	TripKey   string `gorm:"column:trip_key;type:varchar(50);comment:mongdb中行程的key;default:NULL;" json:"trip_key"`    // mongdb中行程的key
	DriverKey string `gorm:"column:driver_key;type:varchar(50);comment:司机实际行程路线的key;default:NULL;" json:"driver_key"` // 司机实际行程路线的key
}

// 乘客表
type LxhPassenger struct {
	Id       int32  `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Name     string `gorm:"column:name;type:varchar(20);comment:姓名;default:NULL;" json:"name"`           // 姓名
	NickName string `gorm:"column:nick_name;type:varchar(20);comment:昵称;default:NULL;" json:"nick_name"` // 昵称
	FileId   int32  `gorm:"column:file_id;type:int;comment:头像文件ID;default:NULL;" json:"file_id"`         // 头像文件ID
	Mobile   string `gorm:"column:mobile;type:char(11);comment:手机号;default:NULL;" json:"mobile"`         // 手机号
	Age      int32  `gorm:"column:age;type:int;comment:年龄;default:NULL;" json:"age"`                     // 年龄
	Sex      string `gorm:"column:sex;type:varchar(2);comment:性别;default:NULL;" json:"sex"`              // 性别
	Mileage  string `gorm:"column:mileage;type:varchar(10);comment:里程数;default:NULL;" json:"mileage"`    // 里程数
}

// 行程表
type Trip struct {
	Id                int32     `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	UserId            int64     `gorm:"column:user_id;type:bigint;comment:用户ID;not null;" json:"user_id"`                                 // 用户ID
	DriverId          int64     `gorm:"column:driver_id;type:bigint;comment:司机ID;default:NULL;" json:"driver_id"`                         // 司机ID
	StartLng          float64   `gorm:"column:start_lng;type:decimal(9, 6);comment:起点经度;not null;" json:"start_lng"`                      // 起点经度
	StartLat          float64   `gorm:"column:start_lat;type:decimal(9, 6);comment:起点纬度;not null;" json:"start_lat"`                      // 起点纬度
	EndLng            float64   `gorm:"column:end_lng;type:decimal(9, 6);comment:终点经度;not null;" json:"end_lng"`                          // 终点经度
	EndLat            float64   `gorm:"column:end_lat;type:decimal(9, 6);comment:终点纬度;not null;" json:"end_lat"`                          // 终点纬度
	StartAddress      string    `gorm:"column:start_address;type:varchar(100);comment:起点文字地址;not null;" json:"start_address"`             // 起点文字地址
	EndAddress        string    `gorm:"column:end_address;type:varchar(100);comment:终点文字地址;not null;" json:"end_address"`                 // 终点文字地址
	VehicleType       int8      `gorm:"column:vehicle_type;type:tinyint;comment:车型 1:"快车" 2:"顺风车" 3:"豪华车";not null;" json:"vehicle_type"` // 车型 1:"快车" 2:"顺风车" 3:"豪华车"
	EstimatedDistance uint32    `gorm:"column:estimated_distance;type:int UNSIGNED;comment:预估距离;not null;" json:"estimated_distance"`     // 预估距离
	ActualDistance    uint32    `gorm:"column:actual_distance;type:int UNSIGNED;comment:实际距离;default:NULL;" json:"actual_distance"`       // 实际距离
	TotalFare         float64   `gorm:"column:total_fare;type:decimal(8, 2);comment:总费用;not null;" json:"total_fare"`                     // 总费用
	CreateTime        time.Time `gorm:"column:create_time;type:datetime;comment:下单时间;not null;" json:"create_time"`                       // 下单时间
	AcceptTime        time.Time `gorm:"column:accept_time;type:datetime;comment:司机接单时间;default:NULL;" json:"accept_time"`                 // 司机接单时间
	StartTime         time.Time `gorm:"column:start_time;type:datetime;comment:上车时间;default:NULL;" json:"start_time"`                     // 上车时间
}

func (u *User) TableName() string {
	return "user"
}
func (u *Auth) TableName() string {
	return "auth"
}
func (u *UserAuth) TableName() string {
	return "user_auth"
}
func (u *MapPoi) TableName() string {
	return "map_poi"
}
func (u *LxhDriver) TableName() string {
	return "lxh_driver"
}
func (u *LxhDriverCheck) TableName() string {
	return "lxh_driver_check"
}
func (u *LxhOrder) TableName() string {
	return "lxh_order"
}
func (u *LxhOrderDetail) TableName() string {
	return "lxh_order_detail"
}
func (u *LxhPassenger) TableName() string {
	return "lxh_passenger"
}
func (u *Trip) TableName() string {
	return "trip"
}
