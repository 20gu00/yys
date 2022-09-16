package manager

import (
	"context"
	"errors"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"
	"go-travel/service/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserDeleteLogic {
	return &UserDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserDeleteLogic) UserDelete(req *types.DeleteUserReq) (resp *types.DeleteUserResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UsercenterRpc.UserDelete(l.ctx, &usercenter.UserDeleteReq{
		Id: req.Id,
	})

	if err != nil {
		return nil, errors.New("delete user of admincenter by id error")
	}

	return &types.DeleteUserResp{
		Pong: res.Pong,
	}, nil
}
