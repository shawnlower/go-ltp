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

pb_include=-I $(shell go list -f '{{ .Dir }}' -m github.com/golang/protobuf)
pb_include+=-I $(shell go list -f '{{ .Dir }}' -m github.com/grpc-ecosystem/grpc-gateway)/third_party/googleapis
proto: $(obj)
	@echo "Rebuilding protobuf stubs"
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
	@export CGO_ENABLED=0
	@export GOOS=linux
	echo "Building $@"
	go build -i -o ./$@ ./$<

ltpd: bin/ltpd

ltpcli: bin/ltpcli

grpc_gw: bin/grpc_gw
