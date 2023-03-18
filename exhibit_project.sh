#!/bin/sh

go install golang.org/x/tools/cmd/goimports@latest
go install github.com/tamirfri/gopy@v0.4.7

python3.11 -m pip install --upgrade pybindgen setuptools wheel pip

set -x

rm -r py_run_go node_modules package*.json main.wasm

npm i --save lodash

gopy build -output=py_run_go -vm=python3.11 github.com/tamirfri/py-run-go-run-js

python3.11 -c "

import py_run_go.lodash as lodash

print(lodash.Find('[{\"a\":1, \"b\":2, \"d\":true}, {\"b\":2, \"c\":3, \"d\":false}, {\"a\":1, \"d\":true}]', '{\"b\":2,\"c\":3,\"d\":false}'))
"

GOOS=js GOARCH=wasm go build -o main.wasm cmd/main.go

npx http-server -c-1
