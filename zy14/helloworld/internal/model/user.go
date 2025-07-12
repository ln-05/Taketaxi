package model

import (
	"gorm.io/gorm"
	"time"
)

// 授权表
type Auth struct {
	Id          int64  `gorm:"column:id;type:int;primaryKey;autoIncrement;not null;" json:"id"`        // 主键
	AppName     string `gorm:"column:app_name;type:varchar(20);default:NULL;" json:"app_name"`         // 应用名称
	AppUnionid  string `gorm:"column:app_unionid;type:varchar(60);default:NULL;" json:"app_unionid"`   // 应用唯一ID
	AccessToken string `gorm:"column:access_token;type:varchar(60);default:NULL;" json:"access_token"` // 访问令牌
	AppRemark   string `gorm:"column:app_remark;type:varchar(255);default:NULL;" json:"app_remark"`    // 备注
}

func (a *Auth) TableName() string {
	return "auth"
}

// 司机表
type LxhDriver struct {
	Id                int64   `gorm:"column:id;type:int;primaryKey;autoIncrement;not null;" json:"id"`              // 主键
	UserId            int64   `gorm:"column:user_id;type:int;default:NULL;" json:"user_id"`                         // 用户ID
	DriverName        string  `gorm:"column:driver_name;type:varchar(50);default:NULL;" json:"driver_name"`         // 司机姓名
	DriverLicense     string  `gorm:"column:driver_license;type:varchar(50);default:NULL;" json:"driver_license"`   // 驾驶证号
	VehicleLicense    string  `gorm:"column:vehicle_license;type:varchar(50);default:NULL;" json:"vehicle_license"` // 行驶证号
	VehicleType       string  `gorm:"column:vehicle_type;type:varchar(50);default:NULL;" json:"vehicle_type"`       // 车辆类型
	ApplicationStatus int     `gorm:"column:application_status;type:int;default:NULL;" json:"application_status"`   // 申请状态
	RejectReason      string  `gorm:"column:reject_reason;type:varchar(255);default:NULL;" json:"reject_reason"`    // 拒绝原因
	AuditorId         int64   `gorm:"column:auditor_id;type:int;default:NULL;" json:"auditor_id"`                   // 审核人id
	Name              string  `gorm:"column:name;type:varchar(20);default:NULL;" json:"name"`                       // 姓名
	NickName          string  `gorm:"column:nick_name;type:varchar(20);default:NULL;" json:"nick_name"`             // 昵称
	AllAcount         float64 `gorm:"column:all_acount;type:decimal(10,2);default:NULL;" json:"all_acount"`         // 总收益
	CarAge            int     `gorm:"column:car_age;type:int;default:NULL;" json:"car_age"`                         // 车龄
	Status            string  `gorm:"column:status;type:varchar(10);default:NULL;" json:"status"`                   // 是否开启接单
	Mobile            string  `gorm:"column:mobile;type:char(11);default:NULL;" json:"mobile"`                      // 联系电话
	FileId            int     `gorm:"column:file_id;type:int;default:NULL;" json:"file_id"`                         // 头像文件ID
}

func (d *LxhDriver) TableName() string {
	return "lxh_driver"
}

func (d *LxhDriver) Create(db *gorm.DB) error {
	return db.Create(d).Error
}

func (d *LxhDriver) GetUserBy(db *gorm.DB, mobile string) error {
	return db.Where("mobile = ?", mobile).First(d).Error
}

func (d *LxhDriver) Update(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).Updates(d).Error
}

// 司机资料审核表
type LxhDriverCheck struct {
	Id                   int64  `gorm:"column:id;type:int;primaryKey;autoIncrement;not null;" json:"id"`                               // 主键
	IdCardFileId         string `gorm:"column:id_card_file_id;type:varchar(100);default:NULL;" json:"id_card_file_id"`                 // 身份证
	DriverLicenseFileId  string `gorm:"column:driver_license_file_id;type:varchar(100);default:NULL;" json:"driver_license_file_id"`   // 驾照
	DrivingLicenseFileId string `gorm:"column:driving_license_file_id;type:varchar(100);default:NULL;" json:"driving_license_file_id"` // 行驶证
	AvatorFileId         string `gorm:"column:avator_file_id;type:varchar(100);default:NULL;" json:"avator_file_id"`                   // 司机头像
	CheckStatus          string `gorm:"column:check_status;type:varchar(5);default:NULL;" json:"check_status"`                         // 审核状态
	Remark               string `gorm:"column:remark;type:varchar(50);default:NULL;" json:"remark"`                                    // 备注
}

