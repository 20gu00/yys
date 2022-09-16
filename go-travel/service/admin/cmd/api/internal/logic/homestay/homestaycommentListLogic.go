package homestay

import (
	"context"
	"go-travel/service/travel/cmd/rpc/pb"
	"go-travel/service/travel/cmd/rpc/travel"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type HomestaycommentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHomestaycommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *HomestaycommentListLogic {
	return &HomestaycommentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type HomestayCommentListApiResp struct {
	Total int64
	List  []*pb.ListHomestaycommentItemResp
}

func (l *HomestaycommentListLogic) HomestaycommentList(req *types.ListHomestaycommentReq) (resp *HomestayCommentListApiResp, err error) {
	// todo: add your logic here and delete this line

	res, err := l.svcCtx.TravelRpc.ListHomestayComment(l.ctx, &travel.ListHomestaycommentReq{
		//Info:     req.Info,
		Page:     req.Page,
		PageSize: req.PageSize,
	})
	if err != nil {
		return nil, err
	}
	return &HomestayCommentListApiResp{
		Total: res.Total,
		List:  res.List,
	}, nil
	return
}
