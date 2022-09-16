package order

import (
	"context"
	"go-travel/service/order/cmd/rpc/order"
	"go-travel/service/order/cmd/rpc/pb"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderListLogic {
	return &OrderListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type OrderListApiResp struct {
	Total int64
	List  []*pb.ListOrderItemResp
}

func (l *OrderListLogic) OrderList(req *types.ListOrderReq) (resp *OrderListApiResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.OrderRpc.ListOrder(l.ctx, &order.ListOrderReq{
		//Info:     req.Info,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}

	return &OrderListApiResp{
		Total: res.Total,
		List:  res.List,
	}, nil
}
