package handler

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/Gonewithmyself/gomicrorpc/example2/lib"
	"github.com/Gonewithmyself/gomicrorpc/example2/proto"
)

type Say struct{}

var _ proto.SayHandler = (*Say)(nil)

func (s *Say) Hello(ctx context.Context, req *proto.SayParam, rsp *proto.SayResponse) error {
	fmt.Println("received", req.Msg)
	rsp.Header = make(map[string]*proto.Pair)
	rsp.Header["name"] = &proto.Pair{Key: 1, Values: "abc"}

	rsp.Msg = "hello world"
	rsp.Values = append(rsp.Values, "a", "b")
	rsp.Type = proto.RespType_DESCEND

	return nil
}

func (s *Say) MyName(ctx context.Context, req *proto.SayParam, rsp *proto.SayParam) error {
	rsp.Msg = "lp"
	return nil
}

/*
 模拟得到一些数据
*/
func (s *Say) Stream(ctx context.Context, req *proto.SRequest, stream proto.Say_StreamStream) error {

	for i := 0; i < int(req.Count); i++ {
		rsp := &proto.SResponse{}
		for j := lib.Random(3, 5); j < 10; j++ {
			rsp.Value = append(rsp.Value, lib.RandomStr(lib.Random(3, 10)))
		}
		if err := stream.Send(rsp); err != nil {
			return err
		}
		// 模拟处理过程
		time.Sleep(time.Microsecond * 50)
	}
	return nil
}

/*
 模拟数据
*/
func (s *Say) BidirectionalStream(ctx context.Context, stream proto.Say_BidirectionalStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(req.Count)
		if err := stream.Send(&proto.SResponse{Value: []string{lib.RandomStr(lib.Random(3, 6))}}); err != nil {
			return err
		}
	}
	fmt.Println("end BidirectionalStream")
	return nil
}
