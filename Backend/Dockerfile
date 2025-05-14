# Etapa 1: Construcción
FROM golang:1.22 AS builder

WORKDIR /app

# Copia todo el código fuente
COPY . .

# Descarga dependencias y compila
RUN go mod download
RUN go build -o main .

# Etapa 2: Imagen final
FROM debian:bookworm-slim

WORKDIR /app

# Copia el binario desde el builder
COPY --from=builder /app/main .

# ✅ Copiar también el .env desde la etapa anterior
COPY --from=builder /app/.env .

# Comando de inicio
CMD ["./main"]
