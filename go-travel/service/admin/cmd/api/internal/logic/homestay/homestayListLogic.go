package homestay

import (
	"context"
	"go-travel/service/travel/cmd/rpc/pb"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayListLogic {
	return &HomestayListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type HomestayListApiResp struct {
	Total int64
	List  []*pb.ListHomestayItemResp
}

func (l *HomestayListLogic) HomestayList(req *types.ListHomestayReq) (resp *HomestayListApiResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.TravelRpc.ListHomestay(l.ctx, &travel.ListHomestayReq{
		//Info:     req.Info,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return &HomestayListApiResp{
		Total: res.Total,
		List:  res.List,
	}, nil
}
