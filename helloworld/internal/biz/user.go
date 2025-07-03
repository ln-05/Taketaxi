package biz

import (
	"github.com/google/wire"
	"gorm.io/gorm"
	"time"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewGreeterUsecase)

type User struct {
	Id       int            `gorm:"column:id;type:int;primaryKey;not null;" json:"id"`
	Mobile   string         `gorm:"column:mobile;type:char(11);default:NULL;" json:"mobile"`
	UserName string         `gorm:"column:user_name;type:varchar(20);default:NULL;" json:"user_name"`
	NickName string         `gorm:"column:nick_name;type:varchar(20);default:NULL;" json:"nick_name"`
	Sex      string         `gorm:"column:sex;type:varchar(5);default:NULL;" json:"sex"`
	Mileage  string         `gorm:"column:mileage;type:varchar(10);default:NULL;" json:"mileage"`
	CreateAt time.Time      `gorm:"column:create_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"create_at"`
	UpdateAt time.Time      `gorm:"column:update_at;type:datetime(3);default:CURRENT_TIMESTAMP(3);" json:"update_at"`
	DeleteAt gorm.DeletedAt `gorm:"column:delete_at;type:datetime(3);default:NULL;" json:"delete_at"`
}

func (u *User) TableName() string {
	return "user"
}
