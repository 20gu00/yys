package logic

import (
	"context"
	"errors"

	"go-travel/service/admin/cmd/rpc/internal/svc"
	"go-travel/service/admin/cmd/rpc/pb"
	"go-travel/service/admin/model"
	"go-travel/service/admin/pkg/salt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
)

type ChangePwdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChangePwdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChangePwdLogic {
	return &ChangePwdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ChangePwdLogic) ChangePwd(in *pb.ChangePwdReq) (*pb.ChangePwdResp, error) {
	// todo: add your logic here and delete this line
	userInfo, err := l.svcCtx.AdminModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("用户不存在,参数:%d,异常:%s", in.Id, err.Error())
		return nil, errors.New("用户不存在")
	default:
		logx.WithContext(l.ctx).Errorf("密码修改操作未成功,参数:%d,异常:%s", in.Id, err.Error())
		return nil, err
	}

	saltPassword := salt.Saltpassword(userInfo.Salt, in.Password)
	data := &model.Admin{
		Id:       in.Id,
		UserName: userInfo.UserName,
		Salt:     userInfo.Salt,
		Password: saltPassword,
		//UserId:
		//CreateAt time.Time `db:"create_at"` // 新增时间
		//UpdateAt time.Time `db:"update_at"` // 更新时间
		Introduction: userInfo.Introduction,
		UserId:       userInfo.UserId,
	}
	err = l.svcCtx.AdminModel.Update(l.ctx, data)

	if err != nil {
		logx.WithContext(l.ctx).Errorf("change passowrd error(rpc)")
		return nil, errors.New("change password error(rpc)")
	}
	return &pb.ChangePwdResp{
		Pong: 1,
	}, nil
}
