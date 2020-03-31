package main

import (
	"center/center"
	"center/config"
	"context"
	"github.com/sirupsen/logrus"
)

var consuladdr []string = []string{"127.0.0.1:8500"}

func main() {
	consul,err := center.NewConsul(consuladdr,nil)
	if err != nil{
		logrus.Error(err)
		return
	}
	config := new(config.RPCConfig)
	key := "rpcx/config.yaml"
	ctx,cancel:=context.WithCancel(context.Background())
	err = center.RegisterCenterMonitor(consul,key,config,ctx,cancel)
	if err!= nil{
		logrus.Error(err)
	}
	select{
	case <-ctx.Done():
		return
	}
}