func (c *LxhDriverCheck) TableName() string {
	return "lxh_driver_check"
}

// 订单表
type LxhOrder struct {
	Id            int64      `gorm:"column:id;type:int;primaryKey;autoIncrement;not null;" json:"id"`            // 主键
	OrderCode     string     `gorm:"column:order_code;type:varchar(50);default:NULL;" json:"order_code"`         // 订单编号
	Amount        float64    `gorm:"column:amount;type:decimal(10,2);default:NULL;" json:"amount"`               // 总价
	OrderStatus   string     `gorm:"column:order_status;type:varchar(10);default:NULL;" json:"order_status"`     // 订单状态
	PassengerId   int        `gorm:"column:passenger_id;type:int;default:NULL;" json:"passenger_id"`             // 乘客ID
	StartAddr     string     `gorm:"column:start_addr;type:varchar(50);default:NULL;" json:"start_addr"`         // 起始地
	EndEnd        string     `gorm:"column:end_end;type:varchar(50);default:NULL;" json:"end_end"`               // 目的地
	Driver        int        `gorm:"column:driver;type:int;default:NULL;" json:"driver"`                         // 司机ID
	StartTime     *time.Time `gorm:"column:start_time;type:datetime;default:NULL;" json:"start_time"`            // 开始时间
	EndTime       *time.Time `gorm:"column:end_time;type:datetime;default:NULL;" json:"end_time"`                // 结束时间
	ConfirmBy     string     `gorm:"column:confirm_by;type:varchar(5);default:NULL;" json:"confirm_by"`          // 取消方
	ConfirmPerson int        `gorm:"column:confirm_person;type:int;default:NULL;" json:"confirm_person"`         // 取消人
	ConfirmReason string     `gorm:"column:confirm_reason;type:varchar(50);default:NULL;" json:"confirm_reason"` // 取消原因
	ConfirmRemark string     `gorm:"column:confirm_remark;type:varchar(50);default:NULL;" json:"confirm_remark"` // 取消备注
	PayStatus     string     `gorm:"column:pay_status;type:varchar(20);default:NULL;" json:"pay_status"`         // 支付状态
	PayType       string     `gorm:"column:pay_type;type:varchar(20);default:NULL;" json:"pay_type"`             // 支付方式
	OrderType     string     `gorm:"column:order_type;type:varchar(20);default:NULL;" json:"order_type"`         // 订单类型
}

func (o *LxhOrder) TableName() string {
	return "lxh_order"
}

// LxhOrder表相关时间处理工具
func FormatTimePtr(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02 15:04:05")
}

func ParseTimePtr(s string) *time.Time {
	if s == "" {
		return nil
	}
	t, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return nil
	}
	return &t
}

// 订单拓展表
type LxhOrderDetail struct {
	Id        int64  `gorm:"column:id;type:int;primaryKey;autoIncrement;not null;" json:"id"`    // 主键
	OrderCode string `gorm:"column:order_code;type:varchar(50);default:NULL;" json:"order_code"` // 订单编号
	TripKey   string `gorm:"column:trip_key;type:varchar(50);default:NULL;" json:"trip_key"`     // 行程key
	DriverKey string `gorm:"column:driver_key;type:varchar(50);default:NULL;" json:"driver_key"` // 司机实际行程key
}

func (d *LxhOrderDetail) TableName() string {
	return "lxh_order_detail"
}

// 乘客表
type LxhPassenger struct {
	Id       int64  `gorm:"column:id;type:int;primaryKey;autoIncrement;not null;" json:"id"`  // 主键
	Name     string `gorm:"column:name;type:varchar(20);default:NULL;" json:"name"`           // 姓名
	NickName string `gorm:"column:nick_name;type:varchar(20);default:NULL;" json:"nick_name"` // 昵称
	FileId   int    `gorm:"column:file_id;type:int;default:NULL;" json:"file_id"`             // 头像文件ID
	Mobile   string `gorm:"column:mobile;type:char(11);default:NULL;" json:"mobile"`          // 手机号
	Age      int    `gorm:"column:age;type:int;default:NULL;" json:"age"`                     // 年龄
	Sex      string `gorm:"column:sex;type:varchar(2);default:NULL;" json:"sex"`              // 性别
	Mileage  string `gorm:"column:mileage;type:varchar(10);default:NULL;" json:"mileage"`     // 里程数
}

