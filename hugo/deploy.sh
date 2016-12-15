#!/bin/bash

cd ../public
rm -rf *
cd ../hugo
hugo
rsync -avz -e "ssh -i ~/.ssh/kp-xak5varz" --progress ~/code/bitbot/public/ tong@207.226.143.243:/home/tong/bitbot.com.cn/

