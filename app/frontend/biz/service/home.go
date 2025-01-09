package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	common "github.com/xmhu2001/gomall/app/frontend/hertz_gen/frontend/common"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (map[string]any, error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "manga-1", "Price": 100, "Picture": "/static/image/manga.png"},
		{"Name": "manga-2", "Price": 110, "Picture": "/static/image/manga.png"},
		{"Name": "manga-3", "Price": 120, "Picture": "/static/image/manga.png"},
		{"Name": "manga-4", "Price": 130, "Picture": "/static/image/manga.png"},
		{"Name": "manga-5", "Price": 140, "Picture": "/static/image/manga.png"},
		{"Name": "manga-6", "Price": 150, "Picture": "/static/image/manga.png"},
	}
	resp["Title"] = "Hot Sales"
	resp["Items"] = items

	return resp, nil
}
