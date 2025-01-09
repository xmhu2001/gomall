package auth

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/xmhu2001/gomall/app/frontend/biz/service"
	auth "github.com/xmhu2001/gomall/app/frontend/hertz_gen/frontend/auth"
	common "github.com/xmhu2001/gomall/app/frontend/hertz_gen/frontend/common"
	"github.coom/xmhu2001/gomall/app/frontend/biz/service"
	"github.coom/xmhu2001/gomall/app/frontend/biz/utils"
	auth "github.coom/xmhu2001/gomall/app/frontend/hertz_gen/frontend/auth"
)

// Login .
// @router /auth/login [POST]
func Login(ctx context.Context, c *app.RequestContext) {
	var err error
	var req auth.LoginReq
	err = c.BindAndValidate(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	// resp := &common.Empty{}
	_, err = service.NewLoginService(ctx, c).Run(&req)
	if err != nil {
		utils.SendErrResponse(ctx, c, consts.StatusOK, err)
		return
	}

	utils.SendSuccessResponse(ctx, c, consts.StatusOK, "done")
}
