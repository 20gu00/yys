package logic

import (
	"context"
	"errors"
	"fmt"

	"go-travel/service/admin/cmd/rpc/admin"
	"go-travel/service/admin/cmd/rpc/internal/svc"
	"go-travel/service/admin/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type AdminInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminInfoLogic {
	return &AdminInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminInfoLogic) AdminInfo(in *pb.UserInfoReq) (*pb.UserInfoResp, error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.AdminModel.FindOneById(l.ctx, in.Id)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Infof("用户不存在,no userId: %s", in.Id)
		return nil, errors.New(fmt.Sprintf("用户不存在,no userId: %s", in.Id))
	default:
		return nil, err
	}
	return &admin.UserInfoResp{
		Id:           userInfo.Id,
		UserName:     userInfo.UserName,
		Introduction: userInfo.Introduction, //"super administrator",
		UserId:       userInfo.UserId,
		//Avatar:       "https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif", //https://gw.alipayobjects.com/zos/antfincdn/XAosXuNZyF/BiazfanxmamNRoxxVxka.png
		//Roles:        []string{"admin"},
	}, nil
}
