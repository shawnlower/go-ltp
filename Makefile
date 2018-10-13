COMMANDS=ltpcli ltpd
BINARIES=$(addprefix bin/,$(COMMANDS))

src = $(wildcard api/*.proto)
obj = $(src:.go=.pb.go)

all: proto binaries

build:
	go build

install:
	go install

proto: $(obj)
	@echo "Rebuilding protobuf stubs"
	@protoc $< --go_out=plugins=grpc:.

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
