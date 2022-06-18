package logic

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"cyntra/product/rpc/internal/svc"
	"cyntra/product/rpc/model"
	"cyntra/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	res, err := l.svcCtx.ProductModel.FindOneByName(l.ctx, in.Name)

	if err != nil && !errors.Is(err, sqlc.ErrNotFound) {
		l.Logger.Error(err)
		return nil, status.Error(codes.Unknown, codes.Unknown.String())
	}

	if res != nil {
		l.Logger.Error("Product with this name already Exists")
		return nil, status.Errorf(codes.AlreadyExists, "Product %v Already Exists", in.Name)
	}

	newProduct := model.Product{
		Name:            in.Name,
		Price:           int64(in.Price),
		Description:     in.Description,
		Size:            in.Size,
		Color:           in.Color,
		Brand:           in.Brand,
		Shade:           in.Shade,
		WashCare:        sql.NullString{String: *in.WashCare, Valid: in.WashCare != nil},
		Stretchable:     sql.NullString{String: *in.Stretchable, Valid: in.Stretchable != nil},
		Features:        sql.NullString{String: *in.Features, Valid: in.Features != nil},
		Fade:            sql.NullString{String: *in.Fade, Valid: in.Fade != nil},
		Fabric:          in.Fabric,
		Category:        in.Category,
		Countryoforigin: in.CountryOfOrigin,
		Discount:        int64(in.Discount),
		Quantity:        int64(in.Quantity),
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		ArchivedAt:      sql.NullTime{Time: time.Now(), Valid: false},
	}

	_, err = l.svcCtx.ProductModel.Insert(l.ctx, &newProduct)

	if err != nil {
		l.Logger.Error(err)
		return nil, status.Error(codes.Unknown, codes.Unknown.String())
	}

	return &product.ProductResponse{
		Name:            newProduct.Name,
		Price:           uint64(newProduct.Price),
		Description:     newProduct.Description,
		Size:            newProduct.Size,
		Color:           newProduct.Color,
		Brand:           newProduct.Brand,
		Shade:           newProduct.Shade,
		WashCare:        in.WashCare,
		Stretchable:     in.Stretchable,
		Features:        in.Features,
		Fade:            in.Fade,
		Fabric:          newProduct.Fabric,
		Category:        newProduct.Category,
		CountryOfOrigin: newProduct.Countryoforigin,
		Discount:        uint32(in.Discount),
		Quantity:        uint64(in.Quantity),
		CreatedAt:       timestamppb.New(newProduct.CreatedAt),
		UpdatedAt:       timestamppb.New(newProduct.CreatedAt),
		ArchivedAt:      nil,
	}, nil
}
