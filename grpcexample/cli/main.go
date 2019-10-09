package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Gonewithmyself/gomicrorpc/grpcexample/common"
	"github.com/Gonewithmyself/gomicrorpc/grpcexample/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client/grpc"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

var etdhosts = []string{
	"http://localhost:32773", "http://localhost:32771", "http://localhost:32769",
}

func main() {
	// 初始化服务

	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etdhosts
	})

	service := micro.NewService(
		micro.Client(grpc.NewClient()),
		micro.Registry(reg),
	)
	service.Init()
	service.Options().Registry.Init(func(options *registry.Options) {
		options.Timeout = time.Second * 2
	})

	sayClent := proto.NewSayService(common.GrpcExampleName, service.Client())

	rsp, err := sayClent.Hello(context.Background(), &proto.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}
