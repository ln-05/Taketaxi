package pkg

import (
	"encoding/json"
	"fmt"
	"github.com/Baidu-AIP/golang-sdk/aip/censor"
)

type Censor struct {
	ConclusionType uint `json:"ConclusionType"`
}

func TextCensor(content string) bool {
	client := censor.NewClient("E1kxZW8WN2tbXu0KiVbCIQCK", "xNwHgQLCFSkC0u5cAIv4lpgSDM0orT6t")
	//如果是百度云ak sk,使用下面的客户端
	res := client.TextCensor(content)
	fmt.Println(res)
	var data Censor
	json.Unmarshal([]byte(res), &data)
	if data.ConclusionType != 1 {
		return false
	}
	return true
}
