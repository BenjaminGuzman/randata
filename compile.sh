#!/usr/bin/env bash

# compile for the local machine
go build

# compile for windows x86_64
GOOS=windows GOARCH=amd64 go build
