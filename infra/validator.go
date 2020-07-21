package infra

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidateRequest(req interface{}) error {
	if err := validate.Struct(req); err != nil {
		return status.Error(codes.InvalidArgument, err.Error())
	}
	return nil
}
