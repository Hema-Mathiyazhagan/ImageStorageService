FROM golang:1.17.0-alpine AS build

RUN apk update \
    && apk add --no-cache go make git gcc musl-dev build-base

WORKDIR /home/app

COPY lib /home/app/lib
COPY model /home/app/model
COPY apis /home/app/apis
COPY docs /home/app/docs
COPY go.mod go.sum main.go Makefile /home/app/

RUN go mod tidy

RUN make

FROM golang:1.17.0-alpine AS prod

RUN mkdir -p /home/app \
    && mkdir -p /home/app/image-db

WORKDIR /home/app

COPY --from=build /home/app/image-storage-service /home/app/

CMD ./image-storage-service