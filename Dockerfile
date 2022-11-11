FROM golang:1.18.8-alpine3.16 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN go build -o go-todo .

FROM alpine:latest

WORKDIR /

COPY --from=builder /app/go-todo ./

EXPOSE 8000

ENV DATABASE_URL postgres://username:password@localhost:5432/todo?sslmode=disable

CMD ["./go-todo"]
