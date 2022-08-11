package main

import (
    "github.com/douglaswei/trade-pt/item-center/handler"
    pb "github.com/douglaswei/trade-pt/item-center/proto"

    log "github.com/asim/go-micro/v3/logger"
)

var (
	service = "item-center"
	version = "latest"
)

func main() {
	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterItemCenterHandler(srv.Server(), new(handler.ItemCenter))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
