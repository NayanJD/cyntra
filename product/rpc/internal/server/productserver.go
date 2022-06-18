// Code generated by goctl. DO NOT EDIT!
// Source: product.proto

package server

import (
	"context"

	"cyntra/product/rpc/internal/logic"
	"cyntra/product/rpc/internal/svc"
	"cyntra/product/rpc/types/product"
)

type ProductServer struct {
	svcCtx *svc.ServiceContext
	product.UnimplementedProductServer
}

func NewProductServer(svcCtx *svc.ServiceContext) *ProductServer {
	return &ProductServer{
		svcCtx: svcCtx,
	}
}

func (s *ProductServer) GetProduct(ctx context.Context, in *product.GetProductByIdRequest) (*product.ProductResponse, error) {
	l := logic.NewGetProductLogic(ctx, s.svcCtx)
	return l.GetProduct(in)
}

func (s *ProductServer) CreateProduct(ctx context.Context, in *product.CreateProductRequest) (*product.ProductResponse, error) {
	l := logic.NewCreateProductLogic(ctx, s.svcCtx)
	return l.CreateProduct(in)
}

func (s *ProductServer) UpdateProduct(ctx context.Context, in *product.UpdateProductRequest) (*product.ProductResponse, error) {
	l := logic.NewUpdateProductLogic(ctx, s.svcCtx)
	return l.UpdateProduct(in)
}

func (s *ProductServer) DeleteProduct(ctx context.Context, in *product.DeleteProductByIdRequest) (*product.ProductResponse, error) {
	l := logic.NewDeleteProductLogic(ctx, s.svcCtx)
	return l.DeleteProduct(in)
}