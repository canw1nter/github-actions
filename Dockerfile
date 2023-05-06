FROM golang:1.20.4-bullseye

COPY . /go-server

RUN cd /go-server

WORKDIR /go-server

ENTRYPOINT ["go", "run", "main.go"]