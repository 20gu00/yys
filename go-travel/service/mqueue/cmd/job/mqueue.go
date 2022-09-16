package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/logx"
	"go-travel/service/mqueue/cmd/job/internal/config"
	"go-travel/service/mqueue/cmd/job/internal/logic"
	"go-travel/service/mqueue/cmd/job/internal/svc"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
)

var configFile = flag.String("f", "etc/mqueue.yaml", "Specify the config file")

func main() {
	flag.Parse()
	var c config.Config

	conf.MustLoad(*configFile, &c, conf.UseEnv())

	// log、prometheus、trace、metricsUrl
	if err := c.SetUp(); err != nil {
		//panic use error here
		//panic(err)
		return
	}

	//logx.DisableStat()

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()
	cronJob := logic.NewCronJob(ctx, svcContext)
	mux := cronJob.Register()

	if err := svcContext.AsynqServer.Run(mux); err != nil {
		logx.WithContext(ctx).Errorf("!!!CronJobErr!!! run err:%+v", err)
		os.Exit(1)
	}
}
