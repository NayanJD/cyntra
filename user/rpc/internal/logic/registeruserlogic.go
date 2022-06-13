package logic

import (
	"context"
	"errors"
	"time"

	"cyntra/user/rpc/internal/svc"
	"cyntra/user/rpc/model"
	"cyntra/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type RegisterUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterUserLogic {
	return &RegisterUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterUserLogic) RegisterUser(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, in.Username)

	if err != nil && !errors.Is(err, sqlc.ErrNotFound) {
		l.Logger.Error(err)
		return nil, errors.New("Unknown error")
	}

	if res != nil {
		l.Logger.Error("User Already Exists")
		return nil, status.Error(codes.AlreadyExists, "User Already Exists")
	}

	role, err := l.svcCtx.RoleModel.FindOneByName(l.ctx, "default_user")

	if err != nil {
		l.Logger.Error(err)
		return nil, status.Error(codes.Unknown, codes.Unknown.String())
	}

	hashedPassword, _ := HashPassword(in.Password)

	newUser := model.User{
		Username:       in.Username,
		HashedPassword: hashedPassword,
		FirstName:      in.FirstName,
		LastName:       in.LastName,
		Gender:         in.Gender,
		Dob:            in.Dob.AsTime(),
		RoleId:         role.Id,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	_, err = l.svcCtx.UserModel.Insert(l.ctx, &newUser)

	if err != nil {
		l.Logger.Error(err)
		return nil, status.Error(codes.Unknown, codes.Unknown.String())
	}

	return &user.RegisterResponse{
		Id:        newUser.Id,
		Username:  newUser.Username,
		FirstName: newUser.FirstName,
		LastName:  newUser.LastName,
		Gender:    newUser.Gender,
		Dob:       timestamppb.New(newUser.Dob),
	}, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
