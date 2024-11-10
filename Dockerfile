FROM golang:1.23-alpine AS builder

# Define o diretório de trabalho no container
WORKDIR /app

# Copia os arquivos go.mod e go.sum e instala as dependências
COPY go.mod go.sum ./
RUN go mod download

# Copia o código-fonte para o diretório de trabalho
COPY . .

# Compila o binário da aplicação
RUN go build -o urbverde-bff

# Etapa de execução
FROM alpine:latest

# Define o diretório de trabalho
WORKDIR /app

# Copia o binário compilado da etapa anterior
COPY --from=builder /app/urbverde-bff .

# Exponha a porta que a aplicação usa
EXPOSE 8080

# Comando para rodar a aplicação
CMD ["./urbverde-bff"]
