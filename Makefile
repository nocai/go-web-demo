LDFLAGS += -X "main.Version=$(shell git describe --tags `git rev-list --tags --max-count=1`)"
LDFLAGS += -X "main.GitHash=$(shell git rev-parse HEAD)"
LDFLAGS += -X "main.BuildTime=$(shell date '+%Y-%m-%d %H:%M:%S')"

all: clean build

build:
	GO111MODULE=on go build -ldflags '$(LDFLAGS)'

clean:
	go clean

# Cross compilation
linux:
	GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags '$(LDFLAGS)'
