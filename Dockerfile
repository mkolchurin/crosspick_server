# syntax=docker/dockerfile:1
FROM golang:alpine AS builder

WORKDIR /build

ADD go.mod .

COPY . .

RUN go build -o crosspicktest main.go

FROM alpine

WORKDIR /build
EXPOSE 9988
COPY --from=builder /build/crosspicktest /build/crosspicktest
COPY www /build/www

CMD ["./crosspicktest"]