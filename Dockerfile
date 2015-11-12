FROM golang:1.5
MAINTAINER xtaci <daniel820313@gmail.com>
ENV GOBIN /go/bin
COPY .godeps /go/.godeps
COPY src /go/src
WORKDIR /go
ENV GOPATH /go:/go/.godeps
RUN go install agent
RUN rm -rf pkg src .godeps
ENTRYPOINT /go/bin/agent
EXPOSE 8888
