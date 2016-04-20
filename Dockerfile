FROM ubuntu:14.04
MAINTAINER Balamurali Pandranki <balamurali@live.com>

# Notes:
# 1. Recommended command to run:
# 2. Default install prefix for all modules is: /usr/local/

RUN \
apt-get update && \
apt-get install -y build-essential && \
apt-get install -y swig python3-dev python3-numpy python3-scipy && \
apt-get install -y sox bison curl

# gcc for cgo
RUN apt-get install -y --no-install-recommends \
		g++ \
		gcc \
		libc6-dev \
		make \
	&& rm -rf /var/lib/apt/lists/*

ENV GOLANG_VERSION 1.6.1
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 6d894da8b4ad3f7f6c295db0d73ccc3646bce630e1c43e662a0120681d47e988

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

COPY go-wrapper /usr/local/bin/


#PocketSphinx
ENV SPHINXBASE   sphinxbase-5prealpha.tar.gz
ENV POCKETSPHINX pocketsphinx-5prealpha.tar.gz
ENV SPHINXTRAIN  sphinxtrain-5prealpha.tar.gz

#ADD  /data/pocketsphinx/sphinxbase-5prealpha.tar.gz /sphinx/
#ADD  /data/pocketsphinx/pocketsphinx-5prealpha.tar.gz /sphinx/
#ADD  /data/pocketsphinx/sphinxtrain-5prealpha.tar.gz  /sphinx/
ADD http://tenet.dl.sourceforge.net/project/cmusphinx/sphinxtrain/5prealpha/sphinxtrain-5prealpha.tar.gz /sphinx/
ADD http://iweb.dl.sourceforge.net/project/cmusphinx/sphinxbase/5prealpha/sphinxbase-5prealpha.tar.gz /sphinx/
ADD http://iweb.dl.sourceforge.net/project/cmusphinx/pocketsphinx/5prealpha/pocketsphinx-5prealpha.tar.gz /sphinx/

RUN mv /sphinx/sphinxbase-5prealpha   /sphinx/sphinxbase
RUN mv /sphinx/pocketsphinx-5prealpha /sphinx/pocketsphinx
RUN mv /sphinx/sphinxtrain-5prealpha  /sphinx/sphinxtrain

WORKDIR /sphinx/sphinxbase
#RUN ls
RUN ./configure --with-swig-python
RUN make
RUN make install
RUN make check

WORKDIR /sphinx/pocketsphinx
#RUN ls
RUN ./configure --with-swig-python
RUN make
RUN make check
RUN make install
RUN make installcheck

WORKDIR /sphinx/sphinxtrain
#RUN ls
RUN ./configure
RUN make
RUN make check
RUN make installcheck

WORKDIR /data

# 'make install' installs all python modules to this dir.
# But Ubuntu recognizes only /usr/local/lib/python3.4/dist-packages
# dir by default. So add this dir to PYTHON_PATH manually.
ENV PYTHONPATH /usr/local/lib/python3.4/site-packages

VOLUME /sphinx
VOLUME /data