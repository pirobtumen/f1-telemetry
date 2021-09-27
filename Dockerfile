FROM golang:1.16-alpine AS base

WORKDIR /f1-telemetry
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
CMD [ "go", "run", "main.go" ]
