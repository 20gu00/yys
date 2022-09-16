package manager

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"go-travel/service/admin/cmd/api/internal/svc"
	"go-travel/service/admin/cmd/api/internal/types"
	"go-travel/service/admin/cmd/rpc/admin"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLoginLogic) AdminLogin(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	//param check
	//strings.TrimSpace(s string)会返回一个string类型的slice，并将最前面和最后面的ASCII定义的空格去掉，中间的空格不会去掉，如果遇到了\0等其他字符会认为是非空格。
	if len(strings.TrimSpace(req.UserName)) == 0 || len(strings.TrimSpace(req.Password)) == 0 {
		reqStr, _ := json.Marshal(req) //func Marshal(v any) ([]byte, error)
		logx.WithContext(l.ctx).Errorf("用户名或密码不能为空,请求信息失败,参数:%s", reqStr)
		return nil, errors.New("用户名或密码不能为空")
	}

	//res resp
	res, err := l.svcCtx.AdminRpc.AdminLogin(l.ctx, &admin.LoginReq{
		UserName: req.UserName,
		Password: req.Password,
	})

	if err != nil {
		logx.WithContext(l.ctx).Errorf("根据用户名: %s和密码: %s查询用户异常:%s", req.UserName, req.Password, err.Error())
		return nil, errors.New("登录失败")
	}

	//now := time.Now().Unix()
	//accessExpire := l.svcCtx.Config.Auth.AccessExpire
	//jwtToken, err := jwtx.GetJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, res.Id)

	return &types.LoginResp{
		AccessToken:  res.Token,
		AccessExpire: res.AccessExpire,
	}, nil
}
