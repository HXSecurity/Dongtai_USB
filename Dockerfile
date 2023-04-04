FROM golang:alpine as builder

WORKDIR /opt
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -o usb .

FROM alpine:latest

WORKDIR /opt/

COPY --from=0 /opt/usb ./

EXPOSE 5005
ENTRYPOINT ./usb