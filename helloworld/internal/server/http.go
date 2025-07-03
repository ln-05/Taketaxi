package server

import (
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/conf"
	"helloworld/internal/service"

	"crypto/sha1"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"time"

	khttp "github.com/go-kratos/kratos/v2/transport/http"
	"github.com/rs/cors"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/skip2/go-qrcode"
)

// 生成随机字符串
func RandString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.GreeterService, logger log.Logger) *khttp.Server {
	var opts = []khttp.ServerOption{
		khttp.Middleware(
			recovery.Recovery(),
		),
	}

	// 添加 CORS 跨域支持
	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // 生产环境建议指定域名
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
	opts = append(opts, khttp.Filter(func(next http.Handler) http.Handler {
		return corsHandler.Handler(next)
	}))

	if c.Http.Network != "" {
		opts = append(opts, khttp.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, khttp.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, khttp.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := khttp.NewServer(opts...)
	v1.RegisterGreeterHTTPServer(srv, greeter)

	// 微信服务器验证专用 handler，带 Token 校验
	const wechatToken = "111" // 必须和微信后台填写一致
	srv.HandleFunc("/wechat/callback", func(w http.ResponseWriter, r *http.Request) {
		signature := r.URL.Query().Get("signature")
		timestamp := r.URL.Query().Get("timestamp")
		nonce := r.URL.Query().Get("nonce")
		echostr := r.URL.Query().Get("echostr")

		arr := []string{wechatToken, timestamp, nonce}
		sort.Strings(arr)
		str := arr[0] + arr[1] + arr[2]
		h := sha1.New()
		h.Write([]byte(str))
		hash := fmt.Sprintf("%x", h.Sum(nil))

		if hash == signature {
			w.Write([]byte(echostr))
			return
		}
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("signature error"))
	})

	// 微信扫码登录二维码生成
	srv.HandleFunc("/wechat/login_qrcode", func(w http.ResponseWriter, r *http.Request) {
		path := "2053b42a.r23.cpolar.top" // TODO: 配置化
		state := RandString(5)
		redirectURI := "http://" + path + "/wechat/login_callback"
		redirectURL := url.QueryEscape(redirectURI)
		wechatLoginURL := fmt.Sprintf(
			"https://open.weixin.qq.com/connect/oauth2/authorize?appid=%s&redirect_uri=%s&response_type=code&state=%s&scope=snsapi_userinfo#wechat_redirect",
			"wxcdda396e027fbf56",
			redirectURL, state)
		w.Header().Set("Content-Type", "image/png")
		qrCode, err := qrcode.Encode(wechatLoginURL, qrcode.Medium, 256)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error generating QR code"))
			return
		}
		w.Write(qrCode)
	})

	// 微信扫码登录回调
	srv.HandleFunc("/wechat/login_callback", func(w http.ResponseWriter, r *http.Request) {
		appid := "wxcdda396e027fbf56"
		appsecret := "5ed0492fceaacf9bbb17fa32cef64577"
		code := r.URL.Query().Get("code")
		if appid == "" || appsecret == "" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("AppID或AppSecret不能为空，请检查配置文件和服务启动顺序"))
			return
		}
		// 获取access_token和openid
		tokenResp, err := http.Get(fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code", appid, appsecret, code))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error,获取token失败"))
			return
		}
		defer tokenResp.Body.Close()
		var tokenData struct {
			AccessToken  string `json:"access_token"`
			ExpiresIn    int    `json:"expires_in"`
			RefreshToken string `json:"refresh_token"`
			OpenID       string `json:"openid"`
			Scope        string `json:"scope"`
		}
		if err1 := json.NewDecoder(tokenResp.Body).Decode(&tokenData); err1 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("error,解析token失败"))
			return
		}
		userInfoURL := fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s", tokenData.AccessToken, tokenData.OpenID)
		userInfoResp, err := http.Get(userInfoURL)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("获取用户信息失败"))
			return
		}
		defer userInfoResp.Body.Close()
		var userData struct {
			OpenID     string `json:"openid"`
			Nickname   string `json:"nickname"`
			HeadImgUrl string `json:"headimgurl"`
		}
		if err1 := json.NewDecoder(userInfoResp.Body).Decode(&userData); err1 != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("解析用户信息失败"))
			return
		}
		// 返回openid、access_token、nickname等
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"openid":       userData.OpenID,
			"nickname":     userData.Nickname,
			"access_token": tokenData.AccessToken,
			"headimgurl":   userData.HeadImgUrl,
		})
	})

	return srv
}
