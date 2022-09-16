package logic

import (
	"context"
	"errors"
	"go-travel/service/order/model"
	"time"

	"go-travel/service/order/cmd/rpc/internal/svc"
	"go-travel/service/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddOrderLogic {
	return &AddOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// provide to admin use
func (l *AddOrderLogic) AddOrder(in *pb.AddOrderReq) (*pb.AddOrderResp, error) {
	// todo: add your logic here and delete this line
	loc, _ := time.LoadLocation("Asia/Shanghai")
	_, err := l.svcCtx.HomestayOrderModel.FindOne(l.ctx, in.Id)
	if err == nil {
		return nil, errors.New("order is exist")
	}
	//if in.Secret == "" {
	//	in.Secret = MD5.MD5(in.AppID) //or base timestamp
	//}

	//string->time.Time
	//mysql datetime	time.Time
	createtime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	updatetime, _ := time.Parse("2006-01-02 15:04:05", in.UpdateTime)
	deletetime, _ := time.Parse("2006-01-02", in.DeleteTime)
	//mysql date
	//livestartdate, _ := time.Parse("2006-01-02 15:04:05", in.LiveStartDate)
	//liveenddate, _ := time.Parse("2006-01-02 15:04:05", in.LiveEndDate)

	_, err = l.svcCtx.HomestayOrderModel.InsertById(l.ctx, &model.HomestayOrder{
		Id:                  in.Id,
		CreateTime:          createtime,
		UpdateTime:          updatetime,
		DeleteTime:          deletetime,
		Version:             in.Version,
		DelState:            in.DelState,
		Sn:                  in.Sn,
		HomestayId:          in.HomestayId,
		UserId:              in.UserId,
		Title:               in.Title,
		SubTitle:            in.SubTitle,
		Cover:               in.Cover,
		Info:                in.Info,
		PeopleNum:           in.PeopleNum,
		RowType:             in.RowType,
		NeedFood:            in.NeedFood,
		FoodInfo:            in.FoodInfo,
		FoodPrice:           in.FoodPrice,
		HomestayPrice:       in.HomestayPrice,
		MarketHomestayPrice: in.MarketHomestayPrice,
		HomestayBusinessId:  in.HomwstayBusinessId,
		HomestayUserId:      in.HomestayUserId,
		LiveStartDate:       time.Unix(in.LiveStartDate, 0).In(loc), //return time.Time
		LiveEndDate:         time.Unix(in.LiveEndDate, 0).In(loc),
		LivePeopleNum:       in.LivePeopleNum,
		TradeState:          in.TradeState,
		TradeCode:           in.TradeCode,
		Remark:              in.Remark,
		OrderTotalPrice:     in.OrderTotalPrice,
		FoodTotalPrice:      in.FoodTotalPrice,
		HomestayTotalPrice:  in.HomestayTotalPrice,
	})
	if err != nil {
		return nil, err
	}
	return &pb.AddOrderResp{
		Pong: 1,
	}, nil
}
