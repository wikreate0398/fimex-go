FROM golang:1.23.4 as builder

RUN go version

COPY . /github.com/wikreate0398/fimex-go/
WORKDIR /github.com/wikreate0398/fimex-go/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/wikreate0398/fimex-go/.bin/app .

CMD ["./app"]