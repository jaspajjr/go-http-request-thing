#!/bin/bash
PATH=$PATH:/usr/sbin

set -e

if [ "$1" = "foo" ]; then
  exec go run /go/src/http-thing/src/main.go
elif [ "$1" = "-t" ]; then
  exec py.test -m src
fi
