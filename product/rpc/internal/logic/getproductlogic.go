package logic

import (
	"context"

	"cyntra/product/rpc/internal/svc"
	"cyntra/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetProductLogic) GetProduct(in *product.GetProductByIdRequest) (*product.ProductResponse, error) {
	// todo: add your logic here and delete this line

	return &product.ProductResponse{}, nil
}