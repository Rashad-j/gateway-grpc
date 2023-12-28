FROM golang:1.21

WORKDIR /app

COPY . .

RUN go build -o bin/server cmd/server/main.go

EXPOSE 8083

CMD ["./bin/server"]