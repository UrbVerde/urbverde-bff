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

# Set the API host based on environment
RUN if [ "$ENV" = "production" ]; then \
      export API_HOST="api.urbverde.com.br"; \
    else \
      export API_HOST="localhost:8080"; \
    fi && \
    # Update the host in main.go before generating Swagger docs
    sed -i "s|@host.*|@host ${API_HOST}|" main.go && \
    # Generate Swagger docs
    swag init --g main.go --output docs

# Build the application binary
RUN go build -o urbverde-bff

# Start a minimal runtime image
FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/urbverde-bff .
COPY --from=builder /app/docs ./docs
EXPOSE 8080
CMD ["./urbverde-bff"]