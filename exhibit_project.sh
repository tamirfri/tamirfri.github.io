#!/bin/sh

pip3 install pybindgen
go install golang.org/x/tools/cmd/goimports@latest
go install github.com/go-python/gopy@latest

set -x

rm -r py_run_go node_modules package*.json

npm i --save lodash

gopy build -output=py_run_go -vm=python3 github.com/tamirfri/py-run-go-run-js

python3 -c "

import py_run_go.lodash as lodash

print(lodash.Find('[{\"a\":1, \"b\":2, \"d\":true}, {\"b\":2, \"c\":3, \"d\":false}, {\"a\":1, \"d\":true}]', '{\"b\":2,\"c\":3,\"d\":false}'))
"
