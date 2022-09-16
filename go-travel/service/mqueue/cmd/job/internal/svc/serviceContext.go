package svc

import (
	"github.com/hibiken/asynq"
	"github.com/silenceper/wechat/v2/miniprogram"
	"github.com/zeromicro/go-zero/zrpc"
	"go-travel/service/mqueue/cmd/job/internal/config"
	"go-travel/service/order/cmd/rpc/order"
	"go-travel/service/usercenter/cmd/rpc/usercenter"
)

type ServiceContext struct {
	Config      config.Config
	AsynqServer *asynq.Server
	MiniProgram *miniprogram.MiniProgram

	OrderRpc      order.Order
	UsercenterRpc usercenter.Usercenter
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:        c,
		AsynqServer:   newAsynqServer(c),
		MiniProgram:   newMiniprogramClient(c),
		OrderRpc:      order.NewOrder(zrpc.MustNewClient(c.OrderRpcConf)),
		UsercenterRpc: usercenter.NewUsercenter(zrpc.MustNewClient(c.UsercenterRpcConf)),
	}
}
