FROM alpine:latest

RUN ls /
RUN mkdir -p /home/liker
WORKDIR /home/liker
COPY . /home/liker

RUN chmod 755 /home/liker/build-binary

#RUN apk --update --no-cache add chromium udev mesa-gl mesa-dri-swrast
