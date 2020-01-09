FROM golang:latest
ENV GOPATH /go
ENV PATH $PATH:$GOPATH/bin
RUN mkdir -p /go/src/github.com/riita10069/jc
COPY . /go/src/github.com/riita10069/jc
WORKDIR /go/src/github.com/riita10069/jc
ENV GO111MODULE=on
