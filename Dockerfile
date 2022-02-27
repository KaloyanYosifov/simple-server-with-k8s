FROM ubuntu:20.04

WORKDIR /app

EXPOSE 7011

RUN apt-get update
RUN apt-get install curl build-essential -y
RUN curl -OL https://golang.org/dl/go1.17.6.linux-arm64.tar.gz && \
    tar -C /usr/local -xvf go1.17.6.linux-arm64.tar.gz && \
    ln -s /usr/local/go/bin/go /usr/bin/go

COPY . /app

CMD go run main.go
