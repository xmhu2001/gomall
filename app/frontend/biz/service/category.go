package service

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	category "github.com/xmhu2001/gomall/app/frontend/hertz_gen/frontend/category"
	"github.com/xmhu2001/gomall/app/frontend/infra/rpc"
	rpcproduct "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/product"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	r, err := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{CategoryName: req.Category})
	if err != nil {
		return nil, err
	}

	return utils.H{
		"title": "Category",
		"items": r.Products,
	}, nil
}
