// Code generated by hertz generator.

package cart

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/xmhu2001/gomall/app/frontend/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return []app.HandlerFunc{middleware.Auth()}
}

func _addcartitemMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _getcartMw() []app.HandlerFunc {
	// your code...
	return nil
}
