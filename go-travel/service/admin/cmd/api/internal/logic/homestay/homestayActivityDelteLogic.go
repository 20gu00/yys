package homestay

import (
	"context"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayActivityDelteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayActivityDelteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayActivityDelteLogic {
	return &HomestayActivityDelteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayActivityDelteLogic) HomestayActivityDelte(req *types.DeleteHomestayActivityReq) (resp *types.DeleteHomestayActivityResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.TravelRpc.DeleteHomestatActivity(l.ctx, &travel.DeleteHomestayActivityReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeleteHomestayActivityResp{
		Pong: res.Pong,
	}, nil
}
