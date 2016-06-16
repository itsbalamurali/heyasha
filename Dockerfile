FROM golang:1.6.2
MAINTAINER Balamurali Pandranki <balamurali@live.com>

VOLUME /go/src/github.com/itsbalamurali/heyasha

ENV PORT 80
ADD . /go/src/github.com/itsbalamurali/heyasha
WORKDIR /go/src/github.com/itsbalamurali/heyasha
RUN go install ./...
ENTRYPOINT /go/bin/heyasha
EXPOSE 80