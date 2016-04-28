FROM golang:1.6.2
MAINTAINER Balamurali Pandranki <balamurali@live.com>


RUN \
apt-get update && \
apt-get install -y build-essential && \
apt-get install -y swig python3-dev python3-numpy python3-scipy && \
apt-get install -y sox bison curl

#ENV PYTHONPATH /usr/local/lib/python3.4/site-packages

#PocketSphinx
ENV SPHINXBASE   sphinxbase-5prealpha.tar.gz
ENV POCKETSPHINX pocketsphinx-5prealpha.tar.gz
#ENV SPHINXTRAIN  sphinxtrain-5prealpha.tar.gz

ADD  ./data/pocketsphinx/sphinxbase-5prealpha.tar.gz /sphinx/
ADD  ./data/pocketsphinx/pocketsphinx-5prealpha.tar.gz /sphinx/
#ADD  ./data/pocketsphinx/sphinxtrain-5prealpha.tar.gz  /sphinx/

RUN mv /sphinx/sphinxbase-5prealpha   /sphinx/sphinxbase
RUN mv /sphinx/pocketsphinx-5prealpha /sphinx/pocketsphinx
#RUN mv /sphinx/sphinxtrain-5prealpha  /sphinx/sphinxtrain

RUN apt-get install -y python-dev

WORKDIR /sphinx/sphinxbase
RUN ./configure --with-swig-python
RUN make
RUN make install
RUN make check

WORKDIR /sphinx/pocketsphinx
RUN ./configure --with-swig-python
RUN make
RUN make check
RUN make install
RUN make installcheck

#WORKDIR /sphinx/sphinxtrain
#RUN ./configure
#RUN make
#RUN make check
#RUN make installcheck

RUN apt-get install -y pkg-config
ENV PORT 8080
ADD . /go/src/github.com/itsbalamurali/heyasha
WORKDIR /go/src/github.com/itsbalamurali/heyasha
RUN go install ./...
ENTRYPOINT /go/bin/heyasha
EXPOSE 8080

WORKDIR /data
VOLUME /sphinx
VOLUME /data