FROM golang:1.24.0 AS builder

WORKDIR /backendcompile

COPY . .

RUN CGO_ENABLED=0 go build -o mycareerapp-api cmd/api/main.go

FROM alpine:latest AS prod

WORKDIR /build

COPY --from=builder /backendcompile/mycareerapp-api .
COPY .env .
EXPOSE 8080
ENTRYPOINT ["./mycareerapp-api"]