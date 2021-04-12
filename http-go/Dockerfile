FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt-get update && \
    apt-get install build-essential -y

ADD ./hello-world/hello-world /go/bin/hello-world

CMD ["/go/bin/hello-world"]