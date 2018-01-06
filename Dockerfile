FROM  golang:1.8-stretch

COPY src /go/src/http-thing

CMD [ "go run /go/src/http-thing/main.go" ]