FROM golang:alpine as builder
WORKDIR /work

COPY . .

RUN CGO_ENABLED=0 go build -o microservice ./cmd/microservice/main.go

FROM alpine
WORKDIR /bin

COPY --from=builder /work/microservice /bin/microservice

CMD /bin/microservice
