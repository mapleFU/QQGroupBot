FROM golang:1.10-alpine

ENV SRC_DIR=/go/src/github.com/mapleFU/QQBot/qqbot

WORKDIR $SRC_DIR

EXPOSE 8085

COPY ./ $SRC_DIR

RUN go build $SRC_DIR/main.go

CMD ["./main"]