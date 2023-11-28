ARG FIREHOSE_IMAGE=""
ARG DEB_PKG=""

FROM ${FIREHOSE_IMAGE}
ARG DEB_PKG
RUN apt update &&\
    apt-get -y install curl ca-certificates libc6 libgcc1 libstdc++6 libtinfo5 zlib1g libusb-1.0-0 libcurl3-gnutls
RUN mkdir -p /var/cache/apt/archives/
ADD ${DEB_PKG} /var/cache/apt/archives/
RUN dpkg -i /var/cache/apt/archives/${DEB_PKG}
RUN rm -rf /var/cache/apt/*
