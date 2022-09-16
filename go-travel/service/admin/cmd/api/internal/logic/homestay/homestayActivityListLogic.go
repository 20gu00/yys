package homestay

import (
	"context"
	"go-travel/service/travel/cmd/rpc/pb"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestayActivityListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestayActivityListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestayActivityListLogic {
	return &HomestayActivityListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type HomestayActivityListApiResp struct {
	Total int64
	List  []*pb.ListHomestayAcitivityItemResp
}

func (l *HomestayActivityListLogic) HomestayActivityList(req *types.ListHomestayActivityReq) (resp *HomestayActivityListApiResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.TravelRpc.ListHomestayActivity(l.ctx, &travel.ListHomestayActivityReq{
		//Info:     req.Info,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return &HomestayActivityListApiResp{
		Total: res.Total,
		List:  res.List,
	}, nil
}
