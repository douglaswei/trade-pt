package main

import (
	"user-center/handler"
	pb "user-center/proto"

	"github.com/asim/go-micro/v3"
	log "github.com/asim/go-micro/v3/logger"
)

var (
	service = "user-center"
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
	pb.RegisterUserCenterHandler(srv.Server(), new(handler.UserCenter))

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}
