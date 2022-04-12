FROM --platform=$BUILDPLATFORM golang:latest AS builder
WORKDIR /tmp
COPY ./signal_sample_server.go /tmp
ARG TARGETOS TARGETARCH
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 go build -a -o signal_sample_server signal_sample_server.go

FROM alpine:latest
WORKDIR /
COPY --from=builder /tmp/signal_sample_server /bin/
CMD ["/bin/signal_sample_server"]
