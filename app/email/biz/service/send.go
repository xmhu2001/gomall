package service

import (
	"context"

	email "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/email"
)

type SendService struct {
	ctx context.Context
} // NewSendService new SendService
func NewSendService(ctx context.Context) *SendService {
	return &SendService{ctx: ctx}
}

// Run create note info
func (s *SendService) Run(req *email.EmailReq) (resp *email.EmailResp, err error) {

	return
}
