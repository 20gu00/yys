package svc

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-travel/service/travel/cmd/rpc/internal/config"
	"go-travel/service/travel/model"
)

type ServiceContext struct {
	Config config.Config

	HomestayModel         model.HomestayModel
	HomestayActivityModel model.HomestayActivityModel
	HomestayBusinessModel model.HomestayBusinessModel
	HomestayCommentModel  model.HomestayCommentModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.DB.DataSource)
	return &ServiceContext{
		Config: c,

		HomestayModel:         model.NewHomestayModel(sqlConn, c.Cache),
		HomestayActivityModel: model.NewHomestayActivityModel(sqlConn, c.Cache),
		HomestayBusinessModel: model.NewHomestayBusinessModel(sqlConn, c.Cache),
		HomestayCommentModel:  model.NewHomestayCommentModel(sqlConn, c.Cache),
	}
}
