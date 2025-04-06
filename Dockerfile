# Etapa 1: build
FROM golang:1.20-alpine AS builder

WORKDIR /app

# Descargamos dependencias primero (mejora el cache)
COPY go.mod go.sum ./
RUN go mod download

# Copiamos el resto del código
COPY . .

# Compilamos la app
RUN go build -o main .

# Etapa 2: imagen final minimalista
FROM alpine:latest

# Instala ca-certificates si haces llamadas HTTPs
RUN apk --no-cache add ca-certificates

# Crea un directorio de trabajo
WORKDIR /root/

# Copia el binario desde el build stage
COPY --from=builder /app/main .

EXPOSE 8080

# Comando por defecto
CMD ["./main"]
