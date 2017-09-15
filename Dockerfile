FROM golang:latest

WORKDIR /root

ENV PATH=${PATH}:/usr/local/go/bin GOROOT=/usr/local/go GOPATH=/root/go

RUN go get -u github.com/valyala/fasthttp && go get -u github.com/mailru/easyjson && go get -u github.com/buaazp/fasthttprouter

ADD . go/src/hl/

RUN go build hl && go install hl

EXPOSE 80

ENTRYPOINT ./go/bin/hl