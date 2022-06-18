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

	res.ArchivedAt = sql.NullTime{Time: time.Now(), Valid: true}

	err = l.svcCtx.ProductModel.Update(l.ctx, res)

	if err != nil {
		l.Logger.Error(err)
		return nil, status.Error(codes.Unknown, codes.Unknown.String())
	}

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
		ArchivedAt:      timestamppb.New(res.ArchivedAt.Time),
	}, nil
}
