FROM golang:1.18.1-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o main main.go

CMD ["/app/main"]