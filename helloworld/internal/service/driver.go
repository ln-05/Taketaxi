package service

import (
	"context"
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
	"log"
)

// 司机入住，并上传资料提交审核
func (c *GreeterService) DriverRegister(_ context.Context, in *v1.DriverRegisterRequest) (*v1.DriverRegisterReply, error) {

	var driver biz.LxhDriver
	if err := c.db.GetDB().Where("id = ?", in.UserId).First(&driver).Error; err == nil {
		return &v1.DriverRegisterReply{
			Code:    500,
			Message: "司机已存在",
		}, nil
	}

	drivers := biz.LxhDriver{
		Id:        in.UserId,
		Name:      in.Name,
		NickName:  in.NickName,
		AllAcount: 0,
		CarAge:    in.CarAge,
		Status:    "未接单",
		Mobile:    in.Mobile,
	}

	if err := c.db.GetDB().Create(&drivers).Error; err != nil {
		return &v1.DriverRegisterReply{
			Code:    500,
			Message: "司机信息保存失败",
		}, nil
	}

	driverCheck := biz.LxhDriverCheck{
		Id:                   in.UserId,
		IdCardFileId:         in.IdCardFileId,
		DriverLicenseFileId:  in.DriverLicenseFileId,
		DrivingLicenseFileId: in.DrivingLicenseFileId,
		AvatorFileId:         in.AvatorFileId,
		CheckStatus:          "待审核",
	}

	if err := c.db.GetDB().Create(&driverCheck).Error; err != nil {
		return &v1.DriverRegisterReply{
			Code:    500,
			Message: "司机审核信息保存失败",
		}, nil
	}

	return &v1.DriverRegisterReply{
		Code:    200,
		Message: "司机入住申请成功，待审核",
	}, nil

}

// 司机选择接单或者不接单
func (c *GreeterService) DriverStatus(_ context.Context, in *v1.DriverStatusRequest) (*v1.DriverStatusReply, error) {
	var driver biz.LxhDriver
	if err := c.db.GetDB().Where("id = ?", in.Id).First(&driver).Error; err != nil {
		return &v1.DriverStatusReply{
			Code:    500,
			Message: "当前司机不存在",
		}, nil
	}

	drivers := biz.LxhDriver{
		Id:     in.Id,
		Status: in.Status,
	}

	if err := c.db.GetDB().Updates(&drivers).Error; err != nil {
		return &v1.DriverStatusReply{
			Code:    500,
			Message: "开始接单失败",
		}, nil
	}

	return &v1.DriverStatusReply{
		Code:    200,
		Message: "开始接单成功",
	}, nil
}

// 完善司机基本信息
func (c *GreeterService) DriverUpdate(_ context.Context, in *v1.DriverUpdateRequest) (*v1.DriverUpdateReply, error) {

	var driver biz.LxhDriver

	if err := c.db.GetDB().Where("id = ?", in.Id).First(&driver); err != nil {
		return &v1.DriverUpdateReply{
			Code:    400,
			Message: "没有当前司机",
		}, nil
	}
	log.Println(driver)
	drivers := biz.LxhDriver{
		Id:       in.Id,
		Name:     in.Name,
		NickName: in.NickName,
		CarAge:   in.CarAge,
		Mobile:   in.Mobile,
	}

	if err := c.db.GetDB().Updates(&drivers).Error; err != nil {
		return &v1.DriverUpdateReply{
			Code:    500,
			Message: "司机信息更新失败",
		}, nil
	}

	return &v1.DriverUpdateReply{
		Code:    200,
		Message: "司机信息更新成功",
	}, nil
}

// 司机接单
func (c *GreeterService) DriverOrder(_ context.Context, in *v1.DriverOrderRequest) (*v1.DriverOrderReply, error) {
	var order biz.LxhOrder
	if err := c.db.GetDB().Where("order_code = ?", in.OrderCode).First(&order).Error; err != nil {
		return &v1.DriverOrderReply{
			Code:    404,
			Message: "订单不存在",
		}, nil
	}

	if order.OrderStatus != "待派单" {
		return &v1.DriverOrderReply{
			Code:    400,
			Message: "订单状态不允许接单",
		}, nil
	}

	order.OrderStatus = "司机已接单"
	order.Driver = in.DriverId
	if err := c.db.GetDB().Updates(&order).Error; err != nil {
		return &v1.DriverOrderReply{
			Code:    500,
			Message: "接单失败",
		}, nil
	}

	return &v1.DriverOrderReply{
		Code:    200,
		Message: "接单成功",
	}, nil
}
