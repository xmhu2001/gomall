package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/xmhu2001/gomall/app/product/biz/dal/mysql"
	"github.com/xmhu2001/gomall/app/product/biz/model"
	product "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/product"
)

type ListProductsService struct {
	ctx context.Context
} // NewListProductsService new ListProductsService
func NewListProductsService(ctx context.Context) *ListProductsService {
	return &ListProductsService{ctx: ctx}
}

// Run create note info
func (s *ListProductsService) Run(req *product.ListProductsReq) (resp *product.ListProductsResp, err error) {
	if req.CategoryName == "" {
		return nil, kerrors.NewGRPCBizStatusError(200401, "category name is required")
	}
	categoryQuery := model.NewCategoryQuery(s.ctx, mysql.DB)

	list, err := categoryQuery.GetProductsByCatrgoryName(req.CategoryName)
	if err != nil {
		return nil, err
	}

	resp = &product.ListProductsResp{}
	for _, v := range list {
		for _, p := range v.Products {
			resp.Products = append(resp.Products, &product.Product{
				Id:          uint32(p.ID),
				Name:        p.Name,
				Description: p.Description,
				Picture:     p.Picture,
				Price:       p.Price,
			})
		}
	}

	return resp, nil
}
