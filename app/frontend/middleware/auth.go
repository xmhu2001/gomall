package middleware

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/hertz-contrib/sessions"
	frontendUtils "github.com/xmhu2001/gomall/app/frontend/utils"
)

func GlobalAuth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, s.Get("user_id"))
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		userID := s.Get("user_id")
		if userID == nil {
			c.Redirect(consts.StatusOK, []byte("/sign-in?next="+c.FullPath()))
			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
