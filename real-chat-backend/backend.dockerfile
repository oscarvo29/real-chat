# build a tiny docker image

FROM alpine:latest

RUN mkdir /app

COPY backendApp /app

CMD ["/app/backendApp"]