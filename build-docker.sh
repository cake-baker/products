#!/bin/bash

env GOOS=linux GOARCH=386 go build -o products cmd/products/main.go
docker build . -t cakebakery/products:v1