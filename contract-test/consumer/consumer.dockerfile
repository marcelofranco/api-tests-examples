FROM alpine:latest

RUN mkdir /app

COPY bin/consumerApp /app

CMD ["/app/consumerApp"]