FROM golang:alpine AS builder
WORKDIR /usr/src/app

COPY main.go .
COPY go.mod .

RUN go mod download
RUN go build -v -o relay

FROM golang:alpine
WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/relay relay
COPY config.json .

EXPOSE 8080
CMD [ "./relay" ]
