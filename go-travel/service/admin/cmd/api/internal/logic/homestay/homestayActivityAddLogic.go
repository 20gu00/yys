package homestay

import (
	"context"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayActivityAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayActivityAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayActivityAddLogic {
	return &HomestayActivityAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayActivityAddLogic) HomestayActivityAdd(req *types.AddHomestayActivityReq) (resp *types.AddHomestayActivityResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.TravelRpc.AddHomestayActivity(l.ctx, &travel.AddHomestayActivityReq{
		Id:         req.Id,
		CreateTime: req.CreateTime,
		UpdateTime: req.UpdateTime,
		DeleteTime: req.DeleteTime,
		DelState:   req.DelState,
		RowType:    req.RowType,
		DataId:     req.DataId,
		RowStatus:  req.RowStatus,
		Version:    req.Version,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddHomestayActivityResp{
		Pong: res.Pong,
	}, nil
}
