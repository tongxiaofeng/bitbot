#!/bin/bash

# argument "p" means Platform: 
# p=mac: GOOS=darwin GOARCH=amd64
# p=win32: GOOS=windows GOARCH=386
# p=win64: GOOS=windows GOARCH=amd64
# p=linux: GOOS=linux GOARCH=amd64
# p=arm: GOOS=linux GOARCH=arm

# argument "w" means builds web documents
GOOS=""
GOARCH=""
Platform=""
Version=`git describe --tags`
BuildWeb=false
BuildLDFlags="-X main.version=$Version -X main.buildstamp=`date -u '+%Y-%m-%d_%I:%M:%S%p'` -X main.githash=`git rev-parse HEAD`" 

echo $BuildLDFlags

while getopts "p:v:w" arg #选项后面的冒号表示该选项需要参数
do
        case $arg in
             p)
                Platform=$OPTARG
                ;;
             w)
                BuildWeb=true
                ;;                 
             ?)  #当有不认识的选项的时候arg为?
             echo "Example usage: build.sh -p mac|win32|win64|linux|arc -v 0.1.0"
             ;;
        esac
done

if [[ -z $Platform ]]; then
    echo "Example usage: build.sh -p mac/win32/win64/linux/arc"
else
    echo "Platform=$Platform"
fi

if [ -z $Platform ]; then
    GOOS="darwin"
    GOARCH="amd64"
elif [ $Platform = "win32" ]; then
    GOOS="windows"
    GOARCH="386"
elif [ $Platform = "win64" ];then
    GOOS="windows"
    GOARCH="amd64"
elif [ $Platform = "linux" ];then
    GOOS="linux"
    GOARCH="amd64"
elif [ $Platform = "arm" ];then
    GOOS="linux"
    GOARCH="arm"
else
    GOOS="darwin"
    GOARCH="amd64"
fi

echo "GOOS=$GOOS"
echo "GOARCH=$GOARCH"

ROOT=$(pwd)
echo $GOROOT
echo $GOPATH
echo $GOOS
echo $GOARCH
echo "############ Running go get. ############"
#go get google.golang.org/grpc
#go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
#go get github.com/robertkrimen/otto
#go get github.com/markcheno/go-talib
#Todo add more go get or use gb tool

echo "############ Build protoc. ############"
cd $ROOT/src/bitbotlib
protoc --go_out=plugins=grpc:. bitbot.proto

echo "############ build and install library. ############"

cd $ROOT/src/bitbotlib
GOOS=$GOOS GOARCH=$GOARCH go install

echo "############ remove all binary files. ############"
cd $ROOT/bin
rm -rf *

echo "############ build and install console. ############"
cd $ROOT/src/console
GOOS=$GOOS GOARCH=$GOARCH go install -ldflags "$BuildLDFlags"

echo "############ build and install robot. ############"
cd $ROOT/src/robot
GOOS=$GOOS GOARCH=$GOARCH go install -ldflags "$BuildLDFlags"

echo "############ copy files to dist folder. ############"
cd $ROOT/dist
if [ -d "$Version" ]; then
   rm -rf $Version
fi
mkdir $Version
cp -aR $ROOT/bin/. $ROOT/dist/$Version

echo "############ upload files to dist folder. ############"


if [ $BuildWeb = true ];then
echo "############ Building website documents. ############"

    echo "############ Build websites. ############"
    cd $ROOT/public
    rm -rf *
    cd $ROOT/hugo
    hugo
    rsync -avz -e "ssh -i ~/.ssh/kp-xak5varz" --progress ~/code/bitbot/public/ tong@207.226.143.243:/home/tong/bitbot.com.cn/
fi

cd $ROOT
echo "############ Done. ############"