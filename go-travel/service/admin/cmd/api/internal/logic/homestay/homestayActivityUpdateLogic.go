package homestay

import (
	"context"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayActivityUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayActivityUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayActivityUpdateLogic {
	return &HomestayActivityUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayActivityUpdateLogic) HomestayActivityUpdate(req *types.UpdateHomestayActivityReq) (resp *types.UpdateHomestayActivityResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.TravelRpc.UpdateHomestayActivity(l.ctx, &travel.UpdateHomestayActivityReq{
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
	return &types.UpdateHomestayActivityResp{
		Pong: res.Pong,
	}, nil
}
