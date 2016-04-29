FROM golang:1.6.2
MAINTAINER Balamurali Pandranki <balamurali@live.com>


RUN \
apt-get update && \
apt-get install -y build-essential && \
apt-get install -y swig && \
apt-get install -y sox bison curl

#PocketSphinx & sphinxbase
ENV SPHINXBASE   sphinxbase-5prealpha.tar.gz
ENV POCKETSPHINX pocketsphinx-5prealpha.tar.gz

ADD  ./data/pocketsphinx/sphinxbase-5prealpha.tar.gz /sphinx/
ADD  ./data/pocketsphinx/pocketsphinx-5prealpha.tar.gz /sphinx/

RUN mv /sphinx/sphinxbase-5prealpha   /sphinx/sphinxbase
RUN mv /sphinx/pocketsphinx-5prealpha /sphinx/pocketsphinx

RUN apt-get install -y python-dev pkg-config

WORKDIR /sphinx/sphinxbase
RUN ./configure
RUN make
RUN make install
#RUN make check

WORKDIR /sphinx/pocketsphinx
RUN ./configure
RUN make
RUN make install
#RUN make check
#RUN make installcheck


ENV PORT 8080
ADD . /go/src/github.com/itsbalamurali/heyasha
WORKDIR /go/src/github.com/itsbalamurali/heyasha
RUN go install ./...
ENTRYPOINT /go/bin/heyasha
EXPOSE 8080

WORKDIR /data
VOLUME /sphinx
VOLUME /data