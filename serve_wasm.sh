#!/bin/sh

GOOS=js GOARCH=wasm go build -o main.wasm cmd/main.go

npx http-server -c-1
