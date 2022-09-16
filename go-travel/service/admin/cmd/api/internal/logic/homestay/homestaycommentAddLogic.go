package homestay

import (
	"context"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestaycommentAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestaycommentAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestaycommentAddLogic {
	return &HomestaycommentAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestaycommentAddLogic) HomestaycommentAdd(req *types.AddHomestaycommentReq) (resp *types.AddHomestaycommentResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.TravelRpc.AddHomestayComment(l.ctx, &travel.AddHomestaycommentReq{
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
	return &types.AddHomestaycommentResp{
		Pong: res.Pong,
	}, nil
}
