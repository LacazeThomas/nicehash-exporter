FROM golang:alpine as builder

RUN apk update && apk add git 
COPY . $GOPATH/src/lacazethomas/nicehashExporter/
WORKDIR $GOPATH/src/lacazethomas/nicehashExporter/
RUN go get -d -v
RUN go build -o /go/bin/nicehashExporter

FROM alpine
EXPOSE 9159
ENV ENVIRONMENT prod
COPY --from=builder /go/bin/nicehashExporter /bin/nicehashExporter
ENTRYPOINT ["/bin/nicehashExporter"]