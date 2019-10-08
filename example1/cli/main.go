package main

import (
	"context"
	"fmt"

	"github.com/Gonewithmyself/gomicrorpc/example1/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
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
		micro.Registry(reg),
	)

	// 2019年源码有变动默认使用的是mdns面不是consul了
	// 如果你用的是默认的注册方式把上面的注释掉用下面的
	/*
		// 初始化服务
		service := micro.NewService(
			micro.Name("lp.srv.eg1"),
		)
	*/
	service.Init()

	sayClent := proto.NewSayService("lp.srv.eg1", service.Client())

	rsp, err := sayClent.Hello(context.Background(), &proto.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)

}
