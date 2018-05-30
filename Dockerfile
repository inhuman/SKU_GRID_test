FROM alpine
LABEL maintainer="msgexec@gmail.com"
COPY ./bin/sku_server /usr/local/bin
RUN chmod +x /usr/local/bin/sku_server
