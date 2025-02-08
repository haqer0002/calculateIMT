FROM golang:alpine

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o calculateIMT main.go

CMD ["./calculateIMT"]