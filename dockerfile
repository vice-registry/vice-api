FROM alpine:latest
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN mkdir -p /opt/vice/
WORKDIR /opt/vice/
ADD vice-api /opt/vice/
RUN chmod +x /opt/vice/vice-api
ENV PORT=8080 \
    RETHINKDB_LOCATION=localhost \
    RETHINKDB_DATABASE=vice \
    RABBITMQ_LOCATION=localhost \
    RABBITMQ_USER=admin \
    RABBITMQ_PASS=admin
CMD /opt/vice/vice-api \
    --port $PORT \
    --host 0.0.0.0 \
    --rethinkdb-location $RETHINKDB_LOCATION \
    --rethinkdb-database $RETHINKDB_DATABASE \
    --rabbitmq-location $RABBITMQ_LOCATION \
    --rabbitmq-user $RABBITMQ_USER \
    --rabbitmq-pass $RABBITMQ_PASS