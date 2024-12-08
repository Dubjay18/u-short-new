FROM golang:1.23-alpine AS build

WORKDIR /app
RUN  go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main cmd/api/main.go


FROM golang:1.23-alpine AS dev
WORKDIR /app
RUN  go install github.com/air-verse/air@latest
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main cmd/api/main.go
EXPOSE ${PORT}
CMD ["air","-c",".air.toml"]

FROM alpine:3.20.1 AS prod
WORKDIR /app

COPY --from=build /app/main /app/main
EXPOSE ${PORT}
CMD ["./main"]


FROM node:20 AS frontend_builder
WORKDIR /frontend

COPY frontend/package*.json ./
RUN npm install
COPY frontend/. .
RUN npm run build

FROM node:23-slim AS frontend
RUN npm install -g serve
COPY --from=frontend_builder /frontend/dist /app/dist
EXPOSE 5173
CMD ["serve", "-s", "/app/dist", "-l", "5173"]
