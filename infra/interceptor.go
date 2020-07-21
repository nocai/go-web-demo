package infra

import (
	"context"
	"gopkg.in/go-playground/validator.v9"
	"runtime/debug"
	"time"

	"github.com/golang/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)


func LoggingInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	glog.Info("==============================START==============================")
	glog.Infof("%s", info.FullMethod)
	glog.Infof("req = { %v }", req)

	var (
		resp interface{}
		err  error
	)
	defer func(start time.Time) {
		glog.Infof("%s, spent time: %v", info.FullMethod, time.Since(start))
		if err != nil {
			glog.Errorf("err = %v", err)
		}
		glog.Infof("resp = { %v }", resp)
		glog.Info("===============================END===============================")
	}(time.Now())

	resp, err = handler(ctx, req)
	if err != nil {
		return resp, err
	}

	return resp, err
}

func RecoveryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()

	return handler(ctx, req)
}

var validate = validator.New()

func ValidateInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	var (
		resp interface{}
		err  error
	)

	if err = validate.Struct(req); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	resp, err = handler(ctx, req)
	if err != nil {
		return resp, err
	}

	if err = validate.Struct(resp); err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}

	return resp, err
}
