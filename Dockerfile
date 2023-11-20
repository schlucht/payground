FROM alpine

RUN apk update && apk add neovim \\
    apk add wget -y

COPY config/ /root/