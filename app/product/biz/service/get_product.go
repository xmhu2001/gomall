package service

import (
	"context"

	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/xmhu2001/gomall/app/product/biz/dal/mysql"
	"github.com/xmhu2001/gomall/app/product/biz/dal/redis"
	"github.com/xmhu2001/gomall/app/product/biz/model"
	product "github.com/xmhu2001/gomall/rpc_gen/kitex_gen/product"
)

type GetProductService struct {
	ctx context.Context
} // NewGetProductService new GetProductService
func NewGetProductService(ctx context.Context) *GetProductService {
	return &GetProductService{ctx: ctx}
}

// Run create note info
func (s *GetProductService) Run(req *product.GetProductReq) (resp *product.GetProductResp, err error) {
	if req.Id == 0 {
		return nil, kerrors.NewGRPCBizStatusError(200401, "product id is required")
	}
	// construct query
	productQuery := model.NewCachedProductQuery(s.ctx, mysql.DB, redis.RedisClient)

	p, err := productQuery.GetById(int(req.Id))

	if err != nil {
		return nil, err
	}

	return &product.GetProductResp{
		Product: &product.Product{
			Id:          uint32(p.ID),
			Picture:     p.Picture,
			Price:       p.Price,
			Name:        p.Name,
			Description: p.Description,
		},
	}, nil
}
