COMMANDS=ltpcli ltpd grpc_gw
BINARIES=$(addprefix bin/,$(COMMANDS))

src = $(wildcard api/proto/*.proto)
obj = $(src:.go=.pb.go)

all: binaries proto

build:
	go build

install:
	go install

pb_include=-I $(shell go list -f '{{ .Dir }}' -m github.com/golang/protobuf)
pb_include+=-I $(shell go list -f '{{ .Dir }}' -m github.com/gogo/protobuf)/protobuf
pb_include+=-I $(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway)/third_party/googleapis
proto: $(obj)
	@echo "Rebuilding protobuf stubs"
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@go install github.com/golang/protobuf/protoc-gen-go
	@protoc $< -I.\
	    $(pb_include) \
	    --grpc-gateway_out=logtostderr=true,paths=source_relative:. \
	    --go_out=paths=source_relative,plugins=grpc:.

test:
	go test

clean:
	rm -f $(BINARIES) api/proto/*.go

binaries: $(BINARIES)

bin/%: cmd/%
	echo "Building $@"
	go get ./$<
	go build -o ./$@ ./$<

static:
	export CGO_ENABLED=0
	export GOOS=linux
static: binaries

ltpd: bin/ltpd

ltpcli: bin/ltpcli

grpc_gw: bin/grpc_gw
