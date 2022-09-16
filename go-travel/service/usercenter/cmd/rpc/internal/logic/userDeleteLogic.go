package logic

import (
	"context"
	"errors"

	"go-travel/service/usercenter/cmd/rpc/internal/svc"
	"go-travel/service/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserDeleteLogic) UserDelete(in *pb.UserDeleteReq) (*pb.UserDeleteResp, error) {
	// todo: add your logic here and delete this line
	if _, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id); err != nil {
		return nil, errors.New("user not found")
	}

	if err := l.svcCtx.UserModel.DeleteById(l.ctx, in.Id); err != nil {
		return nil, errors.New("delete user by id error")
	}

	return &pb.UserDeleteResp{
		Pong: 1,
	}, nil
}
