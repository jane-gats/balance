FROM golang:1.19.3-alpine3.16 as builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .

RUN go install ./cmd/balance

FROM alpine:3.16.0

COPY --from=builder /go/bin/balance /usr/local/bin/balance
ENTRYPOINT ["/usr/local/bin/balance"]
