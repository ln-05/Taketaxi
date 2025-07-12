package service

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	pb "helloworld/api/helloworld/v1"
	"helloworld/internal/data"
	"helloworld/internal/model"
	"helloworld/internal/pkg"
	"math/rand"
	"strings"
	"time"
)

type DriverService struct {
	pb.UnimplementedDriverServer
	data *data.Data
}

func NewDriverService(data *data.Data) *DriverService {
	return &DriverService{
		data: data,
	}
}

func (s *DriverService) SendSms(ctx context.Context, req *pb.DriverSendSmsRequest) (*pb.DriverSendSmsResponse, error) {
	code := rand.Intn(9000) + 1000

	if !pkg.Regex(req.Mobile) {
		return &pb.DriverSendSmsResponse{
			Code:    10000,
			Message: "手机号格式错误",
		}, nil
	}

	get := s.data.Rdb.Get(context.Background(), "send"+req.Mobile+req.Source)
	if get.Val() != "" {
		return nil, fmt.Errorf("短信每分钟只能发送一次")
	}

	//sms, err := pkg.SendSms(in.Mobile, strconv.Itoa(code))
	//if err != nil {
	//	return nil, err
	//}
	//
	//if *sms.Body.Code != "OK" {
	//	return nil, fmt.Errorf(*sms.Body.Message)
	//}

	s.data.Rdb.Set(context.Background(), "sendSms"+req.Mobile+req.Source, code, time.Minute*5)
	s.data.Rdb.Set(context.Background(), "Send"+req.Mobile+req.Source, code, time.Minute*1)

	return &pb.DriverSendSmsResponse{
		Code:    200,
		Message: "短信发送成功",
	}, nil
}
func (s *DriverService) Login(ctx context.Context, req *pb.DriverLoginRequest) (*pb.DriverLoginResponse, error) {
	var d model.LxhDriver
	d.GetUserBy(s.data.DB, req.Mobile)

	if d.Id != 0 {

		d = model.LxhDriver{
			Mobile: req.Mobile,
		}

		if !pkg.Regex(req.Mobile) {
			return &pb.DriverLoginResponse{
				Code:    10000,
				Message: "手机号格式错误",
			}, nil
		}

		err := s.data.Rdb.Get(context.Background(), "sendSms"+req.Source+"register").Err()
		if err != nil {
			return &pb.DriverLoginResponse{
				Code:    10000,
				Message: "短信发送失败",
			}, nil
		}

		err = d.Create(s.data.DB)
		if err != nil {
			return &pb.DriverLoginResponse{
				Code:    10000,
				Message: "注册失败",
			}, nil
		}

		return &pb.DriverLoginResponse{
			Code:    200,
			Message: "注册成功",
		}, nil
	} else {
		if !pkg.Regex(req.Mobile) {
			return &pb.DriverLoginResponse{
				Code:    10000,
				Message: "手机号格式错误",
			}, nil
		}

		err := s.data.Rdb.Get(context.Background(), "sendSms"+req.Source+"login").Err()
		if err != nil {
			return &pb.DriverLoginResponse{
				Code:    10000,
				Message: "短信发送失败",
			}, nil
		}

		token, _ := pkg.NewJWT("2211a").CreateToken(pkg.CustomClaims{
			ID: uint(d.Id),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + 86400,
			},
		})

		return &pb.DriverLoginResponse{
			Code:    200,
			Message: token,
		}, nil
	}
}

func (s *DriverService) Update(ctx context.Context, req *pb.UpdateDriverRequest) (*pb.UpdateDriverResponse, error) {
	var driver model.LxhDriver
	if err := s.data.DB.Where("id = ?", req.DriverId).First(&driver).Error; err != nil {
		return &pb.UpdateDriverResponse{
			Code:    10000,
			Message: "司机不存在",
		}, nil
	}

	var d model.LxhDriver

	d = model.LxhDriver{
		Id:                req.Id,
		UserId:            req.UserId,
		DriverName:        req.DriverName,
		DriverLicense:     req.DriverLicense,
		VehicleLicense:    req.VehicleLicense,
		VehicleType:       req.VehicleType,
		ApplicationStatus: int(req.ApplicationStatus),
		RejectReason:      req.RejectReason,
		AuditorId:         req.AuditorId,
	}

	d.Update(s.data.DB, uint(req.Id))

	return &pb.UpdateDriverResponse{
		Code:    200,
		Message: "司机信息完善成功",
	}, nil
}

