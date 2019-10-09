package subscriber

import (
	"context"
	"fmt"

	"github.com/Gonewithmyself/gomicrorpc/example2/proto"
)

func Handler(ctx context.Context, msg *proto.SayParam) error {
	fmt.Printf("Received message: %s \n", msg.Msg)
	return nil
}
