#!/bin/sh
set -e
(
  cd "$(dirname "$0")"
  go build -o /tmp/simple-stupid-dns-server-go-build app/*.go
)

exec /tmp/simple-stupid-dns-server-go-build "$@"
