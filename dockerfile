FROM golang:1.21.0 

WORKDIR /EOP

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o eopapp main.go

CMD [ "./eopapp" ]