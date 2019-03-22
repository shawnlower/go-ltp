FROM golang:1.12-alpine

# Install tools required for project
RUN apk add --no-cache git

RUN apk add --no-cache bash

ADD . /go

RUN mkdir /app

WORKDIR /go/cmd/ltpd
RUN go build -v -o /app/bin/ltpd

WORKDIR /go/cmd/ltpcli
RUN go build -v -o /app/bin/ltpcli

RUN rm -rf /go

WORKDIR /app

ENTRYPOINT ["sh", "-c"]

CMD ["/app/bin/ltpd --auth-method=insecure --debug"]

EXPOSE 17900
