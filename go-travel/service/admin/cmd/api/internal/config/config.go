package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis struct {
		Address string
		Pass    string
	}
	AdminRpcConf      zrpc.RpcClientConf
	UsercenterRpcConf zrpc.RpcClientConf
	PaymentRpcConf    zrpc.RpcClientConf
	OrderRpcConf      zrpc.RpcClientConf
	TravelRpcConf     zrpc.RpcClientConf
}
