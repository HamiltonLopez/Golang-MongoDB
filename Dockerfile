# Etapa 1: build
FROM golang:1.24.2 AS builder

WORKDIR /app

# Copia go.mod y go.sum para descargar dependencias
COPY go.mod go.sum ./
RUN go mod download

# Asegúrate de que las dependencias estén actualizadas
RUN go mod tidy

# Copia el resto de los archivos
COPY . .

# Verifica los archivos copiados
RUN echo "Archivos en /app:" && ls -la /app

# Construye la aplicación
RUN go build -v -x -o main .

# Etapa 2: run
FROM alpine:latest

# Instala los certificados SSL
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copia el binario desde la etapa de construcción
COPY --from=builder /app/main .

# Expone el puerto
EXPOSE 8080

# Ejecuta la aplicación
CMD ["./main"]
