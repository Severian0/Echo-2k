# syntax=docker/dockerfile:1.7

################## Front-end ##################
FROM node:lts-alpine AS frontend_builder

WORKDIR /frontend
ENV NODE_ENV=production

# 1) Install deps using cached layer
COPY frontend/package*.json ./
RUN npm install --include=dev

# 1.5) cache-buster for frontend build
ARG FRONTEND_BUILD_ID
RUN echo "frontend build id: $FRONTEND_BUILD_ID"

# 2) Copy the rest and build
COPY frontend/ .
RUN npm run build        # writes static files to /frontend/dist


################## Back-end ##################
FROM golang:alpine AS backend_builder

WORKDIR /app
ENV CGO_ENABLED=0

COPY go.mod go.sum ./
RUN go mod download

# 1.5) cache-buster for backend build
ARG BACKEND_BUILD_ID
RUN echo "backend build id: $BACKEND_BUILD_ID"

COPY . .
RUN go build -trimpath -ldflags="-s -w" -o main ./cmd/api


################## Runtime ##################
FROM alpine:latest AS prod

RUN apk add --no-cache curl ca-certificates

WORKDIR /

COPY --from=backend_builder /app/main /main
COPY --from=frontend_builder /frontend/dist /dist

EXPOSE 8080
ENTRYPOINT ["/main"]
