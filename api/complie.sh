#!/usr/bin/env sh

export GOPATH=/Users/liujun/go

protoc \
  -I. \
  -I/usr/local/include \
  -I$GOPATH/src  \
  -I$GOPATH/src/github.com/gogo/protobuf/protobuf \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --gogoslick_out=plugins=grpc:. \
  *.proto

# 这个生成grpc-gateway
protoc -I. \
  -I/usr/local/include \
  -I=$GOPATH/src  \
  -I$GOPATH/src/github.com/gogo/protobuf/protobuf \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true,allow_delete_body=true:. \
  *.proto
#
#protoc -I/usr/local/include -I. \
#  -I$GOPATH/src \
#  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
#  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
#  --swagger_out=logtostderr=true,allow_delete_body=true:. \
#  *.proto
#
#
## 把api.swagger.json传到sawgger目录，并重新打包静态文件
#cd ..
#cp -f api/api.swagger.json swagger/
#go-bindata -pkg=swagger -o swagger/swagger.go \
#  swagger/*.png \
#  swagger/*.html \
#  swagger/*.js \
#  swagger/*.css \
#  swagger/*.json