package service

import (
	"context"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
	"strconv"
)

func (c *GreeterService) DriverRegister(ctx context.Context, in *v1.DriverRegisterRequest) (*v1.DriverRegisterReply, error) {
	var driver biz.Driver
	if err := c.db.GetDB().Where("user_id = ?", in.UserId).First(&driver).Error; err == nil {
		return &v1.DriverRegisterReply{
			Code:    500,
			Message: "你已经是司机了",
		}, nil
	}

	drivers := biz.Driver{
		UserId:       in.UserId,
		DriveStatus:  0,
		OnlineStatus: 0,
		WokeCode:     strconv.Itoa(0),
		Level:        0,
		OrderTotal:   0,
	}

	if err := c.db.GetDB().Create(&drivers).Error; err != nil {
		return &v1.DriverRegisterReply{
			Code:    400,
			Message: "司机入住失败",
		}, nil
	}

	auth := biz.DriverAuth{
		DriverId: drivers.Id,
		IdCord:   in.IdCard,
	}

	if err := c.db.GetDB().Create(&auth).Error; err != nil {
		return &v1.DriverRegisterReply{
			Code:    400,
			Message: "司机认证信息失败",
		}, nil
	}

	return &v1.DriverRegisterReply{
		Code:    200,
		Message: "司机入住成功",
	}, nil
}

func (c *GreeterService) DriverStatus(_ context.Context, in *v1.DriverStatusRequest) (*v1.DriverStatusReply, error) {
	var driver biz.Driver
	if err := c.db.GetDB().Where("id = ?", in.Id).First(&driver).Error; err != nil {
		return &v1.DriverStatusReply{
			Code:    500,
			Message: "当前司机不存在",
		}, nil
	}

	drivers := biz.Driver{
		Id:           in.Id,
		OnlineStatus: int8(in.OnlineStatus),
	}

	if err := c.db.GetDB().Updates(&drivers).Error; err != nil {
		return &v1.DriverStatusReply{
			Code:    500,
			Message: "司机上线失败",
		}, nil
	}

	return &v1.DriverStatusReply{
		Code:    500,
		Message: "司机上线成功",
	}, nil

}
