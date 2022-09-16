package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-travel/service/payment/cmd/rpc/internal/config"
	"go-travel/service/payment/model"

	"github.com/zeromicro/go-queue/kq"
)

type ServiceContext struct {
	Config                             config.Config
	ThirdPaymentModel                  model.ThirdPaymentModel
	KqueuePaymentUpdatePayStatusClient *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {
	//1
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,

		//2
		ThirdPaymentModel:                  model.NewThirdPaymentModel(sqlConn, c.Cache),
		KqueuePaymentUpdatePayStatusClient: kq.NewPusher(c.KqPaymentUpdatePayStatusConf.Brokers, c.KqPaymentUpdatePayStatusConf.Topic),
		//no 1 2 -> rpc error: code = Internal desc = panic: runtime error: invalid memory address or nil pointer dereference
	}
}
