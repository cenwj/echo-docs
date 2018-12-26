#!/usr/bin/env bash

distDir=./dist
packName=echo-docs

if [ ! -d $distDir ]; then
mkdir $distDir
fi
echo "bulding linux version..."
GOOS=linux GOARCH=amd64 go build -o $distDir/$packName server.go
#echo "bulding windows version..."
#GOOS=windows GOARCH=amd64 go build -o $distDir/${packName}.exe server.go
echo "bulding macos version..."
GOOS=darwin GOARCH=amd64 go build -o $distDir/${packName}.mac server.go

echo "copy file..."
rm -rf ${distDir}/public
cp -R public $distDir
rm -rf ${distDir}/template
cp -R template $distDir
