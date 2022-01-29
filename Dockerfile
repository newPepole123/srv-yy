FROM scratch

COPY ./docker/etc/ /etc/
COPY ./bin/srv-yy-linux-amd64 /usr/bin/srv-yy
USER nobody
WORKDIR /app
EXPOSE 8443
ENTRYPOINT ["srv-yy"]

CMD ["serve"]
