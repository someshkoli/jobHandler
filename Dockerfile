FROM golang

ADD . /go/src/github.com/someshkoli/jobHandler

WORKDIR /go/src/github.com/someshkoli/jobHandler

RUN go mod download
RUN go install github.com/someshkoli/jobHandler

EXPOSE 8000

ENTRYPOINT  /go/bin/jobHandler
