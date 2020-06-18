FROM golang:1.14 AS build

WORKDIR /app

COPY . /app

RUN GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build -o ./bin/iaas ./cmd/iaas

FROM alpine:3

COPY --from=build /app/bin/iaas /app/iaas

WORKDIR /app

CMD ["/app/iaas"]
