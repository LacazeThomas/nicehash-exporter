FROM golang:alpine as builder

RUN apk update && apk add git 
COPY . $GOPATH/src/lacazethomas/nicehash-exporter/
WORKDIR $GOPATH/src/lacazethomas/nicehash-exporter/
RUN go get -d -v
RUN go build -o /go/bin/nicehash-exporter

FROM alpine
EXPOSE 9159
ENV ENVIRONMENT prod
COPY --from=builder /go/bin/nicehash-exporter /bin/nicehash-exporter
ENTRYPOINT ["/bin/nicehash-exporter"]