package middleware

import "github.com/cloudwego/hertz/pkg/app/server"

// main初始化时注册中间件
func Register(h *server.Hertz) {
	h.Use(GlobalAuth())
}
