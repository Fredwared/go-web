# Build Golang binary
FROM golang:1.20.2 AS build-golang

WORKDIR /app

COPY ./api .
RUN go get -v && go build -v -o /usr/local/bin/api

EXPOSE 8080
CMD ["api"]
