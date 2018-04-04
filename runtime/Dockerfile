FROM fedora:latest

RUN mkdir -p /home/liker
WORKDIR /home/liker
COPY . /home/liker

RUN chmod 755 /home/liker/build-binary

COPY ./google-chrome.repo /etc/yum.repos.d/google-chrome.repo
RUN dnf install google-chrome-stable -y
