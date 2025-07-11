package pkg

import (
	"fmt"
	"github.com/smartwalle/alipay/v3"
)

func AliPay(Subject, OutTradeNo, TotalAmount, ProductCode string) string {
	privateKey := "MIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQCMm8PA1NAyQbU6EF958YMLHWuEqbMsGwJOh4AkzVyen2FOvPpIs+YHdNV9iK89jTuBSv/DcnTsnDy7+SLx97aRYrX+Hxec23DA+OCkcIU2YfefW9XkINLdvavFi+eutwAuS4Mq2CcBSS04HPNJwUFJLtd4QLBoewTUAep4l4iCeSG6QDiF7wQoENvEj3K7AcUdnz5eee3De4hfs6vo80CLmxOt9XOgHs/j04EImtidF8NE24PqdyLW1/oNMS88cdB/2aOUX+rSVThpUo91zUd8S64Rj+XIQPL0gdbdEJjCCGUIFneHzmafj2jPsaXurOe3jqEZgLhPpuD03273l/CHAgMBAAECggEAOvOPlghjpb6A0fBNZ14HmCBklMFleod0ZyJOap+jRmoCLTX0JYkvV073t2MYPdE58pJUO0iAlzdGC3V/3j0CUWA2d23nw38Jfg6aeLcoUs2uOSUSv0u3vPopy3BF2a2mKMLGIq9vYw2gJNeWdQTOb3VxJ1zP2W5CbvTCrn6x4CltO3m2TZ14EtKGR0Gb1QqG4axxMlgoasRhZviow6wfz+5ozYLyH2m6vqed5xGcICnVmCSdYmr8M8Y2T2715j0br8/HPskGpRrTUPEAPmimC4WRsI6b6dZPLxBkMTXUwDdbmUlEBkBcXxkB8F36RPINXnaXW8fpd6n59Z4ZWj2sAQKBgQD+ESuL+wZZNExKB4nmFcNpNne9VXLErss7hBg9DL7ctzEjpHooq7VcxQxZ+jP4sOA+z8GCCOG/KL+NJ3Dl9imRU9C/NOTr7ryEt97oF1O3TpIJX7OpJA6gkX3nnvAg6bH5y+NjjR7EhE/8z9uKLag3YK3yb0jmkZMDSObJrXDKhwKBgQCNrZ5cZ7FzE37eZlRB8it0rLc/Ai1AB2xeh1g9/XLFc5c4XV+IKc7v4QgWByYHAlhHH+WRJ+nLaDvKnoom69yea1RNamNTpyAjNYjF/8dwVPiysmXcivumcLJOk6VyyRp8L5XyJ3Wcdye/URb1Eiy4dXmEogtsCxicmNye6esqAQKBgQDRI9kw9YaYT9Cs79+4Ixoktc2DcZ90AF8Tsv0w5BkJH1O3/2D/sbktkJdGSgwWH6O4kNS98gnCjgyAaYMqCaTWGKSxgR0gifltVt6LNwiW0HqudLUz+pZ/3sRBsjgBH959vjSEclptcm1VstUJvePHEjKearUb96GJBq3UPo2maQKBgCmhzRkmN1Szc+Jye2vtd6uPXUGcqCNaixOz/dApe+JdtyAPABWJA5AltjSS4S/KFpq/2ruu2Nq2xGsj70DEqk27CEi/bHSWdbEi+BAGu1pc/b+1AI2wIYalMyA37rQuOCQp74v64Gqb8QgdEfvt3UpO6uh4nZagJBlQKfRglIQBAoGBAMcC1/Ktg7359C2sf5N01vCS9Cay3PPpTfMMkgOdJc+BmDMd9RpsQ5nNrxDxzDvZpyr3RSdUE2U6Ca6zXTPCcoBF+akpuZHJgMgvmln8hioicqLwiFeDq55S1ZMafBMYpuzB4XsoS6XkDsMhRMgQjtQriOOwS9TAaGwIQmupQDsh" // 必须，上一步中使用 RSA签名验签工具 生成的私钥
	client, err := alipay.New("9021000143641568", privateKey, false)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	var p = alipay.TradeWapPay{}
	p.NotifyURL = "http://xxx"
	p.ReturnURL = "http://www.baidu.com"
	p.Subject = Subject
	p.OutTradeNo = OutTradeNo
	p.TotalAmount = TotalAmount
	p.ProductCode = ProductCode

	url, err := client.TradeWapPay(p)
	if err != nil {
		fmt.Println(err)
	}

	// 这个 payURL 即是用于打开支付宝支付页面的 URL，可将输出的内容复制，到浏览器中访问该 URL 即可打开支付页面。
	var payURL = url.String()
	fmt.Println(payURL)
	return payURL
}
