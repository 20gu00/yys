package logic

import (
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"go-travel/service/order/model"
	"time"

	"go-travel/service/order/cmd/rpc/internal/svc"
	"go-travel/service/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderLogic {
	return &UpdateOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateOrderLogic) UpdateOrder(in *pb.UpdateOrderReq) (*pb.UpdateOrderResp, error) {
	// todo: add your logic here and delete this line
	loc, _ := time.LoadLocation("Asia/Shanghai")

	_, err := l.svcCtx.HomestayOrderModel.FindOne(l.ctx, in.Id)
	switch err {
	case nil:
	case sqlc.ErrNotFound:
		logx.WithContext(l.ctx).Errorf("order不存在,参数:%d,异常:%s", in.Id, err.Error())
		return nil, errors.New("order不存在")
	default:
		logx.WithContext(l.ctx).Errorf("order search操作未成功,参数:%d,异常:%s", in.Id, err.Error())
		return nil, err
	}
	createtime, _ := time.Parse("2006-01-02 15:04:05", in.CreateTime)
	updatetime, _ := time.Parse("2006-01-02 15:04:05", in.UpdateTime)
	deletetime, _ := time.Parse("2006-01-02", in.DeleteTime)
	//all time insert into mysql need and can use time.Unix()
	//mysql every date type 0000-00-00

	data := &model.HomestayOrder{
		Id:            in.Id,
		CreateTime:    createtime,
		UpdateTime:    updatetime,
		DeleteTime:    deletetime,
		DelState:      in.DelState,
		Version:       in.Version,
		Sn:            in.Sn,
		UserId:        in.UserId,
		HomestayId:    in.HomestayId,
		Title:         in.Title,
		SubTitle:      in.SubTitle,
		Cover:         in.Cover,
		Info:          in.Info,
		PeopleNum:     in.PeopleNum,
		RowType:       in.RowType,
		NeedFood:      in.NeedFood,
		LiveStartDate: time.Unix(in.LiveStartDate, 0).In(loc),
		LiveEndDate:   time.Unix(in.LiveEndDate, 0).In(loc),
	}
	_, err = l.svcCtx.HomestayOrderModel.UpdateNoTrans(l.ctx, data)

	if err != nil {
		//terminal
		logx.WithContext(l.ctx).Errorf("update order error(rpc),%s", err)
		//postman
		return nil, errors.New("update order error(rpc)")
	}
	return &pb.UpdateOrderResp{
		Pong: 1,
	}, nil
}
