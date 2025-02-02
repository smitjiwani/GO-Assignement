FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

RUN mkdir -p /app/tmp && chmod -R 777 /app/tmp

COPY go.mod go.sum ./
RUN go mod download

COPY .air.toml .
COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]