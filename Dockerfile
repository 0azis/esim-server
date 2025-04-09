ARG GO_VERSION
FROM golang:${GO_VERSION}

COPY . /app
WORKDIR /app/cmd/

RUN go mod download

RUN go build -o ./server

CMD ["./server"]
