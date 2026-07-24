FROM golang:1.26.4-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/app/main.go

EXPOSE 8000

COPY config.yml .

CMD ["./main"]