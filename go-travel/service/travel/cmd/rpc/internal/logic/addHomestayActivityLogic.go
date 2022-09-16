package logic

import (
	"context"
	"errors"

	"go-travel/service/travel/cmd/rpc/internal/svc"
	"go-travel/service/travel/cmd/rpc/pb"
	"go-travel/service/travel/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddHomestayActivityLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHomestayActivityLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHomestayActivityLogic {
	return &AddHomestayActivityLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// homestayActivity
func (l *AddHomestayActivityLogic) AddHomestayActivity(in *pb.AddHomestayActivityReq) (*pb.AddHomestayActivityResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.HomestayActivityModel.FindOne(l.ctx, in.Id)
	if err == nil {
		return nil, errors.New("homestayactivity is exist")
	}

	//string->time.Time
	//createtime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	//updatetime, _ := time.Parse("2006-01-02 15:04:05", in.UpdateTime)
	//deletetime, _ := time.Parse("2006-01-02 15:04:05", in.DeleteTime)

	_, err = l.svcCtx.HomestayActivityModel.Insert(l.ctx, &model.HomestayActivity{
		Id: in.Id,
		//CreateTime: createtime,
		//UpdateTime: updatetime,
		//DeleteTime: deletetime,
		DelState:   in.DelState,
		RowType:    in.RowType,
		DataId:     in.DataId,
		RowStatus:  in.RowStatus,
		Version:    in.Version,
		ActivityId: in.ActivityId,
	})

	if err != nil {
		return nil, errors.New("add homestay activity error")
	}
	return &pb.AddHomestayActivityResp{
		Pong: 1,
	}, nil
}
