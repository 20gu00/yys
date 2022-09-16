package manager

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"
	"go-travel/service/admin/cmd/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChangePwdLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChangePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePwdLogic {
	return &ChangePwdLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChangePwdLogic) ChangePwd(req *types.ChangePwdReq) (resp *types.ChangePwdResp, err error) {
	// todo: add your logic here and delete this line
	if len(strings.TrimSpace(fmt.Sprintf("%s", req.Id))) == 0 || len(strings.TrimSpace(fmt.Sprintf("%s", req.Password))) == 0 {
		reqStr, _ := json.Marshal(req) //func Marshal(v any) ([]byte, error)
		logx.WithContext(l.ctx).Errorf("输入的id或者密码不合法,参数:%s", reqStr)
		return nil, errors.New("输入数据不符合")
	}

	_, err = l.svcCtx.AdminRpc.ChangePwd(l.ctx, &admin.ChangePwdReq{
		Id:       req.Id,
		Password: req.Password,
	})
	if err != nil {
		logx.WithContext(l.ctx).Errorf("修改id为 %s的密码失败，%s", req.Id, err.Error())
		return nil, errors.New("修改密码失败")
	}

	return &types.ChangePwdResp{
		Pong: 1,
	}, nil
}
