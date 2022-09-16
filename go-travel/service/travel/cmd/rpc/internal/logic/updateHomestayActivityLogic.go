package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"go-travel/service/travel/model"
	"time"

	"go-travel/service/travel/cmd/rpc/internal/svc"
	"go-travel/service/travel/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateHomestayActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomestayActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomestayActivityLogic {
	return &UpdateHomestayActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateHomestayActivityLogic) UpdateHomestayActivity(in *pb.UpdateHomestayActivityReq) (*pb.UpdateHomestayActivityResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.HomestayActivityModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("homestayActivity不存在,参数:%d,异常:%s", in.Id, err.Error())
		return nil, errors.New("homestayActivity不存在")
	default:
		logx.WithContext(l.ctx).Errorf("homestayActivity search操作未成功,参数:%d,异常:%s", in.Id, err.Error())
		return nil, err
	}

	//string->time.Time
	createtime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	updatetime, _ := time.Parse("2006-01-02 15:04:05", in.UpdateTime)
	deletetime, _ := time.Parse("2006-01-02 15:04:05", in.DeleteTime)
	data := &model.HomestayActivity{
		Id:         in.Id,
		CreateTime: createtime,
		UpdateTime: updatetime,
		DeleteTime: deletetime,
		DelState:   in.DelState,
		RowType:    in.RowType,
		DataId:     in.DataId,
		RowStatus:  in.RowStatus,
		Version:    in.Version,
	}
	err = l.svcCtx.HomestayActivityModel.Update(l.ctx, data)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("update homestayActivity error(rpc)")
		return nil, errors.New("update homestayActivity error(rpc)")
	}

	return &pb.UpdateHomestayActivityResp{
		Pong: 1,
	}, nil
}
