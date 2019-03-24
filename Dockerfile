FROM golang:1.12-alpine

# Install tools required for project
RUN apk add --no-cache bash git make protobuf

ADD . /go/src/github.com/shawnlower/go-ltp

WORKDIR /go/src/github.com/shawnlower/go-ltp

RUN make

FROM golang:1.12-alpine

RUN apk add --no-cache bash libc6-compat

RUN mkdir /app

LABEL maintainer="Shawn Lower <shawn@shawnlower.com>"

ENTRYPOINT ["sh", "-c"]

COPY --from=0 /go/bin /app

CMD ["/app/ltpd --auth-method=insecure --debug"]

EXPOSE 17900
