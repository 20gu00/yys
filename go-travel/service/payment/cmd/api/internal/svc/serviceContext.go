package svc

import (
	"go-travel/service/order/cmd/rpc/order"
	"go-travel/service/payment/cmd/api/internal/config"
	"go-travel/service/payment/cmd/rpc/payment"
	"go-travel/service/usercenter/cmd/rpc/usercenter"

	"github.com/wechatpay-apiv3/wechatpay-go/core"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	WxPayClient *core.Client

	PaymentRpc    payment.Payment
	OrderRpc      order.Order
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		PaymentRpc:    payment.NewPayment(zrpc.MustNewClient(c.PaymentRpcConf)),
		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
