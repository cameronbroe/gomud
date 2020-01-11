FROM golang:1.13 as builder

WORKDIR /go/src/github.com/cameronbroe/gomud

ENV CGO_ENABLED=0 \
    GOARCH=amd64 \
    GOOS=linux

COPY . .

RUN go mod download \
    && mkdir -p build/linux/ \
    && go build -o build/linux/gomud ./cmd/gomud

FROM alpine:latest

RUN apk --update add ca-certificates \
    && apk add sqlite \
    && apk add socat

COPY --from=builder /go/src/github.com/cameronbroe/gomud/build/linux/gomud /usr/bin

CMD ["gomud"]