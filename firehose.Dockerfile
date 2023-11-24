# syntax=docker/dockerfile:1.2

FROM golang:1.21 as firehose
ARG COMMIT
ARG VERSION
RUN mkdir -p /work
ADD . /work
WORKDIR /work
RUN go test ./...
RUN go build -ldflags "-s -w -X main.version=${VERSION} -X main.commit=${COMMIT}" -v -o /work/fireantelope ./cmd/fireantelope

FROM ubuntu:22.04 as base

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get -y install -y \
    ca-certificates vim htop iotop sysstat wget \
    dstat strace lsof curl jq tzdata && \
    rm -rf /var/cache/apt /var/lib/apt/lists/*

RUN wget http://archive.ubuntu.com/ubuntu/pool/main/o/openssl/libssl1.1_1.1.0g-2ubuntu4_amd64.deb
RUN dpkg -i libssl1.1_1.1.0g-2ubuntu4_amd64.deb

RUN rm /etc/localtime && ln -snf /usr/share/zoneinfo/America/Montreal /etc/localtime && dpkg-reconfigure -f noninteractive tzdata

RUN mkdir /tmp/wasmer-install && cd /tmp/wasmer-install && \
    curl -L https://github.com/wasmerio/wasmer/releases/download/2.3.0/wasmer-linux-amd64.tar.gz | tar xzf - && \
    mv lib/libwasmer.a lib/libwasmer.so /usr/lib/ && cd / && rm -rf /tmp/wasmer-install

COPY --from=firehose /work/fireantelope /app/fireantelope
COPY --from=firehose /work/tools/fireantelope/motd_generic /etc/
COPY --from=firehose /work/tools/fireantelope/motd_node_manager /etc/
COPY --from=firehose /work/tools/fireantelope/99-firehose.sh /etc/profile.d/
COPY --from=firehose /work/tools/fireantelope/scripts/* /usr/local/bin

ENV PATH=$PATH:/app
