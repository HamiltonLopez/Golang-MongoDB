#!/bin/bash

APP_NAME="go-mongo-app"
IMAGE_NAME="go-mongo-app-image"
CONTAINER_NAME="go-mongo-app-container"

echo "📦 Construyendo la imagen Docker..."
docker build -t $IMAGE_NAME .

echo "🧹 Deteniendo y eliminando contenedor anterior si existe..."
docker stop $CONTAINER_NAME 2>/dev/null
docker rm $CONTAINER_NAME 2>/dev/null

echo "🚀 Ejecutando el contenedor..."
docker run -d --name $CONTAINER_NAME -p 8080:8080 --env-file .env $IMAGE_NAME

echo "✅ Contenedor '$CONTAINER_NAME' ejecutándose en http://localhost:8080"
