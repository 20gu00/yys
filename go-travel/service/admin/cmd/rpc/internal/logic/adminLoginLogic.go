package logic

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"go-travel/service/admin/cmd/rpc/admin"
	"go-travel/service/admin/cmd/rpc/internal/svc"
	"go-travel/service/admin/cmd/rpc/pb"
	"go-travel/service/admin/pkg/jwtx"
	"go-travel/service/admin/pkg/salt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type AdminLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAdminLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLoginLogic {
	return &AdminLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AdminLoginLogic) AdminLogin(in *pb.LoginReq) (*pb.LoginResp, error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.AdminModel.FindOneByName(l.ctx, in.UserName) //username unique or model limit
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%s,异常:%s", in.UserName, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("用户登录失败,参数:%s,异常:%s", in.UserName, err.Error())
		return nil, err
	}

	saltPassword := salt.Saltpassword(userInfo.Salt, in.Password)
	if userInfo.Password != saltPassword {
		logx.WithContext(l.ctx).Errorf("用户密码不正确,参数:%s", in.Password)
		return nil, errors.New("用户密码不正确")
	}
	now := time.Now().Unix() //second
	accessExpire := l.svcCtx.Config.JWT.AccessExpire
	jwtToken, err := jwtx.GetJwtToken(l.svcCtx.Config.JWT.AccessSecret, now, l.svcCtx.Config.JWT.AccessExpire, userInfo.Id)

	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("生成token失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}

	resp := &admin.LoginResp{
		Token:        jwtToken,
		AccessExpire: now + accessExpire,
		Id:           userInfo.Id,
		UserName:     userInfo.UserName,
		Salt:         userInfo.Salt,
		UserId:       userInfo.UserId,
	}

	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(resp)
	logx.WithContext(l.ctx).Infof("登录成功,参数:%s,响应:%s", reqStr, listStr)
	return resp, nil
}
