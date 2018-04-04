FROM alpine:latest

RUN ls /
RUN mkdir -p /home/liker
WORKDIR /home/liker
COPY . /home/liker

RUN chmod 755 /home/liker/build-binary

RUN dnf install google-chrome-stable -y
