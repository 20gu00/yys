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

type UpdateHomestayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateHomestayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateHomestayLogic {
	return &UpdateHomestayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateHomestayLogic) UpdateHomestay(in *pb.UpdateHomestayReq) (*pb.UpdateHomestayResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.HomestayModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("homestay不存在,参数:%d,异常:%s", in.Id, err.Error())
		return nil, errors.New("homestay not exist")
	default:
		logx.WithContext(l.ctx).Errorf("homestay search操作未成功,参数:%d,异常:%s", in.Id, err.Error())
		return nil, err
	}

	//string->time.Time
	createtime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	updatetime, _ := time.Parse("2006-01-02 15:04:05", in.UpdateTime)
	deletetime, _ := time.Parse("2006-01-02 15:04:05", in.DelTime)
	data := &model.Homestay{
		Id:                  in.Id,
		CreateTime:          createtime,
		UpdateTime:          updatetime,
		DeleteTime:          deletetime,
		Version:             in.Version,
		Title:               in.Title,
		SubTitle:            in.SubTitle,
		Banner:              in.Banner,
		Info:                in.Info,
		PeopleNum:           in.PeopleNum,
		HomestayBusinessId:  in.HomestayBusinessId,
		UserId:              in.UserId,
		RowState:            in.RowState,
		FoodInfo:            in.FoodInfo,
		FoodPrice:           in.FoodPrice,
		HomestayPrice:       in.HomestayPrice,
		MarketHomestayPrice: in.MarketHomestayPrice,
	}
	_, err = l.svcCtx.HomestayModel.UpdateNoSess(l.ctx, data)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("update homestay error(rpc)")
		return nil, errors.New("update homestay error(rpc)")
	}

	return &pb.UpdateHomestayResp{
		Pong: 1,
	}, nil
}
