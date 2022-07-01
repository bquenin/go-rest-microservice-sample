FROM golang:alpine as tester

WORKDIR /root

COPY . .

CMD CGO_ENABLED=0 go test ./...
