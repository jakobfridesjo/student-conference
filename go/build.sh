#!/bin/sh

cd ./ints
go build
cd ../pointers
go build
cd ../memoryoverhead
go build
cd ..