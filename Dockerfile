FROM golang:1.6.2-alpine
MAINTAINER Balamurali Pandranki <balamurali@live.com>


#RUN \
#apt-get update && \
#apt-get install -y build-essential && \
#apt-get install -y swig python3-dev python3-numpy python3-scipy && \
#apt-get install -y sox bison curl


#PocketSphinx
ENV SPHINXBASE   sphinxbase-5prealpha.tar.gz
ENV POCKETSPHINX pocketsphinx-5prealpha.tar.gz
ENV SPHINXTRAIN  sphinxtrain-5prealpha.tar.gz

ADD  ./data/pocketsphinx/sphinxbase-5prealpha.tar.gz /sphinx/
ADD  ./data/pocketsphinx/pocketsphinx-5prealpha.tar.gz /sphinx/
ADD  ./data/pocketsphinx/sphinxtrain-5prealpha.tar.gz  /sphinx/

RUN mv /sphinx/sphinxbase-5prealpha   /sphinx/sphinxbase
RUN mv /sphinx/pocketsphinx-5prealpha /sphinx/pocketsphinx
RUN mv /sphinx/sphinxtrain-5prealpha  /sphinx/sphinxtrain

#WORKDIR /sphinx/sphinxbase
#RUN ./configure --with-swig-python
#RUN make
#RUN make install
#RUN make check

#WORKDIR /sphinx/pocketsphinx
#RUN ./configure --with-swig-python
#RUN make
#RUN make check
#RUN make install
#RUN make installcheck

#WORKDIR /sphinx/sphinxtrain
#RUN ./configure
#RUN make
#RUN make check
#RUN make installcheck

WORKDIR /data

#ENV PYTHONPATH /usr/local/lib/python3.4/site-packages

VOLUME /sphinx
VOLUME /data