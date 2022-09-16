package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-travel/service/order/cmd/api/internal/config"
	"go-travel/service/order/cmd/rpc/order"
	"go-travel/service/payment/cmd/rpc/payment"
	"go-travel/service/travel/cmd/rpc/travel"
)

type ServiceContext struct {
	Config config.Config

	OrderRpc   order.Order
	PaymentRpc payment.Payment
	TravelRpc  travel.Travel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		OrderRpc:   order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		PaymentRpc: payment.NewPayment(zrpc.MustNewClient(c.PaymentRpcConf)),
		TravelRpc:  travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
	}
}
