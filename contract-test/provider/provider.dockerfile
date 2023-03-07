FROM alpine:latest

RUN mkdir /app

COPY bin/providerApp /app

CMD ["/app/providerApp"]