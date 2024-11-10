
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
   cd urbverde-ui
   ```

2. Instale as dependências listadas no `go.mod`:
   ```bash
   go mod download
   ```

3. Inicie o servidor:
   ```bash
   go run main.go
   ```

4. Acesse a API localmente em `[http://localhost:8080/api/v1/address/suggestions]http://localhost:8080/api/v1/address/suggestions`.

### Opção 2: Rodar com Docker Compose

1. Certifique-se de ter o Docker e Docker Compose instalados.

2. Para construir e rodar o projeto com Docker Compose:
   ```bash
   docker-compose up --build
   ```

5. A API estará acessível em `http://localhost:8080/api`.

## Rotas Disponíveis

Você pode listar os endpoints em: http://localhost:8080/api/v1/endpoints

- `GET /api/v1/address/suggestions`: Retorna sugestões de endereços.
- ⚙️ **[em desenvolvimento]** `GET /api/v1/city/bounds`: Retorna os dados de localização e código do município.
- ⚙️ **[em desenvolvimento]** `GET /api/v1/data`: Retorna as camadas e categorias disponiveis para o código do municipio.

---

Este README fornece uma visão geral de como rodar a UrbVerde API localmente, seja usando Go diretamente ou com Docker Compose.
