package logic

import (
	"context"
	"errors"

	"go-travel/service/travel/cmd/rpc/internal/svc"
	"go-travel/service/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteHomestayBusinessLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteHomestayBusinessLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteHomestayBusinessLogic {
	return &DeleteHomestayBusinessLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteHomestayBusinessLogic) DeleteHomestayBusiness(in *pb.DeleteHomestayBusinessReq) (*pb.DeleteHomestayBusinessResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.HomestayBusinessModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("homestay business is not exist")
	}
	err = l.svcCtx.HomestayActivityModel.Delete(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("delete homestay bussiness error")
	}
	return &pb.DeleteHomestayBusinessResp{
		Pong: 1,
	}, nil
}
