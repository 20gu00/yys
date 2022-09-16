package logic

import (
	"context"
	"errors"
	"go-travel/service/travel/cmd/rpc/internal/svc"
	"go-travel/service/travel/cmd/rpc/pb"
	"go-travel/service/travel/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddHomestayLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddHomestayLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddHomestayLogic {
	return &AddHomestayLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// provide to admin use
func (l *AddHomestayLogic) AddHomestay(in *pb.AddHomestayReq) (*pb.AddHomestayResp, error) {
	// todo: add your logic here and delete this line
	_, err := l.svcCtx.HomestayModel.FindOne(l.ctx, in.Id)
	if err == nil {
		return nil, errors.New("homestay is exist")
	}

	//string->time.Time
	//createtime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	//updatetime, _ := time.Parse("2006-01-02 15:04:05", in.UpdateTime)
	//deltime, _ := time.Parse("2006-01-02 15:04:05", in.DelTime)
	_, err = l.svcCtx.HomestayModel.InsertById(l.ctx, &model.Homestay{
		Id: in.Id,
		//CreateTime:          createtime,
		//UpdateTime:          updatetime,
		//DeleteTime:          deltime,
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
	})
	if err != nil {
		return nil, errors.New("add homestay error")
	}

	return &pb.AddHomestayResp{
		Pong: 1,
	}, nil
}
