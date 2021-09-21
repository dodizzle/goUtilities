#!/bin/bash

DIR=$(echo "${PWD##*/}")
lower_DIR=$(echo $DIR| tr '[:upper:]' '[:lower:]')
date=`date '+%Y-%m-%d-%H-%M-%S_sd88sdfsf'`
if test -f $DIR; then
    rm $DIR
fi
docker build -t us.gcr.io/helix-global/spin-test:$date .
docker push us.gcr.io/helix-global/spin-test:$date

if test -f $DIR; then
    rm $DIR
fi