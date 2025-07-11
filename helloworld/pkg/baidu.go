package pkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Baidu-AIP/golang-sdk/aip/censor"
)

type TextCensorResp struct {
	LogId          uint   `json:"log_id"`
	ErrorCode      uint   `json:"error_code"`
	ErrorMsg       string `json:"error_msg"`
	Conclusion     string `json:"conclusion"`
	ConclusionType uint   `json:"conclusionType"`
}

func TextCensor(content string) bool {
	client := censor.NewClient("4cbtVOEaoIWhZWKIFQdmS9gE", "tix1wSQIzBBkfGp7eqciA8RehHX84Rka")
	//如果是百度云ak sk,使用下面的客户端

	res := client.TextCensor(content)
	var resp TextCensorResp
	json.Unmarshal([]byte(res), &resp)
	if resp.ConclusionType == 1 {
		return true
	}
	return false

}

// 新增：通过经纬度获取地理位置
func GetLocationByLatLng(lat, lng float64) (address, province, city, district string, err error) {
	// 百度地图API Key
	ak := "ODwx6VDM3kts8QIsalQt0btSCEIBte20"
	url := fmt.Sprintf("http://api.map.baidu.com/reverse_geocoding/v3/?ak=%s&output=json&coordtype=wgs84ll&location=%f,%f", ak, lat, lng)
	resp, err := http.Get(url)
	if err != nil {
		return "", "", "", "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", "", "", "", err
	}
	type resultStruct struct {
		Result struct {
			FormattedAddress string `json:"formatted_address"`
			AddressComponent struct {
				Province string `json:"province"`
				City     string `json:"city"`
				District string `json:"district"`
			} `json:"addressComponent"`
		} `json:"result"`
		Status int    `json:"status"`
		Msg    string `json:"msg"`
	}
	var res resultStruct
	err = json.Unmarshal(body, &res)
	if err != nil {
		return "", "", "", "", err
	}
	if res.Status != 0 {
		return "", "", "", "", fmt.Errorf("baidu map api error: %s", res.Msg)
	}
	return res.Result.FormattedAddress, res.Result.AddressComponent.Province, res.Result.AddressComponent.City, res.Result.AddressComponent.District, nil
}

// 百度地图API：根据经纬度获取驾车路线距离（单位：米）
func GetBaiduRouteDistance(origin, destination, ak string) (int, error) {
	url := fmt.Sprintf("http://api.map.baidu.com/directionlite/v1/driving?origin=%s&destination=%s&ak=%s", origin, destination, ak)
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	var result struct {
		Status int `json:"status"`
		Result struct {
			Routes []struct {
				Distance int `json:"distance"`
			} `json:"routes"`
		} `json:"result"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}
	if result.Status != 0 || len(result.Result.Routes) == 0 {
		return 0, fmt.Errorf("baidu api error or no route found")
	}
	return result.Result.Routes[0].Distance, nil
}
