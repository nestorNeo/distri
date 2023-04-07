#!/usr/bin/env bash

wr=`pwd`
echo "trabajando en directorio $wr"

rm -Rf "$wr/logs"
mkdir -p "$wr/logs"

serversz=("chela" "tequila" "vodka")

for server in ${serversz[@]}; do

    program="$wr/compilados/$server.exe"
    output="$wr/logs/$server"
    echo $program
    #nohup some_command &> nohup2.out &
    nohup $program  &> $output &
done