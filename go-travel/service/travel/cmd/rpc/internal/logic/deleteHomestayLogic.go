package logic

import (
	"context"
	"errors"

	"go-travel/service/travel/cmd/rpc/internal/svc"
	"go-travel/service/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHomestayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHomestayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHomestayLogic {
	return &DeleteHomestayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteHomestayLogic) DeleteHomestay(in *pb.DeleteHomestayReq) (*pb.DeleteHomestayResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.HomestayModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("homestay is not exist")
	}
	err = l.svcCtx.HomestayModel.DeleteById(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("delete homestay errors")
	}
	return &pb.DeleteHomestayResp{
		Pong: 1,
	}, nil
}
