package manager

import (
	"context"
	"encoding/json"
	"errors"
	"go-travel/service/admin/cmd/rpc/admin"
	"strconv"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ManagerInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewManagerInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ManagerInfoLogic {
	return &ManagerInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ManagerInfoLogic) ManagerInfo() (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()
	res, err := l.svcCtx.AdminRpc.AdminInfo(l.ctx, &admin.UserInfoReq{
		Id: userId,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据userId: %s,查询用户异常:%s", strconv.FormatInt(userId, 10), err.Error())
		return nil, errors.New("查询用户失败")
	}
	//copier.Copy(resp,res)   <=> field similar
	return &types.UserInfoResp{
		Id:           res.Id,
		UserId:       res.UserId,
		Name:         res.UserName,
		Introduction: res.Introduction,
	}, nil
}
