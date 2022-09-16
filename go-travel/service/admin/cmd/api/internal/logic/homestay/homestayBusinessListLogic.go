package homestay

import (
	"context"
	"go-travel/service/travel/cmd/rpc/pb"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayBusinessListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayBusinessListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayBusinessListLogic {
	return &HomestayBusinessListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type HomestayBusinessListApiResp struct {
	Total int64
	List  []*pb.ListHomestayBusinessItemResp
}

func (l *HomestayBusinessListLogic) HomestayBusinessList(req *types.ListHomestayBusinessReq) (resp *HomestayBusinessListApiResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.TravelRpc.ListHomestayBussiness(l.ctx, &travel.ListHomestayBusinessReq{
		//Info:     req.Info,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return &HomestayBusinessListApiResp{
		Total: res.Total,
		List:  res.List,
	}, nil
}
