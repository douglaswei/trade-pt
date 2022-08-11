FROM golang:alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /go/src/item-center
RUN apk --update --no-cache add ca-certificates gcc libtool make musl-dev protoc
COPY item-center /go/src/item-center
RUN make init proto tidy build

FROM scratch
COPY --from=builder /etc/ssl/certs /etc/ssl/certs
COPY --from=builder /go/src/item-center/item-center /item-center
ENTRYPOINT ["/item-center"]
CMD []
