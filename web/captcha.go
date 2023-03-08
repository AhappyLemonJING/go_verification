package web

import (
	"fmt"
	"go_verification_code/captcha"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
)

func CaptchaHandler(w http.ResponseWriter, req *http.Request) {
	c := captcha.NewCaptcha()
	request := &captcha.CaptchaRequest{}
	request.NeedGetCaptchaTime = 0
	request.Randstr = req.FormValue("rand_str")
	request.SceneId, _ = strconv.ParseUint(req.FormValue("scene_id"), 10, 64)
	request.BusinessId, _ = strconv.ParseUint(req.FormValue("business_id"), 10, 64)
	userIP := ClientPublicIP(req)
	if userIP != "" {
		request.UserIp = userIP
	} else {
		request.UserIp = "127.0.0.1"
	}
	request.Ticket = req.FormValue("ticket")
	request.Imei = req.FormValue("imei")
	request.MacAddress = req.FormValue("mac_address")
	fmt.Printf("#{request}")
	res, err := c.DescribeCaptchaResult(request)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(w, err.Error())
		return
	}
	fmt.Fprintf(w, res)

}

// ClientPublicIP 尽最大努力实现获取客户端公网 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientPublicIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" && !HasLocalIPddr(ip) {
			return ip
		}
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" && !HasLocalIPddr(ip) {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		if !HasLocalIPddr(ip) {
			return ip
		}
	}
	return ""
}

// HasLocalIPddr 检测 IP 地址字符串是否是内网地址
func HasLocalIPddr(ip string) bool {
	return HasLocalIP(net.ParseIP(ip))
}

// HasLocalIP 检测 IP 地址是否是内网地址
func HasLocalIP(ip net.IP) bool {
	if ip.IsLoopback() {
		return true
	}
	ip4 := ip.To4()
	if ip4 == nil {
		return false
	}
	return ip4[0] == 10 ||
		(ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31) ||
		(ip4[0] == 169 && ip4[1] == 254) ||
		(ip4[0] == 192 && ip4[1] == 168)
}
