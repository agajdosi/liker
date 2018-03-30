FROM golang:1.9
USER nobody

RUN mkdir -p /go/src/github.com/agajdosi/liker/build-binary
WORKDIR /go/src/github.com/agajdosi/liker/build-binary
COPY . /go/src/github.com/agajdosi/liker/build-binary
RUN go-wrapper download && go-wrapper install
RUN CGO_ENABLED=0 go build -v
RUN ls
RUN pwd
