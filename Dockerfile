# Dockerfile para rodar os testes e iniciar o serviço

# Etapa de build
FROM golang:1.19-alpine AS builder

WORKDIR /app

# Instalar gcc e musl-dev para rodar os testes e compilar
RUN apk add --no-cache gcc musl-dev

# Copiar os arquivos de dependência
COPY go.mod .
# COPY go.sum .
RUN go mod download

# Copiar o código do projeto
COPY . .

# Rodar os testes
RUN go test -v

# Compilar a aplicação
RUN go build -o main .

# Etapa final para a execução da aplicação
FROM alpine:latest

WORKDIR /app

# Copiar o binário compilado da etapa anterior
COPY --from=builder /app/main .

RUN echo "WEATHER_API_KEY: $WEATHER_API_KEY"

EXPOSE 8080

# Executar o servidor
CMD ["./main"]