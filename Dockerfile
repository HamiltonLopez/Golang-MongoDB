# Etapa 1: build
FROM golang:1.24.2 AS builder

WORKDIR /app


COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN echo "Archivos en /app:" && ls -la /app


RUN go build -v -x -o main .


FROM alpine:latest


RUN apk --no-cache add ca-certificates

WORKDIR /root/


COPY --from=builder /app/main .

EXPOSE 8080


CMD ["./main"]

# Etapa 2: run


