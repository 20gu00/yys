package homestay

import (
	"context"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayUpdateLogic {
	return &HomestayUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayUpdateLogic) HomestayUpdate(req *types.UpdateHomestayReq) (resp *types.UpdateHomestayResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.TravelRpc.UpdateHomestay(l.ctx, &travel.UpdateHomestayReq{
		Id:                  req.Id,
		CreateTime:          req.CreateTime,
		UpdateTime:          req.UpdateTime,
		DelTime:             req.DelTime,
		Version:             req.Version,
		Title:               req.Title,
		SubTitle:            req.SubTitle,
		Banner:              req.Banner,
		Info:                req.Info,
		PeopleNum:           req.PeopleNum,
		HomestayBusinessId:  req.HomestayBusinessId,
		UserId:              req.UserId,
		RowState:            req.RowState,
		FoodInfo:            req.FoodInfo,
		FoodPrice:           req.FoodPrice,
		HomestayPrice:       req.HomestayPrice,
		MarketHomestayPrice: req.MarketHomestayPrice,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateHomestayResp{
		Pong: res.Pong,
	}, nil
}
