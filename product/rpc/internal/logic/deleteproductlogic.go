package logic

import (
	"context"

	"cyntra/product/rpc/internal/svc"
	"cyntra/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteProductLogic) DeleteProduct(in *product.DeleteProductByIdRequest) (*product.ProductResponse, error) {
	// todo: add your logic here and delete this line

	return &product.ProductResponse{}, nil
}
