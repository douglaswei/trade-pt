module github.com/douglaswei/trade-pt/user

go 1.16

require (
	common v0.0.0-00010101000000-000000000000
	github.com/zeromicro/go-zero v1.4.0
	github.com/zeromicro/zero-contrib/zrpc/registry/consul v0.0.0-20220708020647-d13169872f75
	google.golang.org/grpc v1.48.0
	google.golang.org/protobuf v1.28.1
	gorm.io/driver/sqlite v1.3.6
	gorm.io/gen v0.3.14
	gorm.io/gorm v1.23.9-0.20220713102635-3262daf8d468
	gorm.io/plugin/dbresolver v1.2.2
)

replace (
	common => ./../common
	github.com/douglaswei/trade-pt/user => ./
)
