package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-redis/redis/v8"
	pb "helloworld/api/helloworld/v1"
	"helloworld/internal/data"
	"helloworld/internal/model"
	"helloworld/internal/pkg"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type UserService struct {
	pb.UnimplementedUserServer
	data *data.Data
}

func NewUserService(data *data.Data) *UserService {
	return &UserService{
		data: data,
	}
}

// 短信发送
func (s *UserService) SendSms(ctx context.Context, req *pb.UserSendSmsRequest) (*pb.UserSendSmsResponse, error) {
	code := rand.Intn(9000) + 1000

	if !pkg.Regex(req.Mobile) {
		return &pb.UserSendSmsResponse{
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

	return &pb.UserSendSmsResponse{
		Code:    200,
		Message: "短信发送成功",
	}, nil
}

// 登录注册一体化
func (s *UserService) Login(ctx context.Context, req *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {

	var u model.User
	u.GetUserBy(s.data.DB, req.Mobile)

	if u.Id != 0 {

		u = model.User{
			Mobile: req.Mobile,
		}

		if !pkg.Regex(req.Mobile) {
			return &pb.UserLoginResponse{
				Code:    10000,
				Message: "手机号格式错误",
			}, nil
		}

		err := s.data.Rdb.Get(context.Background(), "sendSms"+req.Source+"register").Err()
		if err != nil {
			return &pb.UserLoginResponse{
				Code:    10000,
				Message: "短信发送失败",
			}, nil
		}

		err = u.Create(s.data.DB)
		if err != nil {
			return &pb.UserLoginResponse{
				Code:    10000,
				Message: "注册失败",
			}, nil
		}

		return &pb.UserLoginResponse{
			Code:    200,
			Message: "注册成功",
		}, nil
	} else {
		if !pkg.Regex(req.Mobile) {
			return &pb.UserLoginResponse{
				Code:    10000,
				Message: "手机号格式错误",
			}, nil
		}

		err := s.data.Rdb.Get(context.Background(), "sendSms"+req.Source+"login").Err()
		if err != nil {
			return &pb.UserLoginResponse{
				Code:    10000,
				Message: "短信发送失败",
			}, nil
		}

		token, _ := pkg.NewJWT("2211a").CreateToken(pkg.CustomClaims{
			ID: uint(u.Id),
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Unix() + 86400,
			},
		})

		return &pb.UserLoginResponse{
			Code:    200,
			Message: token,
		}, nil
	}

}

// 完善信息
func (s *UserService) Update(ctx context.Context, req *pb.UserUpdateRequest) (*pb.UserUpdateResponse, error) {
	//验证用户是否存在
	var user model.User
	err := s.data.DB.Where("id = ?", req.Id).First(&user).Error
	if err != nil {
		return &pb.UserUpdateResponse{
			Code:    10000,
			Message: "用户不存在",
		}, nil
	}

	u := &model.User{
		Id:       req.Id,
		Mobile:   req.Mobile,
		UserName: req.UserName,
		NickName: req.NickName,
		IdCard:   req.IdCard,
		Sex:      req.Sex,
		Mileage:  req.Mileage,
	}

	err = s.data.DB.Save(u).Error
	if err != nil {
		return &pb.UserUpdateResponse{
			Code: 10000, Message: "更新失败",
		}, nil
	}

	return &pb.UserUpdateResponse{
		Code:    200,
		Message: "用户完善信息成功",
	}, nil
}

// 获取当前位置
// 通过百度地图IP定位API获取用户当前位置信息
// 打车应用中的用户定位、默认起点设置、附近服务查询等
func (s *UserService) GetCurrentLocation(ctx context.Context, req *pb.GetCurrentLocationRequest) (*pb.GetCurrentLocationResponse, error) {
	// 验证用户是否存在
	var user model.User
	err := s.data.DB.Where("id = ?", req.UserId).First(&user).Error
	if err != nil {
		return &pb.GetCurrentLocationResponse{
			Code:    10000,
			Message: "用户不存在",
		}, nil
	}

	//  调用百度地图IP定位API获取当前位置
	url := fmt.Sprintf("https://api.map.baidu.com/location/ip?ak=%s&coor=bd09ll", req.BaiduAk)

	// 配置HTTP客户端，设置10秒超时避免长时间等待
	httpClient := &http.Client{
		Timeout: 10 * time.Second, // 超时设置，防止网络问题导致接口长时间无响应
	}

	// 发起HTTP GET请求
	resp, err := httpClient.Get(url)
	if err != nil {
		return &pb.GetCurrentLocationResponse{
			Code:    10000,
			Message: fmt.Sprintf("请求百度地图API失败: %v", err),
		}, nil
	}
	defer resp.Body.Close() // 确保响应体被正确关闭，避免内存泄漏

	// 读取响应内容
	// 使用io.ReadAll一次性读取所有响应数据到内存
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return &pb.GetCurrentLocationResponse{
			Code:    10000,
			Message: fmt.Sprintf("读取响应失败: %v", err),
		}, nil
	}

	// 解析JSON响应
	// 定义匿名结构体用于解析百度地图API返回的JSON数据
	// 使用匿名结构体的原因：只用于一次性解析，不需要复用，减少代码复杂度
	var baiduResp struct {
		Status  int    `json:"status"`  // API响应状态码：0表示成功，其他表示错误
		Message string `json:"message"` // API响应消息：成功时为"Success"，失败时为错误描述
		Content struct {
			Point struct {
				Lat string `json:"lat"` // 纬度（字符串格式，需要转换为float64）
				Lng string `json:"lng"` // 经度（字符串格式，需要转换为float64）
			} `json:"point"`
			Address       string `json:"address"` // 详细地址，如"北京市北京市东城区天安门"
			AddressDetail struct {
				City     string `json:"city"`     // 城市，如"北京市"
				Province string `json:"province"` // 省份，如"北京市"
				District string `json:"district"` // 区县，如"东城区"
			} `json:"address_detail"`
		} `json:"content"`
	}

	// 使用json.Unmarshal将JSON字节数组解析为Go结构体
	if err := json.Unmarshal(body, &baiduResp); err != nil {
		return &pb.GetCurrentLocationResponse{
			Code:    10000,
			Message: fmt.Sprintf("解析JSON失败: %v", err),
		}, nil
	}

	// 检查API响应状态
	if baiduResp.Status != 0 {
		return &pb.GetCurrentLocationResponse{
			Code:    10000,
			Message: fmt.Sprintf("百度地图API错误: %s", baiduResp.Message),
		}, nil
	}

	// 解析经纬度
	// 将字符串格式的经纬度转换为float64类型，便于后续计算和存储
	latitude, err := strconv.ParseFloat(baiduResp.Content.Point.Lat, 64)
	if err != nil {
		return &pb.GetCurrentLocationResponse{
			Code:    10000,
			Message: "解析纬度失败",
		}, nil
	}
	longitude, err := strconv.ParseFloat(baiduResp.Content.Point.Lng, 64)
	if err != nil {
		return &pb.GetCurrentLocationResponse{
			Code:    10000,
			Message: "解析经度失败",
		}, nil
	}

	// 存储到Redis GEO数据结构
	geoKey := "china:users:location" // Redis键名，用于存储所有用户的地理位置
	geoName := req.UserId            // 用户ID作为位置名称，便于后续查询和更新

	// 使用GEOADD命令将用户位置添加到Redis GEO数据结构
	err = s.data.Rdb.GeoAdd(ctx, geoKey, &redis.GeoLocation{
		Name:      geoName,   // 位置名称（用户ID）
		Longitude: longitude, // 经度
		Latitude:  latitude,  // 纬度
	}).Err()
	if err != nil {
		// 只打日志，不影响主流程
		// 原因：位置获取是主要功能，Redis存储是辅助功能
		// 即使Redis存储失败，用户仍能获得位置信息
		fmt.Printf("存储用户地理位置到Redis失败: %v\n", err)
	}

	//  返回位置信息
	// 返回完整的用户位置信息，包括经纬度坐标和详细地址
	return &pb.GetCurrentLocationResponse{
		Latitude:  latitude,                                 // 纬度
		Longitude: longitude,                                // 经度
		Address:   baiduResp.Content.Address,                // 详细地址
		City:      baiduResp.Content.AddressDetail.City,     // 城市
		Province:  baiduResp.Content.AddressDetail.Province, // 省份
		District:  baiduResp.Content.AddressDetail.District, // 区县
		Code:      200,                                      // 成功状态码
		Message:   "获取位置成功",                                 // 成功消息
	}, nil
}

// 搜索记录

func (s *UserService) UserSearchHistory(ctx context.Context, req *pb.UserSearchRequest) (*pb.UserSearchResponse, error) {
	var user model.User
	err := s.data.DB.Where("id = ?", req.UserId).First(&user).Error
	if err != nil {
		return &pb.UserSearchResponse{
			Code:    10000,
			Message: "用户不存在",
		}, nil
	}

	// 限制搜索记录最多10条
	var count int64
	s.data.DB.Model(&model.UserSearchHistory{}).Where("user_id = ?", req.UserId).Count(&count)
	if count >= 10 {
		s.data.DB.Where("user_id = ?", req.UserId).Order("search_time asc").Limit(1).Delete(&model.UserSearchHistory{})
	}

	//去重：如果用户上一次搜索的关键词和本次相同，则只更新时间，不插入新记录
	var lastRecord model.UserSearchHistory
	err = s.data.DB.Where("user_id = ?", req.UserId).Order("search_time desc").First(&lastRecord).Error
	if err == nil && lastRecord.Keyword == req.Keyword {
		// 更新时间
		s.data.DB.Model(&model.UserSearchHistory{}).Where("id = ?", lastRecord.Id).Update("search_time", time.Now())
	} else {
		// 插入新记录
		s.data.DB.Create(&model.UserSearchHistory{
			UserId:     uint(req.UserId),
			Keyword:    req.Keyword,
			SearchTime: time.Now(),
		})
	}

	return &pb.UserSearchResponse{
		Code:    200,
		Message: "搜索成功",
	}, nil
}

// 创建行程接口
func (s *UserService) CreateTrip(ctx context.Context, req *pb.CreateTripRequest) (*pb.CreateTripResponse, error) {
	//根据 user_id 查询该用户是否存在
	var user model.User
	err := s.data.DB.Where("id = ?", req.UserId).First(&user).Error
	if err != nil {
		return &pb.CreateTripResponse{
			Code:    10000,
			Message: "用户不存在",
		}, nil
	}

	//验证起点地址不为空
	if req.StartLocation == "" {
		return &pb.CreateTripResponse{
			Code:    10000,
			Message: "起点地址不能为空",
		}, nil
	}

	//验证终点地址不为空
	if req.EndLocation == "" {
		return &pb.CreateTripResponse{
			Code:    10000,
			Message: "终点地址不能为空",
		}, nil
	}

	//验证车型不为空
	if req.CarType == "" {
		return &pb.CreateTripResponse{
			Code:    10000,
			Message: "车型不能为空",
		}, nil
	}

	//车型校验与单价获取
	var pricePerKm float64
	switch req.CarType {
	case "economy":
		pricePerKm = 2.0
	case "comfort":
		pricePerKm = 5.0
	case "luxury":
		pricePerKm = 10.0
	default:
		return &pb.CreateTripResponse{
			Code:    10000,
			Message: "不支持的车型",
		}, nil
	}

	//计算距离
	ak := pkg.BaiduMapAPIKey
	distance, err := pkg.EstimateDistanceByBaidu(req.StartLocation, req.EndLocation, ak)
	if err != nil {
		return &pb.CreateTripResponse{
			Code:    10000,
			Message: "距离计算失败",
		}, nil
	}

	// 计价逻辑
	baseFee := 17.0 // 起步价
	price := baseFee + distance*pricePerKm

	//生成行程ID
	tripID := fmt.Sprintf("TRIP%v", time.Now().UnixNano())

	return &pb.CreateTripResponse{
		TripId:  tripID,
		Price:   price,
		CarType: req.CarType,
		Message: "行程创建成功",
		Code:    200,
	}, nil
}

// 用户叫车
// 用户在叫车时，如果有未支付的订单，平台会要求先支付上一个订单才可以继续操作
// 叫车时平台根据 passenger_id 查看是哪个用户，然后为该用户进行叫车
// 所以我们需要根据 passenger_id 查看用户是否存在，还要有出发地址才能叫车，如果有未支付订单也不能叫车

func (s *UserService) CreateUserOrder(ctx context.Context, req *pb.CreateUserOrderRequest) (*pb.CreateUserOrderResponse, error) {
	// 根据 passenger_id 查询该用户是否存在
	var user model.User
	err := s.data.DB.Where("id = ?", req.PassengerId).First(&user).Error
	if err != nil {
		return &pb.CreateUserOrderResponse{
			Code:    10000,
			Message: "乘客不存在",
		}, nil
	}

	// 用户叫车时的出发地址不能为空
	if req.StartAddr == "" {
		return &pb.CreateUserOrderResponse{
			Code:    10000,
			Message: "出发地址不能为空",
		}, nil
	}

	//查询用户上个订单是否付款
	var existingOrder model.LxhOrder
	err = s.data.DB.Where("passenger_id = ? AND order_status IN ('待接单', '已接单', '行程中')", req.PassengerId).First(&existingOrder).Error
	if err == nil {
		return &pb.CreateUserOrderResponse{
			Code:    10000,
			Message: "请先支付上一个订单",
		}, nil
	}

	//获取起点经纬度
	ak := pkg.BaiduMapAPIKey
	startLat, startLng, err := pkg.GetLatLng(req.StartAddr, ak)
	if err != nil {
		return &pb.CreateUserOrderResponse{
			Code:    10000,
			Message: "起点地址无法解析经纬度，无法分配司机",
		}, nil
	}

	//分配最近司机，支持自动扩大范围
	radiusList := []float64{3, 5, 10} // 单位：公里
	var nearestDriverId string
	var nearestDriver model.LxhDriver
	for _, radius := range radiusList {
		drivers, err := s.data.Rdb.GeoRadius(ctx, "china:drivers:location", startLng, startLat, &redis.GeoRadiusQuery{
			Radius: radius,
			Unit:   "km",
			Sort:   "ASC",
			Count:  10,
		}).Result()
		if err != nil || len(drivers) == 0 {
			continue
		}
		for _, d := range drivers {
			// 查数据库确认司机可用
			var driver model.LxhDriver
			err := s.data.DB.Where("id = ? AND status = ?", d.Name, 1).First(&driver).Error
			if err == nil {
				nearestDriverId = d.Name
				nearestDriver = driver
				break
			}
		}
		if nearestDriverId != "" {
			break
		}
	}
	if nearestDriverId == "" {
		return &pb.CreateUserOrderResponse{
			Code:    10000,
			Message: "附近暂无可用司机，请稍后再试",
		}, nil
	}

	// 创建订单，分配司机
	order := &model.LxhOrder{
		OrderCode:   req.OrderCode,
		Amount:      req.Amount,
		OrderStatus: req.OrderStatus,
		PassengerId: int(req.PassengerId),
		StartAddr:   req.StartAddr,
		EndEnd:      req.EndEnd,
		Driver:      int(nearestDriver.Id),
		StartTime:   model.ParseTimePtr(req.StartTime),
		EndTime:     model.ParseTimePtr(req.EndTime),
		PayStatus:   req.PayStatus,
		PayType:     req.PayType,
		OrderType:   req.OrderType,
	}
	err = s.data.DB.Create(&order).Error
	if err != nil {
		return &pb.CreateUserOrderResponse{
			Code:    10000,
			Message: "创建订单失败",
		}, nil
	}

	return &pb.CreateUserOrderResponse{
		Code:    200,
		Message: "创建订单成功",
	}, nil
}

// 评价司机
// 评论肯定是订单结束之后评论的，订单结束后我们可以查找订单，通过订单找到该司机，去进行评论
// 所以我们需要去判断该用户、司机、订单是否存在，然后在评论时检查是否有敏感词

func (s *UserService) CommentDriver(ctx context.Context, req *pb.UserCommentDriverRequest) (*pb.UserCommentDriverResponse, error) {
	// 根据 passenger_id 查询该用户是否存在
	var user model.User
	err := s.data.DB.Where("id = ?", req.PassengerId).First(&user).Error
	if err != nil {
		return &pb.UserCommentDriverResponse{
			Code:    10000,
			Message: "该用户不存在",
		}, nil
	}

	//根据 driver_id 查询该司机是否存在
	var driver model.LxhDriver
	err = s.data.DB.Where("id = ?", req.DriverId).First(&driver).Error
	if err != nil {
		return &pb.UserCommentDriverResponse{
			Code:    10000,
			Message: "该司机不存在",
		}, nil
	}

	//根据 order_id 查询该订单是否存在
	var order model.LxhOrder
	err = s.data.DB.Where("id = ?", req.OrderId).First(&order).Error
	if err != nil {
		return &pb.UserCommentDriverResponse{
			Code:    10000,
			Message: "该订单不存在",
		}, nil
	}

	//评论时检查是否有敏感词
	if !pkg.TextCensor(req.Content) {
		return &pb.UserCommentDriverResponse{
			Code:    10000,
			Message: "该评论含有敏感词",
		}, nil
	}

	//评论入库
	com := &model.DriverEvaluation{
		OrderId:     uint(req.OrderId),
		PassengerId: uint(req.PassengerId),
		DriverId:    uint(req.DriverId),
		Score:       uint(req.Score),
		Content:     req.Content,
	}

	err = s.data.DB.Create(&com).Error
	if err != nil {
		return &pb.UserCommentDriverResponse{
			Code:    10000,
			Message: "评价失败",
		}, nil
	}

	return &pb.UserCommentDriverResponse{
		Code:    200,
		Message: "评价成功",
	}, nil

}
