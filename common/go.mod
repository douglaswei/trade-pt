module github.com/douglaswei/trade-pt/common

go 1.16

require (
	gorm.io/driver/mysql v1.3.5
	gorm.io/gorm v1.23.8
)

// This can be removed once etcd becomes go gettable, version 3.4 and 3.5 is not,
// see https://github.com/etcd-io/etcd/issues/11154 and https://github.com/etcd-io/etcd/issues/11931.
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace github.com/douglaswei/trade-pt/common => ./
