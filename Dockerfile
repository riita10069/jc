FROM asia.gcr.io/lc-dev-01/go-build:0.0.3 as builder

COPY ./ $GOPATH/src/github.com/riita10069/
RUN cd $GOPATH/src/github.com/riita10069/jc && go build


FROM alpine:latest

RUN apk add --no-cache ca-certificates tzdata

ENV GOPATH=/go

RUN mkdir -p $GOPATH/src/github.com/riita10069/jc
WORKDIR $GOPATH/src/github.com/riita10069/jc

COPY --from=builder $GOPATH/src/github.com/riita10069/jc/jc jc

RUN chmod +x lc-servejc

CMD ["./jc"]