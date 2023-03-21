#!/bin/sh

set -x

npm i --save lodash

gopy build -output=py_run_go -vm=python3.11 github.com/tamirfri/py-run-go-run-js

python3.11 -c "

import py_run_go.lodash as lodash

print(lodash.Find('[{\"user\":\"barney\",\"age\":36,\"active\":true},{\"user\":\"fred\",\"age\":40,\"active\":false}]', '{\"age\":36,\"active\":true}'))
"
