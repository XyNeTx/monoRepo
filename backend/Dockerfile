
FROM golang:1.25-alpine AS go-builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/cosmtrek/air@latest

EXPOSE 8080
CMD [ "air" ]