// 司机实名认证
func (s *DriverService) RealName(ctx context.Context, req *pb.RealNameRequest) (*pb.RealNameResponse, error) {

	var driver model.LxhDriver
	if err := s.data.DB.Where("id = ?", req.DriverId).First(&driver).Error; err != nil {
		return &pb.RealNameResponse{
			Code:    10000,
			Message: "司机不存在",
		}, nil
	}

	if !pkg.RealName(req.IdCard, req.RealName) {
		return &pb.RealNameResponse{
			Code:    10000,
			Message: "实名认证失败",
		}, nil
	}

	return &pb.RealNameResponse{
		Code:    200,
		Message: "实名认证成功",
	}, nil
}

// 首页统计接口
func (s *DriverService) HomeStat(ctx context.Context, req *pb.HomeStatRequest) (*pb.HomeStatResponse, error) {
	var drivers model.LxhDriver
	if err := s.data.DB.Where("id = ?", req.DriverId).First(&drivers).Error; err != nil {
		return &pb.HomeStatResponse{
			Code:    10000,
			Message: "司机不存在",
		}, nil
	}

	// 今日收益
	var todayIncome float64
	err := s.data.DB.Model(&model.LxhOrder{}).
		Where("driver = ? AND DATE(start_time) = ? AND pay_status = ?", req.DriverId, time.Now().Format("2006-01-02"), "payed").
		Select("SUM(actual_price)").Scan(&todayIncome).Error
	if err != nil {
		todayIncome = 0
	}

	// 近7天平均收益
	var weekIncome float64
	var weekOrderCount int64
	weekAgo := time.Now().AddDate(0, 0, -6).Format("2006-01-02")
	err = s.data.DB.Model(&model.LxhOrder{}).
		Where("driver = ? AND DATE(start_time) >= ? AND order_status NOT IN (?, ?)", req.DriverId, weekAgo, "异常", "已取消").
		Select("SUM(actual_price)").Scan(&weekIncome).Error
	if err != nil {
		weekIncome = 0
	}
	err = s.data.DB.Model(&model.LxhOrder{}).
		Where("driver = ? AND DATE(start_time) >= ? AND order_status NOT IN (?, ?)", req.DriverId, weekAgo, "异常", "已取消").
		Count(&weekOrderCount).Error
	if err != nil {
		weekOrderCount = 0
	}
	avgIncome := 0.0
	if weekOrderCount > 0 {
		avgIncome = weekIncome / float64(weekOrderCount)
	}

	// 剩余拒单次数
	var driver model.LxhDriver
	if err := s.data.DB.Where("id = ?", req.DriverId).First(&driver).Error; err != nil {
		return &pb.HomeStatResponse{
			Code:    10002,
			Message: "司机不存在",
		}, nil
	}
	dailyRejectLimit := 3
	var todayRejectCount int64
	err = s.data.DB.Model(&model.LxhOrder{}).
		Where("driver = ? AND DATE(start_time) = ? AND order_status = ?", req.DriverId, time.Now().Format("2006-01-02"), "rejected").
		Count(&todayRejectCount).Error
	if err != nil {
		todayRejectCount = 0
	}
	remainRejectCount := int32(dailyRejectLimit) - int32(todayRejectCount)
	if remainRejectCount < 0 {
		remainRejectCount = 0
	}

	// 七日成单率
	var weekSuccessCount int64
	var weekTotalCount int64
	err = s.data.DB.Model(&model.LxhOrder{}).
		Where("driver = ? AND DATE(start_time) >= ? AND order_status = ?", req.DriverId, weekAgo, "已完成").
		Count(&weekSuccessCount).Error
	if err != nil {
		weekSuccessCount = 0
	}
	err = s.data.DB.Model(&model.LxhOrder{}).
		Where("driver = ? AND DATE(start_time) >= ?", req.DriverId, weekAgo).
		Count(&weekTotalCount).Error
	if err != nil {
		weekTotalCount = 0
	}
	weekSuccessRate := 0.0
	if weekTotalCount > 0 {
		weekSuccessRate = float64(weekSuccessCount) / float64(weekTotalCount) * 100
	}

	return &pb.HomeStatResponse{
		TodayIncome:      todayIncome,
		AvgIncome:        avgIncome,
		RemainOrderCount: remainRejectCount,
		WeekSuccessRate:  weekSuccessRate,
		Code:             200,
		Message:          "首页统计获取成功",
	}, nil
}

