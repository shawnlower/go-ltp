COMMANDS=ltpcli ltpd grpc_gw
BINARIES=$(addprefix bin/,$(COMMANDS))

src = $(wildcard api/proto/*.proto)
obj = $(src:.go=.pb.go)

all: proto binaries

build:
	go build

install:
	go install

proto: $(obj)
	@echo "Rebuilding protobuf stubs"
	@protoc $< -I. -I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
	    --grpc-gateway_out=logtostderr=true:. \
	    --go_out=plugins=grpc:.

test:
	go test

clean:
	rm -f $(BINARIES)

binaries: $(BINARIES)

bin/%: cmd/%
	@echo "Building $@"
	go build -o ./$@ ./$<

ltpd: bin/ltpd

ltpcli: bin/ltpcli

grpc_gw: bin/grpc_gw
