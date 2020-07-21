package infra

import (
	"github.com/golang/protobuf/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func FailedPrecondition(module, code, message string) *status.Status {
	st := status.New(codes.FailedPrecondition, message)
	messages := []proto.Message{
		&errdetails.PreconditionFailure_Violation{
			Type:        module,
			Subject:     code,
			Description: message,
		},
	}

	var err error
	if st, err = st.WithDetails(messages...); err != nil {
		panic(err)
	}
	return st
}
