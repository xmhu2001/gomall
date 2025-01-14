package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
	auth "github.com/xmhu2001/gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/xmhu2001/gomall/app/frontend/hertz_gen/frontend/common"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (resp *common.Empty, err error) {
	session := sessions.Default(h.RequestContext)
	session.Set("user_id", 1)
	err = session.Save()
	if err != nil {
		return nil, err
	}
	return
}
