package middleware

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/pkg/endpoint"
	"time"
)

func Middleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) error {
		begin := time.Now()
		err := next(ctx, req, resp)
		fmt.Printf("middleware took %v\n", time.Since(begin))
		return err
	}
}
