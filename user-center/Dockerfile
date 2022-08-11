FROM golang:alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/user-center
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc
COPY . /go/src/user-center
RUN make init proto tidy build

FROM scratch
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/user-center/user-center /user-center
ENTRYPOINT ["/user-center"]
CMD []
