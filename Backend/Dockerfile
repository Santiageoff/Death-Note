# Etapa 1: build
FROM golang:1.22 AS builder

WORKDIR /app

# Copiar archivos de dependencias
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copiar el código fuente
COPY . .

# Compilar el binario para Linux sin dependencias C
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# Etapa 2: runtime
FROM alpine:latest

WORKDIR /app

# Instalar curl y bash y descargar wait-for-it
RUN apk --no-cache add curl bash \
    && curl -sSLo /usr/local/bin/wait-for-it https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh \
    && chmod +x /usr/local/bin/wait-for-it

# Copiar binario y variables de entorno
COPY --from=builder /app/main .
COPY --from=builder /app/.env .env

EXPOSE 8080

# Ejecutar backend esperando que MySQL esté listo
CMD ["wait-for-it", "mysql:3306", "--timeout=30", "--", "./main"]
