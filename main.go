package main

import (
	_ "go_verification_code/web"
	"net/http"
)

func main() {
	http.ListenAndServe(":80", nil)
}
