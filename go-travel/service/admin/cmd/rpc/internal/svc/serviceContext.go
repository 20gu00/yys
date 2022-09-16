package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-travel/service/admin/cmd/rpc/internal/config"
	"go-travel/service/admin/model"
)

type ServiceContext struct {
	Config      config.Config
	RedisClient *redis.Redis
	AdminModel  model.AdminModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	sqlConn := sqlx.NewMysql(c.Mysql.Datasource)
	return &ServiceContext{
		Config: c,
		RedisClient: redis.New(c.Redis.Host, func(r *redis.Redis) {
			r.Type = c.Redis.Type
			r.Pass = c.Redis.Pass
		}),
		//AdminModel: model.NewAdminModel(sqlConn, c.Cache),
		AdminModel: model.NewAdminModel(sqlConn),
	}
}
