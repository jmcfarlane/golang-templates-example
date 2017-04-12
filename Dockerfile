FROM golang

EXPOSE 8080

ENV PROJECT /go/src/github.com/jmcfarlane/golang-templates-example

ADD . $PROJECT
WORKDIR $PROJECT

RUN go env
RUN go version
RUN go get github.com/GeertJohan/go.rice/rice
RUN go get -t ./...
run go generate
RUN go install

CMD ["/go/bin/golang-templates-example"]


