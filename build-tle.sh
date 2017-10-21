#!/usr/bin/env bash

env GOOS=windows GOARCH=386 go build
mv tikalinkextract.exe tikalinkextract-win386.exe
env GOOS=windows GOARCH=amd64 go build
mv tikalinkextract.exe tikalinkextract-win64.exe
env GOOS=linux GOARCH=amd64 go build
mv tikalinkextract tikalinkextract-linux64
