FROM golang:alpine AS builder

WORKDIR /build

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

FROM alpine:3.11

COPY --from=builder /build/main /

ENTRYPOINT ["/main"]
