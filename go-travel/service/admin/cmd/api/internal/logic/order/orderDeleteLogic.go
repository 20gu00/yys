package order

import (
	"context"
	"go-travel/service/order/cmd/rpc/order"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderDeleteLogic {
	return &OrderDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderDeleteLogic) OrderDelete(req *types.DeleteOrderReq) (resp *types.DeleteOrderResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.OrderRpc.DeleteOrder(l.ctx, &order.DeleteOrderReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}

	return &types.DeleteOrderResp{
		Pong: res.Pong,
	}, nil
}
