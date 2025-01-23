# build a tiny docker image

FROM alpine:latest

RUN mkdir /app

COPY chatListenerApp /app

CMD ["/app/chatListenerApp"]