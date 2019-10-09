package main

import (
	"time"

	"github.com/Gonewithmyself/gomicrorpc/example2/common"
	"github.com/Gonewithmyself/gomicrorpc/example2/handler"
	"github.com/Gonewithmyself/gomicrorpc/example2/proto"
	"github.com/Gonewithmyself/gomicrorpc/example2/subscriber"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/server"
	"github.com/micro/go-plugins/registry/etcdv3"
)

var etdhosts = []string{
	"http://localhost:32773", "http://localhost:32771", "http://localhost:32769",
}

func main() {
	// 我这里用的etcd 做为服务发现
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etdhosts
	})

	// 初始化服务
	service := micro.NewService(
		micro.Name(common.ServiceName),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*20),
		micro.Registry(reg),
	)

	// 2019年源码有变动默认使用的是mdns面不是consul了
	// 如果你用的是默认的注册方式把上面的注释掉用下面的
	/*
		// 初始化服务
		service := micro.NewService(
			micro.Name(common.ServiceName),
			micro.RegisterTTL(time.Second*30),
			micro.RegisterInterval(time.Second*20),
		)
	*/

	service.Init()
	// 注册 Handler
	proto.RegisterSayHandler(service.Server(), new(handler.Say))

	// Register Subscribers
	if err := server.Subscribe(server.NewSubscriber(common.Topic1, subscriber.Handler)); err != nil {
		panic(err)
	}

	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
