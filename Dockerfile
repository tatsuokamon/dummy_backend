FROM golang:1.26 AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod tidy
COPY main.go ./

RUN go build -o app main.go

FROM gcr.io/distroless/cc-debian12
WORKDIR /app

COPY --from=builder /app/app .
EXPOSE 80

# CMD ["./app"]
