FROM golang:latest

WORKDIR /app

COPY ./Comment-Service .

RUN go mod download

RUN go build .

CMD ["./commentservice"]