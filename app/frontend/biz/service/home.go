package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"

	common "github.com/xmhu2001/gomall/app/frontend/hertz_gen/frontend/common"
	"github.com/xmhu2001/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/product"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *common.Empty) (res map[string]any, err error) {
	products, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{})
	if err != nil {
		return nil, err
	}
	
	return utils.H{
		"title": "Hot Sale",
		"items": products.Products,
	}, nil
}
