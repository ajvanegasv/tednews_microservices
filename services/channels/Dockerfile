FROM golang:1.22-alpine3.19 as builder

WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /app/main .

FROM alpine:3.19

WORKDIR /

RUN set -x && apk update && apk add --no-cache ca-certificates && rm -rf /var/cache/apk/*

COPY --from=builder /app/main /app/main

ENV GIN_MODE release
EXPOSE 8080

WORKDIR /app

# Run the Go Gin binary.
ENTRYPOINT ["/app/main"]