// 司机出车
func (s *DriverService) Depart(ctx context.Context, req *pb.DriverDepartRequest) (*pb.DriverDepartResponse, error) {
	//  校验司机ID是否存在
	var driver model.LxhDriver
	if err := s.data.DB.Where("id = ?", req.DriverId).First(&driver).Error; err != nil {
		return &pb.DriverDepartResponse{
			Code:    10000,
			Message: "司机不存在",
		}, nil
	}

	//  校验车辆ID是否存在且绑定该司机
	var vehicle model.LxhDriver
	if err := s.data.DB.Where("id = ? AND driver_id = ?", req.VehicleId, req.DriverId).First(&vehicle).Error; err != nil {
		return &pb.DriverDepartResponse{
			Code:    10000,
			Message: "车辆不存在或未绑定该司机",
		}, nil
	}

	//  校验司机当前状态（必须为离线）
	var statusDriver model.LxhDriver
	if err := s.data.DB.Where("id = ? AND status = ?", req.DriverId, "offline").First(&statusDriver).Error; err != nil {
		return &pb.DriverDepartResponse{
			Code:    10000,
			Message: "司机当前不是离线状态，无法出车",
		}, nil
	}

	//  更新司机状态为“出车”
	if err := s.data.DB.Model(&model.LxhDriver{}).Where("id = ?", req.DriverId).Update("status", "departed").Error; err != nil {
		return &pb.DriverDepartResponse{
			Code:    10000,
			Message: "出车失败，状态更新异常",
		}, nil
	}
	return &pb.DriverDepartResponse{
		Code:    200,
		Message: "出车成功",
	}, nil
}

// 司机接单
func (s *DriverService) AcceptOrder(ctx context.Context, req *pb.DriverAcceptOrderRequest) (*pb.DriverAcceptOrderResponse, error) {
	// 校验订单是存在和接单状态
	var order model.LxhOrder
	if err := s.data.DB.Where("id = ? AND order_status = ?", req.OrderId, "pending").First(&order).Error; err != nil {
		return &pb.DriverAcceptOrderResponse{
			Code:    10000,
			Message: "订单不存在或状态不可接单",
		}, nil
	}

	// 校验司机是否存在和在线状态
	var driver model.LxhDriver
	if err := s.data.DB.Where("id = ? AND status = ?", req.DriverId, "online").First(&driver).Error; err != nil {
		return &pb.DriverAcceptOrderResponse{
			Code:    10000,
			Message: "司机不存在或不在线",
		}, nil
	}

	// 距离校验
	if order.StartAddr != "" && driver.Name != "" && !strings.Contains(order.StartAddr, driver.Name) {
		return &pb.DriverAcceptOrderResponse{
			Code:    10000,
			Message: "订单起点与司机信息不符，无法接单",
		}, nil
	}

	// 车型要求
	if order.OrderType != "" && req.CarType != order.OrderType {
		return &pb.DriverAcceptOrderResponse{Code: 10000,
			Message: "车型不符，无法接单",
		}, nil
	}

	// 防重复接单
	if order.OrderStatus != "pending" {
		return &pb.DriverAcceptOrderResponse{
			Code:    10000,
			Message: "订单已被接单",
		}, nil
	}

	// 更新订单状态为已接单，记录司机ID
	if err := s.data.DB.Model(&model.LxhOrder{}).Where("id = ?", req.OrderId).Updates(map[string]interface{}{
		"order_status": "accepted",
		"driver":       req.DriverId,
	}).Error; err != nil {
		return &pb.DriverAcceptOrderResponse{
			Code:    10000,
			Message: "接单失败，状态更新异常",
		}, nil
	}
	return &pb.DriverAcceptOrderResponse{
		Code:    200,
		Message: "接单成功",
	}, nil
}
