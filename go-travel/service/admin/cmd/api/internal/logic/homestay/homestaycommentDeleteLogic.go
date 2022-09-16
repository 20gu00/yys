package homestay

import (
	"context"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestaycommentDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestaycommentDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestaycommentDeleteLogic {
	return &HomestaycommentDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestaycommentDeleteLogic) HomestaycommentDelete(req *types.DeleteHomestaycommentReq) (resp *types.DeleteHomestaycommentResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.TravelRpc.DeleteHomestayComment(l.ctx, &travel.DeleteHomestaycommentReq{
		Id: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &types.DeleteHomestaycommentResp{
		Pong: res.Pong,
	}, nil
	return
}
