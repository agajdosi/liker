FROM alpine

RUN mkdir -p /liker
WORKDIR /liker
COPY . /liker

RUN ls /
RUN ls /liker

RUN chmod 755 /liker/build-binary
