package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"os/signal"

	"github.com/Gonewithmyself/gomicrorpc/example2/common"
	"github.com/Gonewithmyself/gomicrorpc/example2/lib"
	"github.com/Gonewithmyself/gomicrorpc/example2/proto"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/client"
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
		)
	*/
	service.Init()
	service.Client().Init(client.Retries(3),
		client.PoolSize(5))
	sayClent := proto.NewSayService(common.ServiceName, service.Client())

	SayHello(sayClent)
	NotifyTopic(service)
	GetStreamValues(sayClent)
	TsBidirectionalStream(sayClent)

	st := make(chan os.Signal)
	signal.Notify(st, os.Interrupt)

	<-st
	fmt.Println("server stopped.....")
}

func SayHello(client proto.SayService) {
	rsp, err := client.Hello(context.Background(), &proto.SayParam{Msg: "hello server"})
	if err != nil {
		panic(err)
	}

	fmt.Println(rsp)
}

// test stream
func GetStreamValues(client proto.SayService) {
	rspStream, err := client.Stream(context.Background(), &proto.SRequest{Count: 10})
	if err != nil {
		panic(err)
	}

	idx := 1
	for {
		rsp, err := rspStream.Recv()

		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		fmt.Printf("test stream get idx %d  data  %v\n", idx, rsp)
		idx++
	}
	// close the stream
	if err := rspStream.Close(); err != nil {
		fmt.Println("stream close err:", err)
	}
	fmt.Println("Read Value End")
}

func TsBidirectionalStream(client proto.SayService) {
	rspStream, err := client.BidirectionalStream(context.Background())
	if err != nil {
		panic(err)
	}
	for i := int64(0); i < 7; i++ {
		if err := rspStream.Send(&proto.SRequest{Count: i}); err != nil {
			fmt.Println("send error", err)
			break
		}
		rsp, err := rspStream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}

		fmt.Printf("test stream get idx %d  data  %v\n", i, rsp)
	}
	// close the stream
	if err := rspStream.Close(); err != nil {
		fmt.Println("stream close err:", err)
	}
	fmt.Println("TsBidirectionalStream: Read Value End")
}

func NotifyTopic(service micro.Service) {
	p := micro.NewPublisher(common.Topic1, service.Client())
	p.Publish(context.TODO(), &proto.SayParam{Msg: lib.RandomStr(lib.Random(3, 10))})
}
