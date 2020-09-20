#!/bin/sh

ARGS="$@"

GOPATH=$(pwd)
GOBIN="$(pwd)/bin"

echo -e ">>>>> SET LOCAL ENV\n"
echo ".....GOPATH='$GOPATH'"
echo ".....GOBIN='$GOBIN'"
echo -e "\n"
eval "export GOPATH='$GOPATH'"
eval "export GOBIN='$GOBIN'"
echo -e ">>>>> RUN GO\n"
eval "go $ARGS"