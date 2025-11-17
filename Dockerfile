# syntax=docker/dockerfile:1.7         # Enable BuildKit features

################## Front-end ##################
FROM node:22.15.0-alpine AS frontend_builder

WORKDIR /frontend
ENV NODE_ENV=production

# ----- deps (cached) -----
COPY frontend/package*.json ./
RUN --mount=type=cache,target=/root/.npm \
    npm ci --omit=dev

# ----- sources -----
COPY frontend/ .
RUN npm run build      # -> /frontend/dist

# 2) BACKEND BUILD – static, stripped
FROM golang:1.24.2-alpine AS backend_builder
WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod go mod download
COPY . .
# strip debug info with -s -w
RUN --mount=type=cache,target=/root/.cache/go-build \
    go build -trimpath -ldflags="-s -w" -o main ./cmd/api

# 3) RUNTIME – scratch for zero libc overhead
FROM scratch AS prod
# copy only the binary and assets
COPY --from=backend_builder /app/main /main
COPY --from=frontend_builder /frontend/dist /dist
EXPOSE 8080

# directly exec the static binary
ENTRYPOINT ["/main"]