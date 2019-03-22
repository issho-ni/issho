//go:generate protoc --gogofaster_out=plugins=grpc,paths=source_relative:.. -I=$GOPATH/pkg/mod -I.. common/common.proto

package common
