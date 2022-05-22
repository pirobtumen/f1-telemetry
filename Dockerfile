FROM golang:1.17-alpine AS base
WORKDIR /f1-telemetry

RUN go install github.com/cosmtrek/air@latest

COPY go.mod .
COPY go.sum .
RUN go mod download

CMD [ "air", "-c", "air.toml" ]


FROM base AS prod
COPY . .
CMD [ "go", "run", "main.go" ]
