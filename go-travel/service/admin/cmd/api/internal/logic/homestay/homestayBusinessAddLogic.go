package homestay

import (
	"context"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBusinessAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayBusinessAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessAddLogic {
	return &HomestayBusinessAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *HomestayBusinessAddLogic) HomestayBusinessAdd(req *types.AddHomestayBusinessReq) (resp *types.AddHomestayBusinessResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.TravelRpc.AddHomestayBusiness(l.ctx, &travel.AddHomestayBusinessReq{
		Id:          req.Id,
		CreateTime:  req.CreateTime,
		UpdateTime:  req.UpdateTime,
		DeleteTime:  req.DeleteTime,
		DelState:    req.DelState,
		Title:       req.Title,
		UserId:      req.UserId,
		Info:        req.Info,
		BossInfo:    req.BossInfo,
		LicenseFron: req.LicenseFron,
		LicenseBack: req.LicenseBack,
		RowState:    req.RowState,
		Star:        float32(req.Star), //star
		Tag:         req.Tag,
		Cover:       req.Cover,
		HeaderImg:   req.HeaderImg,
		Version:     req.Version,
	})
	if err != nil {
		return nil, err
	}
	return &types.AddHomestayBusinessResp{
		Pong: res.Pong,
	}, nil
}
