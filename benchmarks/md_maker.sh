#!/bin/sh

DIR_NAME=${PWD##*/}
GOVERSION=$(go version)
NAME="$(tr '[:lower:]' '[:upper:]' <<< ${DIR_NAME:0:1})${DIR_NAME:1}"
NAME_LOWER="$(tr '[:upper:]' '[:lower:]' <<< $NAME)"
NAME="${NAME//'_'/ }"

BODY="## $NAME Results\n"
BODY+="Benchmark Name|Iterations|Per-Iteration\n"
BODY+="----|----|----\n"
TABLE="$(grep Benchmark | tr '\t' '|'| tr '\n' '#')"
BODY+="${TABLE//'#'/\n}"
BODY+="\n"
BODY+="Generated using $GOVERSION\n"

MD_FILE_NAME=$NAME_LOWER".md"
echo $(rm -f "./$MD_FILE_NAME")
echo $BODY >> ./$MD_FILE_NAME

echo "- [$NAME](./$DIR_NAME/$MD_FILE_NAME)" >> "../index.md"