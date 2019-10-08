package main

import (
	"context"
	"fmt"

	"github.com/Gonewithmyself/gomicrorpc/example1/proto"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
)

type Say struct{}

var etdhosts = []string{
	"http://localhost:32773", "http://localhost:32771", "http://localhost:32769",
}

func (s *Say) Hello(ctx context.Context, req *proto.SayParam, rsp *proto.SayResponse) error {
	fmt.Println("received", req.Msg)
	rsp.Header = make(map[string]*proto.Pair)
	rsp.Header["name"] = &proto.Pair{Key: 1, Values: "abc"}

	rsp.Msg = "hello world"
	rsp.Values = append(rsp.Values, "a", "b")
	rsp.Type = proto.RespType_DESCEND

	return nil
}

func main() {
	// 我这里用的etcd 做为服务发现
	reg := etcdv3.NewRegistry(func(op *registry.Options) {
		op.Addrs = etdhosts
	})

	// 初始化服务
	service := micro.NewService(
		micro.Name("lp.srv.eg1"),
		micro.Registry(reg),
	)
	service.Init()

	// 注册 Handler
	proto.RegisterSayHandler(service.Server(), new(Say))

	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
