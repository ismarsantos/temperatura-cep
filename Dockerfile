# Dockerfile para rodar os testes e iniciar o serviço
FROM golang:1.19-alpine AS builder

WORKDIR /app

# Instalar gcc e musl-dev para rodar os testes e compilar
RUN apk add --no-cache gcc musl-dev

# Copiar os arquivos de dependência
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copiar o código do projeto
COPY . .

# Rodar os testes
RUN go test -v

# Compilar a aplicação
RUN go build -o /app/main .

# Etapa final
FROM alpine:latest

WORKDIR /app

# Copiar o binário compilado da etapa anterior
COPY --from=builder /app/main .

# Copiar o arquivo .env para a imagem final
COPY .env .

# Instalar dependências para rodar a aplicação
RUN apk add --no-cache ca-certificates

# Expor a porta que será usada pelo serviço
EXPOSE 8080

# Executar o servidor
CMD ["/bin/sh", "-c", "./main"]
