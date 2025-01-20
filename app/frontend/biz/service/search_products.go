package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	product "github.com/xmhu2001/gomall/app/frontend/hertz_gen/frontend/product"
	"github.com/xmhu2001/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/product"
)

type SearchProductsService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewSearchProductsService(Context context.Context, RequestContext *app.RequestContext) *SearchProductsService {
	return &SearchProductsService{RequestContext: RequestContext, Context: Context}
}

func (h *SearchProductsService) Run(req *product.SearchProductsReq) (resp map[string]any, err error) {
	r, err := rpc.ProductClient.SearchProducts(h.Context, &rpcproduct.SearchProductsReq{Query: req.Q})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"items": r.Results,
		"q":     req.Q,
	}, nil
}