func (p *LxhPassenger) TableName() string {
	return "lxh_passenger"
}

// 用户表
type User struct {
	Id       int64          `gorm:"column:id;type:int;primaryKey;autoIncrement;not null;" json:"id"`                  // 主键
	Mobile   string         `gorm:"column:mobile;type:char(11);default:NULL;" json:"mobile"`                          // 手机号
	UserName string         `gorm:"column:user_name;type:varchar(20);default:NULL;" json:"user_name"`                 // 用户名
	NickName string         `gorm:"column:nick_name;type:varchar(20);default:NULL;" json:"nick_name"`                 // 昵称
	IdCard   string         `gorm:"column:id_card;type:varchar(18);default:NULL;" json:"id_card"`                     // 身份证
	Sex      string         `gorm:"column:sex;type:varchar(5);default:NULL;" json:"sex"`                              // 性别
	Mileage  string         `gorm:"column:mileage;type:varchar(10);default:NULL;" json:"mileage"`                     // 里程数
	CreateAt *time.Time     `gorm:"column:create_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"create_at"` // 创建时间
	UpdateAt *time.Time     `gorm:"column:update_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"update_at"` // 更新时间
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;type:datetime(3);default:NULL;" json:"delete_at"`                 // 删除时间
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}

func (u *User) GetUserBy(db *gorm.DB, mobile string) error {
	return db.Where("mobile=?", mobile).First(&u).Error
}

// 用户授权表
type UserAuth struct {
	Id     int64 `gorm:"column:id;type:int;primaryKey;autoIncrement;not null;" json:"id"` // 主键
	UserId int   `gorm:"column:user_id;type:int;default:NULL;" json:"user_id"`            // 用户ID
	AuthId int   `gorm:"column:auth_id;type:int;default:NULL;" json:"auth_id"`            // 授权ID
}

func (a *UserAuth) TableName() string {
	return "user_auth"
}

type DriverEvaluation struct {
	Id          uint      `gorm:"column:id;type:bigint UNSIGNED;comment:评价唯一标识;primaryKey;not null;" json:"id"`                          // 评价唯一标识
	OrderId     uint      `gorm:"column:order_id;type:bigint UNSIGNED;comment:关联订单 id，关联 order_info 表;not null;" json:"order_id"`        // 关联订单 id，关联 order_info 表
	PassengerId uint      `gorm:"column:passenger_id;type:bigint UNSIGNED;comment:乘客用户 id，关联 user_info 表;not null;" json:"passenger_id"` // 乘客用户 id，关联 user_info 表
	DriverId    uint      `gorm:"column:driver_id;type:bigint UNSIGNED;comment:司机用户 id，关联 driver 表;not null;" json:"driver_id"`          // 司机用户 id，关联 driver 表
	Score       uint      `gorm:"column:score;type:tinyint;comment:评价分数，1 - 5 分;not null;" json:"score"`                                 // 评价分数，1 - 5 分
	Content     string    `gorm:"column:content;type:varchar(500);comment:评价内容;default:NULL;" json:"content"`                            // 评价内容
	EvalTime    time.Time `gorm:"column:eval_time;type:datetime;comment:评价时间;not null;default:CURRENT_TIMESTAMP;" json:"eval_time"`      // 评价时间
}

func (d *DriverEvaluation) TableName() string {
	return "driver_evaluation"
}

func (d *DriverEvaluation) Create(db *gorm.DB) error {
	return db.Create(&d).Error
}

type UserSearchHistory struct {
	Id         uint           `gorm:"column:id;type:bigint;primaryKey;not null;" json:"id"`
	UserId     uint           `gorm:"column:user_id;type:bigint;not null;" json:"user_id"`
	Keyword    string         `gorm:"column:keyword;type:varchar(255);not null;" json:"keyword"`
	SearchTime time.Time      `gorm:"column:search_time;type:datetime(3);not null;default:CURRENT_TIMESTAMP(3);" json:"search_time"`
	DeletedAt  gorm.DeletedAt `gorm:"column:deleted_at;type:datetime(3);default:NULL;" json:"deleted_at"`
}

func (u *UserSearchHistory) TableName() string {
	return "user_search_history"
}

func (u *UserSearchHistory) Create(db *gorm.DB) error {
	return db.Create(&u).Error
}
