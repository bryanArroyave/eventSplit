

# Usa la versión correcta de Go
FROM golang:1.23.6 AS builder

WORKDIR /app

RUN ls -l 

# Copiar archivos de módulos
COPY go.mod go.sum ./
RUN go mod tidy

# Copiar todo el código fuente
COPY . .


# Compilar la aplicación
RUN go build -o journey-service ./cmd/main.go


# Etapa final (imagen más ligera)
FROM golang:1.23.6 AS final
WORKDIR /app
COPY --from=builder /app/journey-service .

# Ejecutar el binario
CMD ["./journey-service"]


# FROM golang:1.23.6 AS builder
# WORKDIR /app
# COPY . .

# RUN GOOS=linux GOARCH=arm64 go build -tags musl -tags dynamic -o main cmd/main.go

# FROM golang:1.23-alpine3.20@sha256:ac67716dd016429be8d4c2c53a248d7bcdf06d34127d3dc451bda6aa5a87bc06

# RUN apk update && apk add build-base

# RUN addgroup -S appgroup && adduser -S appuser -G appgroup --uid 10001

# WORKDIR /app
# USER appuser
# EXPOSE 8080

# CMD ["./main"]
