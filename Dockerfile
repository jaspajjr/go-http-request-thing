FROM  golang:1.8-stretch

COPY container-entrypoint.sh /entry
RUN chmod +x /entry

COPY src /go/src/http-thing

CMD [ "foo" ]
ENTRYPOINT [ "/entry" ]