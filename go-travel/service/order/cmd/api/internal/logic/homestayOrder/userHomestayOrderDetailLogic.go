package homestayOrder

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"go-travel/common/ctxdata"
	"go-travel/common/tool"
	"go-travel/common/xerr"
	"go-travel/service/order/cmd/rpc/order"
	"go-travel/service/order/model"
	"go-travel/service/payment/cmd/rpc/payment"

	"go-travel/service/order/cmd/api/internal/svc"
	"go-travel/service/order/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserHomestayOrderDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserHomestayOrderDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserHomestayOrderDetailLogic {
	return &UserHomestayOrderDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserHomestayOrderDetailLogic) UserHomestayOrderDetail(req *types.UserHomestayOrderDetailReq) (resp *types.UserHomestayOrderDetailResp, err error) {
	// todo: add your logic here and delete this line
	userId := ctxdata.GetUidFromCtx(l.ctx)

	res, err := l.svcCtx.OrderRpc.HomestayOrderDetail(l.ctx, &order.HomestayOrderDetailReq{
		Sn: req.Sn,
	})
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrMsg("get homestay order detail fail"), " rpc get HomestayOrderDetail err:%v , sn : %s", err, req.Sn)
	}

	var typesOrderDetail types.UserHomestayOrderDetailResp
	if res.HomestayOrder != nil && res.HomestayOrder.UserId == userId {

		copier.Copy(&typesOrderDetail, res.HomestayOrder)

		//重置价格.
		typesOrderDetail.OrderTotalPrice = tool.Fen2Yuan(res.HomestayOrder.OrderTotalPrice)
		typesOrderDetail.FoodTotalPrice = tool.Fen2Yuan(res.HomestayOrder.FoodTotalPrice)
		typesOrderDetail.HomestayTotalPrice = tool.Fen2Yuan(res.HomestayOrder.HomestayTotalPrice)
		typesOrderDetail.HomestayPrice = tool.Fen2Yuan(res.HomestayOrder.HomestayPrice)
		typesOrderDetail.FoodPrice = tool.Fen2Yuan(res.HomestayOrder.FoodPrice)
		typesOrderDetail.MarketHomestayPrice = tool.Fen2Yuan(res.HomestayOrder.MarketHomestayPrice)

		//支付信息.
		if typesOrderDetail.TradeState != model.HomestayOrderTradeStateCancel && typesOrderDetail.TradeState != model.HomestayOrderTradeStateWaitPay {
			paymentResp, err := l.svcCtx.PaymentRpc.GetPaymentSuccessRefundByOrderSn(l.ctx, &payment.GetPaymentSuccessRefundByOrderSnReq{
				OrderSn: res.HomestayOrder.Sn,
			})
			if err != nil {
				logx.WithContext(l.ctx).Errorf("Failed to get order payment information err : %v , orderSn:%s", err, res.HomestayOrder.Sn)
			}

			if paymentResp.PaymentDetail != nil {
				typesOrderDetail.PayTime = paymentResp.PaymentDetail.PayTime
				typesOrderDetail.PayType = paymentResp.PaymentDetail.PayMode
			}
		}

		return &typesOrderDetail, nil
	}

	return nil, nil
}
