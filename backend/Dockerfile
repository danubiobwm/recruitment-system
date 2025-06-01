FROM golang:1.24

WORKDIR /app

COPY . .

RUN chmod +x wait-for-it.sh

RUN go mod download

CMD ["go", "run", "main.go"]
