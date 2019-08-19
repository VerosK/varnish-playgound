FROM alpine:3.3
MAINTAINER verosk
COPY ./working-server /bin/working-server
COPY ./broken-server /bin/broken-server
EXPOSE 8983
ENTRYPOINT [ "/bin/working-server" ]
