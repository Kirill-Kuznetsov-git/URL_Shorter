FROM golang:1.17.3-alpine3.14 AS builder
ADD app /app
WORKDIR /app
RUN apk add --no-cache ca-certificates git
RUN go mod download
RUN	go build -o main .

FROM alpine:3.10 AS app
WORKDIR /app
COPY --from=builder /app/main /app
ENTRYPOINT ["./main"]
