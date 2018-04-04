FROM alpine:latest

USER chrome

RUN sudo mkdir -p /liker
WORKDIR /liker
COPY . /liker

RUN chmod 755 /liker/build-binary
