#!/usr/bin/env bash

CGO_ENABLED=0 go build -ldflags "-s -w" -o /tmp/ytmp/socks5server ./cmd/socks5server/main.go
