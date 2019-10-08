#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o bin/main
zip bin/main.zip bin/main
