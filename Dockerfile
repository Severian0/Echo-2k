# syntax=docker/dockerfile:1.7

################## Front-end ##################
FROM node:lts-alpine AS frontend_builder

WORKDIR /frontend
ENV NODE_ENV=production

# 1) Install deps using cached layer
COPY frontend/package*.json ./
RUN npm install --include=dev

# 2) Copy the rest and build
COPY frontend/ .
RUN npm run build        # writes static files to /frontend/dist


################## Back-end ##################
FROM golang:alpine AS backend_builder

WORKDIR /app

# Build a static-ish-ish binary (no CGO) for the container OS/arch
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags="-s -w" -o main ./cmd/api


################## Runtime ##################
FROM alpine:latest AS prod

# curl = for your docker-compose healthcheck
# ca-certificates = for any HTTPS calls the Go app might make
RUN apk add --no-cache curl ca-certificates

WORKDIR /

# Copy Go binary and static frontend
COPY --from=backend_builder /app/main /main
COPY --from=frontend_builder /frontend/dist /dist

EXPOSE 8080

ENTRYPOINT ["/main"]
