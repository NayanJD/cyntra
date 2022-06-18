package logic

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"cyntra/product/rpc/internal/svc"
	"cyntra/product/rpc/model"
	"cyntra/product/rpc/types/product"

	"github.com/google/uuid"
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

	var washCare, stretchable, features, fade string

	if in.WashCare != nil {
		washCare = *in.WashCare
	} else {
		washCare = ""
	}

	if in.Stretchable != nil {
		stretchable = *in.Stretchable
	} else {
		stretchable = ""
	}

	if in.Features != nil {
		features = *in.Features
	} else {
		features = ""
	}

	if in.Fade != nil {
		fade = *in.Fade
	} else {
		fade = ""
	}

	newId := uuid.New()
	newProduct := model.Product{
		Id:              newId.String(),
		Name:            in.Name,
		Price:           int64(in.Price),
		Description:     in.Description,
		Size:            in.Size,
		Color:           in.Color,
		Brand:           in.Brand,
		Shade:           in.Shade,
		WashCare:        sql.NullString{String: washCare, Valid: in.WashCare != nil},
		Stretchable:     sql.NullString{String: stretchable, Valid: in.Stretchable != nil},
		Features:        sql.NullString{String: features, Valid: in.Features != nil},
		Fade:            sql.NullString{String: fade, Valid: in.Fade != nil},
		Fabric:          in.Fabric,
		Category:        in.Category,
		CountryOfOrigin: in.CountryOfOrigin,
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
		Id:              newProduct.Id,
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
		CountryOfOrigin: newProduct.CountryOfOrigin,
		Discount:        uint32(in.Discount),
		Quantity:        uint64(in.Quantity),
		CreatedAt:       timestamppb.New(newProduct.CreatedAt),
		UpdatedAt:       timestamppb.New(newProduct.UpdatedAt),
		ArchivedAt:      nil,
	}, nil
}
