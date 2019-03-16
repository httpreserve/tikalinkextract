#!/usr/bin/env bash
set -eux

TLE="tikalinkextract"
DIR="build"
rm -rf "$DIR"
mkdir "$DIR"
env GOOS=windows GOARCH=386 go build
mv "$TLE".exe "${DIR}/${TLE}"-win386.exe
env GOOS=windows GOARCH=amd64 go build
mv "$TLE".exe "${DIR}/${TLE}"-win64.exe
env GOOS=linux GOARCH=amd64 go build
mv "$TLE" "${DIR}/${TLE}"-linux64
env GOOS=darwin GOARCH=386 go build
mv "$TLE" "${DIR}/${TLE}"-darwin386
env GOOS=darwin GOARCH=amd64 go build
mv "$TLE" "${DIR}/${TLE}"-darwinAmd64
