# urbverde-bff/Dockerfile
FROM golang:1.23-alpine AS builder

# Install tools and dependencies
RUN apk add --no-cache git bash

# Install swag CLI
RUN go install github.com/swaggo/swag/cmd/swag@latest

# Set the environment argument (default: develop)
ARG ENV=develop
ENV ENV=${ENV}

# Define the working directory in the container
WORKDIR /app

# Copy the source code and dependencies
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Set the API host based on environment and generate docs
RUN if [ "$ENV" = "production" ]; then \
      export API_HOST="api.urbverde.com.br"; \
    else \
      export API_HOST="localhost:8080"; \
    fi && \
    sed -i "s|@host.*|@host ${API_HOST}|" main.go && \
    swag init --g main.go --output docs

# Build the application binary
RUN go build -o urbverde-bff

# Start a minimal runtime image
FROM alpine:latest
WORKDIR /app

# Copy the binary and docs
COPY --from=builder /app/urbverde-bff .
COPY --from=builder /app/docs /app/docs_source
COPY --from=builder /app/repositories/address/data /app/repositories/address/data
COPY --from=builder /app/repositories/categories/data /app/repositories/categories/data

# Ensure directories exist and have correct permissions
RUN mkdir -p /app/docs && \
    chmod -R 777 /app/docs && \
    chmod -R 777 /app/docs_source && \
    chmod -R 755 /app/repositories

EXPOSE 8080

# Copy docs to mounted volume on container start
CMD cp -r /app/docs_source/* /app/docs/ && ./urbverde-bff