package logic

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"cyntra/product/rpc/internal/svc"
	"cyntra/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateProductLogic) UpdateProduct(in *product.UpdateProductRequest) (*product.ProductResponse, error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.ProductModel.FindOne(l.ctx, in.Id)

	if err != nil {

		if errors.Is(err, sqlc.ErrNotFound) {
			l.Logger.Error("Product does not exist")
			return nil, status.Errorf(codes.NotFound, "Product %v does not exist", in.Id)
		} else {
			l.Logger.Error(err)
			return nil, status.Error(codes.Unknown, codes.Unknown.String())
		}
	}

	if in.Name != nil {
		res.Name = *in.Name
	}

	if in.Price != nil {
		res.Price = int64(*in.Price)
	}

	if in.Description != nil {
		res.Description = *in.Description
	}

	if in.Size != nil {
		res.Size = *in.Size
	}

	if in.Color != nil {
		res.Color = *in.Color
	}

	if in.Brand != nil {
		res.Brand = *in.Brand
	}

	if in.Shade != nil {
		res.Shade = *in.Shade
	}

	if in.WashCare != nil {
		res.WashCare = sql.NullString{String: *in.WashCare, Valid: true}
	}

	if in.Stretchable != nil {
		res.Stretchable = sql.NullString{String: *in.Stretchable, Valid: true}
	}

	if in.Distress != nil {
		res.Distress = sql.NullString{String: *in.Distress, Valid: true}
	}

	if in.Features != nil {
		res.Distress = sql.NullString{String: *in.Distress, Valid: true}
	}

	if in.Fade != nil {
		res.Fade = sql.NullString{String: *in.Fade, Valid: true}
	}

	if in.Fabric != nil {
		res.Fabric = *in.Fabric
	}

	if in.Category != nil {
		res.Category = *in.Category
	}
	if in.CountryOfOrigin != nil {
		res.CountryOfOrigin = *in.CountryOfOrigin
	}

	if in.Discount != nil {
		res.Discount = int64(*in.Discount)
	}

	if in.Quantity != nil {
		res.Quantity = int64(*in.Quantity)
	}

	res.UpdatedAt = time.Now()

	err = l.svcCtx.ProductModel.Update(l.ctx, res)

	var washCare, stretchable, features, fade *string

	if res.WashCare.Valid {
		washCare = &res.WashCare.String
	}

	if res.Stretchable.Valid {
		stretchable = &res.Stretchable.String
	}

	if res.Features.Valid {
		features = &res.Features.String
	}

	if res.Fade.Valid {
		fade = &res.Fade.String
	}

	return &product.ProductResponse{
		Id:              res.Id,
		Name:            res.Name,
		Price:           uint64(res.Price),
		Description:     res.Description,
		Size:            res.Size,
		Color:           res.Color,
		Brand:           res.Brand,
		Shade:           res.Shade,
		WashCare:        washCare,
		Stretchable:     stretchable,
		Features:        features,
		Fade:            fade,
		Fabric:          res.Fabric,
		Category:        res.Category,
		CountryOfOrigin: res.CountryOfOrigin,
		Discount:        uint32(res.Discount),
		Quantity:        uint64(res.Quantity),
		CreatedAt:       timestamppb.New(res.CreatedAt),
		UpdatedAt:       timestamppb.New(res.UpdatedAt),
		ArchivedAt:      nil,
	}, nil
}
