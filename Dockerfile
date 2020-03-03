FROM golang:alpine AS builder

ADD ./ /go/src/github.com/S117Carlos/golang-simple-todo-list
RUN mkdir -p /github.com/S117Carlos/golang-simple-todo-list
WORKDIR /github.com/S117Carlos/golang-simple-todo-list
RUN go build -o main github.com/S117Carlos/golang-simple-todo-list
CMD ["/github.com/S117Carlos/golang-simple-todo-list/main"]
EXPOSE 8081