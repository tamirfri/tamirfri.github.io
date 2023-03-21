#!/bin/sh

go install golang.org/x/tools/cmd/goimports@latest
go install github.com/tamirfri/gopy@v0.4.7

python3.11 -m pip install --quiet --upgrade pybindgen setuptools wheel pip

rm -rf py_run_go node_modules package*.json main.wasm
