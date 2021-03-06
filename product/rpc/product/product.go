// Code generated by goctl. DO NOT EDIT!
// Source: product.proto

package product

import (
	"context"

	"cyntra/product/rpc/types/product"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CreateProductRequest     = product.CreateProductRequest
	DeleteProductByIdRequest = product.DeleteProductByIdRequest
	GetProductByIdRequest    = product.GetProductByIdRequest
	ProductResponse          = product.ProductResponse
	UpdateProductRequest     = product.UpdateProductRequest

	Product interface {
		GetProduct(ctx context.Context, in *GetProductByIdRequest, opts ...grpc.CallOption) (*ProductResponse, error)
		CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
		UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error)
		DeleteProduct(ctx context.Context, in *DeleteProductByIdRequest, opts ...grpc.CallOption) (*ProductResponse, error)
	}

	defaultProduct struct {
		cli zrpc.Client
	}
)

func NewProduct(cli zrpc.Client) Product {
	return &defaultProduct{
		cli: cli,
	}
}

func (m *defaultProduct) GetProduct(ctx context.Context, in *GetProductByIdRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.GetProduct(ctx, in, opts...)
}

func (m *defaultProduct) CreateProduct(ctx context.Context, in *CreateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.CreateProduct(ctx, in, opts...)
}

func (m *defaultProduct) UpdateProduct(ctx context.Context, in *UpdateProductRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.UpdateProduct(ctx, in, opts...)
}

func (m *defaultProduct) DeleteProduct(ctx context.Context, in *DeleteProductByIdRequest, opts ...grpc.CallOption) (*ProductResponse, error) {
	client := product.NewProductClient(m.cli.Conn())
	return client.DeleteProduct(ctx, in, opts...)
}
