COMMANDS=ltpcli ltpd
BINARIES=$(addprefix bin/,$(COMMANDS))

all: binaries

build:
	go build

install:
	go install

test:
	go test

clean:
	rm -f $(BINARIES)

binaries: $(BINARIES)
	@echo "Building $@"

bin/%: cmd/%
	go build -o ./$@ ./$<
