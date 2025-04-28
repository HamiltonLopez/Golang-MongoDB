# Etapa 1: build
FROM golang:1.20-alpine AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .


RUN go build -o main .


FROM alpine:latest


RUN apk --no-cache add ca-certificates

WORKDIR /root/


COPY --from=builder /app/main .

EXPOSE 8080


CMD ["./main"]


# Etapa 2: production
# FROM alpine:latest
# RUN apk --no-cache add ca-certificates