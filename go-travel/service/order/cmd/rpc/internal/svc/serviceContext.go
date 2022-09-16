package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
	"go-travel/service/order/cmd/rpc/internal/config"
	"go-travel/service/order/model"
	"go-travel/service/travel/cmd/rpc/travel"

	"github.com/hibiken/asynq"
)

type ServiceContext struct {
	Config config.Config

	AsynqClient        *asynq.Client
	TravelRpc          travel.Travel
	HomestayOrderModel model.HomestayOrderModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,

		AsynqClient:        newAsynqClient(c),
		TravelRpc:          travel.NewTravel(zrpc.MustNewClient(c.TravelRpcConf)),
		HomestayOrderModel: model.NewHomestayOrderModel(sqlx.NewMysql(c.DB.DataSource), c.Cache),
	}
}
