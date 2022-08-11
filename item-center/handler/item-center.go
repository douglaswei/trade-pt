package handler

import (
	"context"
	"io"
	"time"

	log "github.com/asim/go-micro/v3/logger"

	pb "github.com/douglaswei/trade-pt/item-center/proto"
)

type ItemCenter struct{}

func (e *ItemCenter) Call(ctx context.Context, req *pb.CallRequest, rsp *pb.CallResponse) error {
	log.Infof("Received ItemCenter.Call request: %v", req)
	rsp.Msg = "Hello " + req.Name
	return nil
}

func (e *ItemCenter) ClientStream(ctx context.Context, stream pb.ItemCenter_ClientStreamStream) error {
	var count int64
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			log.Infof("Got %v pings total", count)
			return stream.SendMsg(&pb.ClientStreamResponse{Count: count})
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		count++
	}
}

func (e *ItemCenter) ServerStream(ctx context.Context, req *pb.ServerStreamRequest, stream pb.ItemCenter_ServerStreamStream) error {
	log.Infof("Received ItemCenter.ServerStream request: %v", req)
	for i := 0; i < int(req.Count); i++ {
		log.Infof("Sending %d", i)
		if err := stream.Send(&pb.ServerStreamResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
		time.Sleep(time.Millisecond * 250)
	}
	return nil
}

func (e *ItemCenter) BidiStream(ctx context.Context, stream pb.ItemCenter_BidiStreamStream) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Infof("Got ping %v", req.Stroke)
		if err := stream.Send(&pb.BidiStreamResponse{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
