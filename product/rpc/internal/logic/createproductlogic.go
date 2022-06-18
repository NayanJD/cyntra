package logic

import (
	"context"

	"cyntra/product/rpc/internal/svc"
	"cyntra/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateProductLogic) CreateProduct(in *product.CreateProductRequest) (*product.ProductResponse, error) {
	// todo: add your logic here and delete this line

	return &product.ProductResponse{}, nil
}
