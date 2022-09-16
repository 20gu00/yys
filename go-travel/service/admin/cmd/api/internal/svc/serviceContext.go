package svc

import (
	"go-travel/service/admin/cmd/api/internal/config"
	"go-travel/service/admin/cmd/rpc/admin"
	"go-travel/service/order/cmd/rpc/order"
	"go-travel/service/payment/cmd/rpc/payment"
	"go-travel/service/travel/cmd/rpc/travel"
	"go-travel/service/usercenter/cmd/rpc/usercenter"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config config.Config

	AdminRpc      admin.Admin
	UsercenterRpc usercenter.Usercenter
	PaymentRpc    payment.Payment
	OrderRpc      order.Order
	TravelRpc     travel.Travel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		AdminRpc:      admin.NewAdmin(zrpc.MustNewClient(c.AdminRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
		PaymentRpc:    payment.NewPayment(zrpc.MustNewClient(c.PaymentRpcConf)),
		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		TravelRpc:     travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
	}
}
