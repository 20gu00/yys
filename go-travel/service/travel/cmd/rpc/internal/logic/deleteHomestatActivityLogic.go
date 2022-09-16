package logic

import (
	"context"
	"errors"

	"go-travel/service/travel/cmd/rpc/internal/svc"
	"go-travel/service/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHomestatActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHomestatActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHomestatActivityLogic {
	return &DeleteHomestatActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteHomestatActivityLogic) DeleteHomestatActivity(in *pb.DeleteHomestayActivityReq) (*pb.DeleteHomestayResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.HomestayActivityModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("homestay activity is not exist")
	}
	err = l.svcCtx.HomestayActivityModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("delete homestay activity error")
	}
	return &pb.DeleteHomestayResp{
		Pong: 1,
	}, nil
}
