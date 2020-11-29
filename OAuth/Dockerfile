FROM golang:latest

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

ENV PORT=8083

RUN go build -o main .

CMD ["/app/main"]