package logic

import (
	"context"
	"errors"

	"go-travel/service/order/cmd/rpc/internal/svc"
	"go-travel/service/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteOrderLogic {
	return &DeleteOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteOrderLogic) DeleteOrder(in *pb.DeleteOrderReq) (*pb.DeleteOrderResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.HomestayOrderModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("order is not exist")
	}
	err = l.svcCtx.HomestayOrderModel.DeleteById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteOrderResp{
		Pong: 1,
	}, nil
}
