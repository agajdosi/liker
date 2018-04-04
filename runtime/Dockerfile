FROM alpine:latest

RUN mkdir -p /liker
WORKDIR /liker
COPY . /liker

RUN chmod 755 /liker/build-binary

#RUN apk --update --no-cache add chromium udev mesa-gl mesa-dri-swrast
