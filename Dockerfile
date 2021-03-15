FROM golang:buster
RUN apt-get install gcc
COPY go.mod go.sum hello.go ./prog/
WORKDIR prog/
RUN go mod tidy
RUN go build .
ENTRYPOINT ./torrenthygge
