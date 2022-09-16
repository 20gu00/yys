package manager

import (
	"context"
	"errors"
	"go-travel/service/usercenter/cmd/rpc/pb"
	"go-travel/service/usercenter/cmd/rpc/usercenter"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type UserListApiResp struct {
	Total int64
	List  []*pb.UserListItemResp
}

func NewUserListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserListLogic {
	return &UserListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserListLogic) UserList(req *types.ListUserReq) (resp *UserListApiResp, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UsercenterRpc.UserList(l.ctx, &usercenter.UserListReq{
		PageNo:   req.Page,
		PageSize: req.PageSize,
		//Info:     req.Info,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("page: %s and pagesize:%s 查询用户异常:%s", req.Page, req.PageSize, err.Error())
		return nil, errors.New("user list error")
	}

	//use copier.Copy id(the first field will always be 0) because my [] field one is Id one is ID
	//outlist := make([]types.ListUserItemResp, 0)
	//var outlist []types.ListUserItemResp
	//err = copier.Copy(&outlist, res.List) //dest use &

	return &UserListApiResp{
		Total: res.Total,
		List:  res.List,
	}, nil
}
