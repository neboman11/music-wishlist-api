# syntax=docker/dockerfile:1

FROM golang:1.26-alpine3.24

RUN apk upgrade --no-cache

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY api ./api
COPY docs ./docs
COPY models ./models
RUN go build

EXPOSE 3001

CMD [ "/app/music-wishlist-api" ]
