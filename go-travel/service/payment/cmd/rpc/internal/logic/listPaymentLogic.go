package logic

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"go-travel/service/payment/cmd/rpc/payment"

	"go-travel/service/payment/cmd/rpc/internal/svc"
	"go-travel/service/payment/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPaymentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPaymentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPaymentLogic {
	return &ListPaymentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// provide to admin use
func (l *ListPaymentLogic) ListPayment(in *pb.ListPaymentReq) (*pb.ListPaymentResp, error) {
	// todo: add your logic here and delete this line
	countbuilder := l.svcCtx.ThirdPaymentModel.CountBuilder("id")
	count, err := l.svcCtx.ThirdPaymentModel.FindCount(l.ctx, countbuilder)
	if err != nil {
		logx.WithContext(l.ctx).Errorf("查询payment count信息失败,参数,异常:%s", err.Error())
		return nil, errors.New("find payment count error(rpc)")
	}

	rowBuilder := l.svcCtx.ThirdPaymentModel.RowBuilder()
	all, err := l.svcCtx.ThirdPaymentModel.FindPageListByPage(l.ctx, rowBuilder, in.Page, in.PageSize, "id")
	if err != nil {
		reqStr, _ := json.Marshal(in)
		logx.WithContext(l.ctx).Errorf("分页查询所有payment table信息失败,参数:%s,异常:%s", reqStr, err.Error())
		return nil, errors.New("list payment error(rpc)")
	}

	var outList []*payment.ListPaymentItemResp
	for _, paymentlist := range all {
		createtime := paymentlist.CreateTime.Format("2006-01-02 15:04:05")
		updatetime := paymentlist.UpdateTime.Format("2006-01-02 15:04:05")
		deletetime := paymentlist.DeleteTime.Format("2006-01-02 15:04:05")
		paytime := paymentlist.PayTime.Format("2006-01-02 15:04:05")
		outItem := &payment.ListPaymentItemResp{ //& not&
			Id:             paymentlist.Id,
			CreateTime:     createtime,
			UpdateTime:     updatetime,
			DeleteTime:     deletetime,
			DelState:       paymentlist.DelState,
			Version:        paymentlist.Version,
			UserId:         paymentlist.UserId,
			PayMode:        paymentlist.PayMode,
			TradeState:     paymentlist.TradeState,
			TradeType:      paymentlist.TradeType,
			TransactionId:  paymentlist.TransactionId,
			TradeStateDesc: paymentlist.TradeStateDesc,
			OrderSn:        paymentlist.OrderSn,
			ServiceType:    paymentlist.ServiceType,
			PayStatus:      paymentlist.PayStatus,
			PayTotal:       paymentlist.PayTotal,
			PayTime:        paytime,
		}
		outList = append(outList, outItem)
	}
	reqStr, _ := json.Marshal(in)
	listStr, _ := json.Marshal(outList)
	logx.WithContext(l.ctx).Infof("分页查询所有payment信息,参数：%s,响应：%s", reqStr, listStr)

	return &pb.ListPaymentResp{
		Total: count,
		List:  outList,
	}, nil
}
