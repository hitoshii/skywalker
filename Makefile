
all: get skywalker forctl

skywalker: src/skywalker/skywalker.go
	GOPATH=`pwd` go build -o bin/skywalker $?

forctl: src/forctl/forctl.go
	GOPATH=`pwd` go build -o bin/forctl $?

get:
	GOPATH=`pwd` go get --fix skywalker
	GOPATH=`pwd` go get --fix forctl

fmt:
	find ./src/forctl ./src/skywalker -name '*.go'|xargs gofmt -w

proto:
	PATH="$(PATH):./bin" GOPATH=`pwd` protoc --go_out=./ ./src/skywalker/message/*.proto

proto-gen:
	GOPATH=`pwd` go get -u github.com/golang/protobuf/protoc-gen-go

clean: bin/skywalker bin/forctl
	rm -rf $?
