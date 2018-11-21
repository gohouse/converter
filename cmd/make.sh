#!/usr/bin/env bash
go build -o table2struct.darwin.bin cli.go
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o table2struct.linux.bin cli.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o table2struct.win.exe cli.go