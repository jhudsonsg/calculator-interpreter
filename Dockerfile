FROM ubuntu
WORKDIR /root/go/src/github.com/jhudsonsg/calculator-interpreter
COPY . .

RUN apt-get update -y \
    && apt-get upgrade -y \
    && apt-get install curl -y \
    && apt-get install make -y \
    && curl -O https://dl.google.com/go/go1.13.7.linux-amd64.tar.gz \
    && sha256sum go1.13.7.linux-amd64.tar.gz \
    && tar xvf go1.13.7.linux-amd64.tar.gz \
    && rm go1.13.7.linux-amd64.tar.gz \
    && chown -R root:root ./go \
    && mv ./go /usr/local

ENV GOPATH="${HOME}/work"
ENV PATH="${PATH}:/usr/local/go/bin:${GOPATH}/bin"

RUN echo "Start you service"