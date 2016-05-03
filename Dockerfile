FROM golang:1.6.2
MAINTAINER Balamurali Pandranki <balamurali@live.com>

RUN \
apt-get update && \
apt-get install -y build-essential swig sox bison curl python-dev pkg-config libsasl2-dev

ENV LD_LIBRARY_PATH=/usr/local/lib
ENV PKG_CONFIG_PATH=/usr/local/lib/pkgconfig


#PocketSphinx & sphinxbase
ENV SPHINXBASE   sphinxbase-5prealpha.tar.gz
ENV POCKETSPHINX pocketsphinx-5prealpha.tar.gz

ADD  ./data/pocketsphinx/sphinxbase-5prealpha.tar.gz /sphinx/
ADD  ./data/pocketsphinx/pocketsphinx-5prealpha.tar.gz /sphinx/

RUN mv /sphinx/sphinxbase-5prealpha   /sphinx/sphinxbase
RUN mv /sphinx/pocketsphinx-5prealpha /sphinx/pocketsphinx

WORKDIR /sphinx/sphinxbase
RUN ./autogen.sh && \
./configure && \
make && \
make install

WORKDIR /sphinx/pocketsphinx
RUN ./configure && \
make && \
make install

ENV PORT 8080
ADD . /go/src/github.com/itsbalamurali/heyasha
WORKDIR /go/src/github.com/itsbalamurali/heyasha
RUN go install ./...
ENTRYPOINT /go/bin/heyasha
EXPOSE 8080

WORKDIR /data
VOLUME /sphinx
VOLUME /data