FROM golang:1.12-alpine

# Install tools required for project
RUN apk add --no-cache git make protobuf

COPY . build

WORKDIR ./build

RUN make

FROM golang:1.12-alpine

#RUN apk add --no-cache bash libc6-compat
RUN apk add --no-cache bash

RUN mkdir /app

LABEL maintainer="Shawn Lower <shawn@shawnlower.com>"

ENTRYPOINT ["sh", "-c"]

COPY --from=0 /go/bin /app

CMD ["/app/ltpd --auth-method=insecure --debug"]

EXPOSE 17900
