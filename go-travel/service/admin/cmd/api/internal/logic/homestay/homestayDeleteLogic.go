package homestay

import (
	"context"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayDeleteLogic {
	return &HomestayDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayDeleteLogic) HomestayDelete(req *types.DeleteHomestayReq) (resp *types.DeleteHomestayResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.TravelRpc.DeleteHomestay(l.ctx, &travel.DeleteHomestayReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeleteHomestayResp{
		Pong: res.Pong,
	}, nil
}
