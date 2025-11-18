# syntax=docker/dockerfile:1.7

################## Front-end (SvelteKit static) ##################
FROM node:22-slim AS frontend_builder

WORKDIR /frontend

# 1) Install deps (including devDeps so Vite/SvelteKit are available)
COPY frontend/package*.json ./
RUN npm ci --include=dev || npm install --include=dev

# 2) Copy source and build
COPY frontend/ .
# Assumes adapter-static outputs to "dist"
RUN npm run build      # -> /frontend/dist


################## Back-end (Go API) ##################
FROM golang:1.25-alpine AS backend_builder

WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64

# 1) Go deps
COPY go.mod go.sum ./
RUN go mod download

# 2) Copy sources and build
COPY . .
RUN go build -trimpath -ldflags="-s -w" -o main ./cmd/api


################## Runtime ##################
FROM alpine:latest AS prod

# Need curl + sh for your docker-compose healthcheck:
#   test: ["CMD-SHELL", "curl -f http://localhost:${PORT}/ || exit 1"]
RUN apk add --no-cache curl

# Copy the Go binary and static frontend
COPY --from=backend_builder /app/main /main
COPY --from=frontend_builder /frontend/dist /dist

EXPOSE 8080
ENTRYPOINT ["/main"]
