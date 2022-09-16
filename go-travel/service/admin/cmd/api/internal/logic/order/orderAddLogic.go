package order

import (
	"context"
	"go-travel/service/order/cmd/rpc/order"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type OrderAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewOrderAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrderAddLogic {
	return &OrderAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *OrderAddLogic) OrderAdd(req *types.AddOrderReq) (resp *types.AddOrderResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.OrderRpc.AddOrder(l.ctx, &order.AddOrderReq{
		Id:                  req.Id,
		CreateTime:          req.CreateTime,
		UpdateTime:          req.UpdateTime,
		DeleteTime:          req.DeleteTime,
		Version:             req.Version,
		DelState:            req.DelState,
		Sn:                  req.Sn,
		HomestayId:          req.HomestayId,
		UserId:              req.UserId,
		Title:               req.Title,
		SubTitle:            req.SubTitle,
		Cover:               req.Cover,
		Info:                req.Info,
		PeopleNum:           req.PeopleNum,
		RowType:             req.RowType,
		NeedFood:            req.NeedFood,
		FoodInfo:            req.FoodInfo,
		FoodPrice:           req.FoodPrice,
		HomestayPrice:       req.HomestayPrice,
		MarketHomestayPrice: req.MarketHomestayPrice,
		HomwstayBusinessId:  req.HomwstayBusinessId,
		HomestayUserId:      req.HomestayUserId,
		LiveStartDate:       req.LiveStartDate,
		LiveEndDate:         req.LiveEndDate,
		LivePeopleNum:       req.LivePeopleNum,
		TradeState:          req.TradeState,
		TradeCode:           req.TradeCode,
		Remark:              req.Remark,
		OrderTotalPrice:     req.OrderTotalPrice,
		FoodTotalPrice:      req.FoodTotalPrice,
		HomestayTotalPrice:  req.HomestayTotalPrice,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddOrderResp{
		Pong: res.Pong,
	}, nil
}
