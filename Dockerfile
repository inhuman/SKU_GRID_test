FROM alpine
LABEL maintainer="msgexec@gmail.com"
COPY ./bin/sku_server /usr/local/bin
RUN chmod +x /usr/local/bin/sku_server
RUN apk update && apk add ca-certificates
ENTRYPOINT /usr/local/bin/sku_server