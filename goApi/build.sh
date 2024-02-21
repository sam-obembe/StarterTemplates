#!/bin/sh

rm api
GOOS=linux GOARCH=amd64 go build