#!/bin/bash

#github.com/golang/protobuf/proto
#TODO go get first

echo "build and install library"

cd ./src/bitbotlib
go install

echo "build and install console"
cd ../console
go install

echo "build and install robot"
cd ../robot
go install

echo "copy files to bin folder"
cd ../../bin
mv console ../dist/
mv robot ../dist/

echo "back to original folder"
cd ..

echo "Build websites"
cd hugo

source deploy.sh

echo "Done."