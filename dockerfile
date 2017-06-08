FROM alpine:latest
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN mkdir -p /opt/vice/
WORKDIR /opt/vice/
ADD vice-api /opt/vice/
RUN chmod +x /opt/vice/vice-api
CMD /opt/vice/vice-api