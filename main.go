package main

import (
	"context"
	"github.com/sirupsen/logrus"
)

var consuladdr []string = []string{"127.0.0.1:8500"}

func main() {
	consul, err := NewConsul(consuladdr, nil)
	if err != nil {
		logrus.Error(err)
		return
	}
	change := make(chan Config)
	defer close(change)
	config := new(LogConfig)
	key := "rpc/config.yaml"
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		err = RegisterCenterMonitor(consul, key, config, ctx, cancel, change)
	}()
	if err != nil {
		cancel()
		logrus.Error(err)
		return
	}
BreakPoint:
	select {
	case <-ctx.Done():
		return
	case conf := <-change:
		if conf == nil {
			break BreakPoint
		}
		logrus.Println(conf)
	}
}
