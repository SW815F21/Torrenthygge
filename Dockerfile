FROM golang:buster
RUN apt-get install gcc
COPY go.mod go.sum hello.go denmark-latest.osm.pbf.torrent ./prog/
WORKDIR prog/
RUN go mod tidy
RUN go build .
ENV TORRENT_CLIENT_PORT=42069
EXPOSE $TORRENT_CLIENT_PORT/tcp
EXPOSE $TORRENT_CLIENT_PORT/udp
ENTRYPOINT ./torrenthygge
