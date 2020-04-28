package controller

import (
	"net/http"
)

type Index interface {
	//app_http.Controller
}

//定义控制器
func Initialize(base http.Handler) http.Handler {

	base = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	return base
}
