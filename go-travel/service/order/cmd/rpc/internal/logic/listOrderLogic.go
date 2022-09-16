package logic

import (
	"context"
	"encoding/json"
	"go-travel/service/order/cmd/rpc/order"
	"time"

	"go-travel/service/order/cmd/rpc/internal/svc"
	"go-travel/service/order/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListOrderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListOrderLogic {
	return &ListOrderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListOrderLogic) ListOrder(in *pb.ListOrderReq) (*pb.ListOrderResp, error) {
	// todo: add your logic here and delete this line
	countbuilder := l.svcCtx.HomestayOrderModel.CountBuilder("id")
	count, err := l.svcCtx.HomestayOrderModel.FindCount(l.ctx, countbuilder)
	if err != nil {
		restr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("查询homestay_order信息count失败,参数%s,异常:%s", restr, err.Error())
		return nil, err
	}
	rowBuilder := l.svcCtx.HomestayOrderModel.RowBuilder()
	all, err := l.svcCtx.HomestayOrderModel.FindPageListByPage(l.ctx, rowBuilder, in.Page, in.PageSize, "id") //Pageno Current
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("分页查询所有homestay_order信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, err
	}
	loc, _ := time.LoadLocation("Asia/Shanghai")
	var outList []*order.ListOrderItemResp
	for _, orderlist := range all {
		//livestartdate := orderlist.LiveStartDate.Format("2006-01-02 15:04:05")
		//liveenddate := orderlist.LiveEndDate.Format("2006-01-02 15:04:05")

		outItem := &order.ListOrderItemResp{
			Id:                  orderlist.Id,
			Sn:                  orderlist.Sn,
			UserId:              orderlist.UserId, // 用户id
			Version:             orderlist.Version,
			DelState:            orderlist.DelState,
			HomestayId:          orderlist.HomestayId,
			Title:               orderlist.Title,
			SubTitle:            orderlist.SubTitle,
			Cover:               orderlist.Cover,
			Info:                orderlist.Info,
			PeopleNum:           orderlist.PeopleNum,
			RowType:             orderlist.RowType,
			NeedFood:            orderlist.NeedFood,
			FoodInfo:            orderlist.FoodInfo,
			FoodPrice:           orderlist.FoodPrice,
			HomestayPrice:       orderlist.HomestayPrice,
			MarketHomestayPrice: orderlist.MarketHomestayPrice,
			HomwstayBusinessId:  orderlist.HomestayBusinessId,
			HomestayUserId:      orderlist.HomestayUserId,
			LiveStartDate:       orderlist.LiveStartDate.In(loc).Unix(), //livestartdate,
			LiveEndDate:         orderlist.LiveEndDate.In(loc).Unix(),
			LivePeopleNum:       orderlist.LivePeopleNum,
			TradeState:          orderlist.TradeState,
			TradeCode:           orderlist.TradeCode,
			Remark:              orderlist.Remark,
			OrderTotalPrice:     orderlist.OrderTotalPrice,
			FoodTotalPrice:      orderlist.FoodTotalPrice,
			HomestayTotalPrice:  orderlist.HomestayTotalPrice,
		}
		outList = append(outList, outItem)
	}
	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(outList)
	logx.WithContext(l.ctx).Infof("分页查询所有order信息,参数：%s,响应：%s", reqStr, listStr)

	return &pb.ListOrderResp{
		Total: count,
		List:  outList,
	}, nil
}
