package homestay

import (
	"context"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestaycommentUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestaycommentUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestaycommentUpdateLogic {
	return &HomestaycommentUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestaycommentUpdateLogic) HomestaycommentUpdate(req *types.UpdateHomestaycommentReq) (resp *types.UpdateHomestaycommentResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.TravelRpc.UpdateHomestayComment(l.ctx, &travel.UpdateHomestaycommentReq{
		Id:         req.Id,
		CreateTime: req.CreateTime,
		UpdateTime: req.UpdateTime,
		DeleteTime: req.DeleteTime,
		DelState:   req.DelState,
		Version:    req.Version,
		HomestayId: req.HomestayId,
		UserId:     req.UserId,
		Star:       float32(req.Star),
		Content:    req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &types.UpdateHomestaycommentResp{
		Pong: res.Pong,
	}, nil
}
