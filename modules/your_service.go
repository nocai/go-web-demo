package modules

import (
	"context"
	"github.com/golang/glog"
	"github.com/nocai/go-web-demo/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type service struct {
}

// NewYourServiceServer new a service server
func NewYourServiceServer() api.YourServiceServer {
	return &service{}
}

func (*service) Echo(ctx context.Context, req *api.StringMessage) (*api.StringMessage, error) {
	glog.Info(req)

	return &api.StringMessage{
		Value: req.Value,
	}, status.Error(codes.Unknown, "这是一个错误，出错啦")
}
