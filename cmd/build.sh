#!/usr/bin/env bash

wr=`pwd`
echo "trabajando en directorio $wr"
rm -Rf "$wr/compilados"
mkdir -p "$wr/compilados"

serversz=("chela" "tequila" "vodka")

for server in ${serversz[@]}; do
    echo "$wr/$server";
    cd "$wr/$server";
    go build;
    mv $server "$wr/compilados";
done