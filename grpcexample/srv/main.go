package main

import (
	"context"
	"fmt"

	"github.com/Gonewithmyself/gomicrorpc/grpcexample/common"
	"github.com/Gonewithmyself/gomicrorpc/grpcexample/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/server/grpc"
)

type Say struct{}

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
	// 初始化服务
	service := micro.NewService(
		micro.Server(grpc.NewServer()),
		micro.Name(common.GrpcExampleName),
	)

	service.Init()

	// 注册 Handler
	if err := proto.RegisterSayHandler(service.Server(), new(Say)); err != nil {
		panic(err)
	}

	// run server
	if err := service.Run(); err != nil {
		panic(err)
	}
}
