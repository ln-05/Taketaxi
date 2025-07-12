package pkg

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

// 百度地图API配置
const (
	BaiduMapAPIKey = "N3RoKOXTLUJgX6bchYUFqntOtcwTwD7u" //
	BaiduMapAPIURL = "https://api.map.baidu.com"
)

// LocationInfo 位置信息结构体
type LocationInfo struct {
	Latitude     float64 `json:"latitude"`      // 纬度
	Longitude    float64 `json:"longitude"`     // 经度
	Address      string  `json:"address"`       // 详细地址
	City         string  `json:"city"`          // 城市
	District     string  `json:"district"`      // 区县
	Street       string  `json:"street"`        // 街道
	StreetNumber string  `json:"street_number"` // 门牌号
}

// BaiduMapResponse 百度地图API响应结构体
type BaiduMapResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  struct {
		Location struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"location"`
		FormattedAddress string `json:"formatted_address"`
		AddressComponent struct {
			City         string `json:"city"`
			District     string `json:"district"`
			Street       string `json:"street"`
			StreetNumber string `json:"street_number"`
		} `json:"addressComponent"`
	} `json:"result"`
}

// CoordinateConvertResponse 坐标转换响应结构体
type CoordinateConvertResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Result  []struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
	} `json:"result"`
}

// BaiduMapService 百度地图服务
type BaiduMapService struct {
	apiKey string
	client *http.Client
}

// 计算起点和终点的距离（演示用，实际可接入地图API）
func EstimateDistance(start, end string) float64 {
	rand.Seed(time.Now().UnixNano())
	return 5 + rand.Float64()*15 // 5~20公里
}

// 地址转经纬度
func GetLatLng(address, ak string) (lat, lng float64, err error) {
	api := fmt.Sprintf("http://api.map.baidu.com/geocoding/v3/?address=%s&output=json&ak=%s", url.QueryEscape(address), ak)
	resp, err := http.Get(api)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	var result struct {
		Status int `json:"status"`
		Result struct {
			Location struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"location"`
		} `json:"result"`
	}
	body, _ := io.ReadAll(resp.Body)
	err = json.Unmarshal(body, &result)
	if err != nil {
		return
	}
	if result.Status != 0 {
		err = fmt.Errorf("百度地图地理编码失败")
		return
	}
	lat = result.Result.Location.Lat
	lng = result.Result.Location.Lng
	return
}

// 计算真实驾车距离（单位：米）
func EstimateDistanceByBaidu(startAddr, endAddr, ak string) (float64, error) {
	// 1. 地址转经纬度
	startLat, startLng, err := GetLatLng(startAddr, ak)
	if err != nil {
		return 0, fmt.Errorf("起点地址转经纬度失败: %v", err)
	}
	endLat, endLng, err := GetLatLng(endAddr, ak)
	if err != nil {
		return 0, fmt.Errorf("终点地址转经纬度失败: %v", err)
	}

	// 2. 路线规划API
	api := fmt.Sprintf(
		"http://api.map.baidu.com/directionlite/v1/driving?origin=%f,%f&destination=%f,%f&ak=%s",
		startLat, startLng, endLat, endLng, ak,
	)
	resp, err := http.Get(api)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	var result struct {
		Status int `json:"status"`
		Result struct {
			Routes []struct {
				Distance float64 `json:"distance"`
			} `json:"routes"`
		} `json:"result"`
	}
	body, _ := io.ReadAll(resp.Body)
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, err
	}
	if result.Status != 0 || len(result.Result.Routes) == 0 {
		return 0, fmt.Errorf("百度地图路线规划失败")
	}
	// 返回米数
	return result.Result.Routes[0].Distance, nil
}

// CalcDistance 计算两点间球面距离（单位：米）
func CalcDistance(lat1, lng1, lat2, lng2 float64) float64 {
	const earthRadius = 6371000.0 // 地球半径（米）
	rad := func(d float64) float64 { return d * math.Pi / 180.0 }
	lat1Rad, lat2Rad := rad(lat1), rad(lat2)
	deltaLat := lat1Rad - lat2Rad
	deltaLng := rad(lng1) - rad(lng2)
	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) + math.Cos(lat1Rad)*math.Cos(lat2Rad)*math.Sin(deltaLng/2)*math.Sin(deltaLng/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return earthRadius * c
}

// CheckServiceScore 校验服务分是否达标
func CheckServiceScore(score, threshold float64) bool {
	return score >= threshold
}
