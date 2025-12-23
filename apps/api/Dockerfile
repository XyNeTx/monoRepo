
FROM golang:1.25-alpine AS go-builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN go install github.com/air-verse/air@latest

WORKDIR /app/cmd/server

RUN mv .air_linux.toml .air.toml

EXPOSE 8080

CMD [ "air" ]