GOPATH:=$(shell go env GOPATH)

.PHONY: init
init:
	@go get -u google.golang.org/protobuf/proto
	@go install github.com/golang/protobuf/protoc-gen-go@latest
	@go install github.com/asim/go-micro/cmd/protoc-gen-micro/v3@latest

.PHONY: proto
proto:
	@protoc --proto_path=. --micro_out=. --go_out=:. proto/user-center.proto

.PHONY: tidy
tidy:
	@go mod tidy

.PHONY: build
build:
	@go build -o user-center *.go

.PHONY: test
test:
	@go test -v ./... -cover

.PHONY: docker
docker:
	@docker build -t user-center:latest .
