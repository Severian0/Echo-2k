# syntax=docker/dockerfile:1.7

################## Front-end ##################
FROM node:22-slim AS frontend_builder

WORKDIR /frontend
ENV NODE_ENV=production

# Install deps
COPY frontend/package*.json ./
RUN npm install --include=dev

# Copy app and build
COPY frontend/ .
RUN npm run build      # -> /frontend/dist

################## Back-end ##################
FROM golang:1.25-alpine AS backend_builder

WORKDIR /app
ENV CGO_ENABLED=0 GOOS=linux

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags="-s -w" -o main ./cmd/api

################## Runtime ##################
FROM alpine:latest AS prod

WORKDIR /
RUN apk add --no-cache ca-certificates curl

COPY --from=backend_builder /app/main /main
COPY --from=frontend_builder /frontend/dist /dist

EXPOSE 8080
ENTRYPOINT ["/main"]
