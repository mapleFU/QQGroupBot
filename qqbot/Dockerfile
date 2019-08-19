FROM golang:1.12-alpine

ENV SRC_DIR=/go/src/github.com/mapleFU/QQBot/qqbot
ENV GO111MODULE on

WORKDIR $SRC_DIR

EXPOSE 8085

COPY ./ $SRC_DIR

RUN go build -mod=vendor $SRC_DIR/main.go

CMD ["./main"]