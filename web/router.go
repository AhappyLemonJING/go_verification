package web

import "net/http"

func init() {
	http.HandleFunc("/index", IndexHandler)
	http.HandleFunc("/captcha", CaptchaHandler)
}
