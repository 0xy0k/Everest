#!/bin/bash

# To get tsukid version, should install binary with `make install`
make install

tsukid version
# 1.0.0

tsukid version --long
# name: tsuki
# server_name: tsukid
# version: 1.0.0
# commit: cad573dbd60f477799e241589410a155f988d120
# build_tags: ""
# go: go version go1.16 darwin/amd64
# build_deps:
# - github.com/99designs/keyring@v1.1.6
