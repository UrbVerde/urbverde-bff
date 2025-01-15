<!-- urbverde-bff/README.md -->
# UrbVerde API

API desenvolvida em Go para fornecer sugestões de endereço e outros serviços relacionados.

## Pré-requisitos

- [Go](https://golang.org/doc/install) 1.16 ou superior
- [Docker](https://www.docker.com/) e [Docker Compose](https://docs.docker.com/compose/)

## Como Rodar a API Localmente

### Opção 1: Rodar Localmente com Go

1. Clone o repositório:
   ```bash
   git clone https://github.com/UrbVerde/urbverde-bff.git
   cd urbverde-bff
   ```

2. Instale as dependências listadas no `go.mod`:
   ```bash
   go mod download
   go mod tidy
   ```

3. Inicie o servidor:
   ```bash
   go run main.go
   ```

4. Acesse a API localmente em `http://localhost:8080/v1/address/suggestions`.

### Opção 2: Rodar com Docker Compose

1. Certifique-se de ter o Docker e Docker Compose instalados.

2. Para construir e rodar o projeto com Docker Compose:
   ```bash
   docker-compose up --build
   ```

3. Para testar o swagger com a api online use:
   ```bash
   ENV=production docker-compose build && docker-compose up -d 
   ```

5. A API estará acessível em `http://localhost:8080/`.

## Documentação Swagger

A API possui documentação interativa utilizando o Swagger, que pode ser gerada e atualizada automaticamente com base nos comentários do código.

### Como acessar a documentação Swagger

1. Certifique-se de que a API esteja rodando localmente.
2. Acesse a interface do Swagger no navegador em:
   [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)
3. Explore os endpoints documentados e teste as funcionalidades diretamente pela interface do Swagger.

### Como atualizar a documentação Swagger

Se houver mudanças nas rotas ou nos controladores, você precisará atualizar a documentação Swagger:

1. Certifique-se de que o **Swag CLI** está instalado. Se ainda não estiver, instale com:
   ```bash
   go install github.com/swaggo/swag/cmd/swag@latest
   ```
2. Gere ou atualize os arquivos de documentação com o comando:
   ```bash
   swag init
   ```
3. Certifique-se de que a pasta `docs/` foi atualizada corretamente. Essa pasta contém os arquivos `swagger.json` e `swagger.yaml` utilizados pelo Swagger.
