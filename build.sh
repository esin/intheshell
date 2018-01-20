#!/bin/bash

### Build 32 bit
env GOOS=linux GOARCH=386 go build -o intheshell .

### Build 64 bit
env GOOS=linux GOARCH=amd64 go build -o intheshell64 .

### Build arm
env GOOS=linux GOARCH=arm go build -o intheshell_arm .
