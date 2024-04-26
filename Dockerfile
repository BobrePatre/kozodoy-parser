FROM golang:latest AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o parserapp ./cmd/server

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/parserapp .

EXPOSE 8080

CMD ["./parserapp"]
