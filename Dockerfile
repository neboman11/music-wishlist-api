# syntax=docker/dockerfile:1

FROM golang:latest

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY api ./api
COPY docs ./docs
RUN go build

EXPOSE 3001

CMD [ "/app/music-wishlist-api" ]
