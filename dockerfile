FROM golang:latest

WORKDIR /EOP

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD ["go", "run", "main.go"]
