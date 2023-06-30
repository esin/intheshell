#!/bin/bash

dt=$(date +%y%m%d%H%M)
docker build --load -t es1n/intheshell:latest -t es1n/intheshell:$dt -f Dockerfile .
if [ $? -ne 0 ]; then
  exit 1;
fi

docker push es1n/intheshell:$dt

docker push es1n/intheshell:latest

exit 0
