COMMANDS=ltpcli ltpd grpc_gw
BINARIES=$(addprefix bin/,$(COMMANDS))

# If a GOPATH is defined, use that (e.g. for a Docker build)
# otherwise, build in cwd, under a temporary build/ directory

src = $(wildcard api/proto/*.proto)
obj = $(src:.go=.pb.go)

all: proto binaries

build:
	go build

install:
	go install

proto: $(obj)
	@echo "Rebuilding protobuf stubs"
	@go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	@go get github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
	@go get github.com/golang/protobuf/protoc-gen-go
	@go get github.com/gogo/protobuf/protoc-gen-gofast
	@protoc $< -I.\
	    -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	    -I$$GOPATH/src/github.com/golang/protobuf \
	    -I$$GOPATH/src/github.com/gogo/protobuf/protobuf \
	    -I$$GOPATH/src/github.com/gogo/protobuf/protobuf/ \
	    --grpc-gateway_out=logtostderr=true:. \
	    --go_out=plugins=grpc:.

test:
	go test

clean:
	rm -f $(BINARIES) api/proto/*.go

binaries: $(BINARIES)

bin/%: cmd/%
	@export CGO_ENABLED=0
	@export GOOS=linux
	echo "Building $@"
	go build -i -o ./$@ ./$<

ltpd: bin/ltpd

ltpcli: bin/ltpcli

grpc_gw: bin/grpc_gw
