echo off

REM TODO go get first

echo  build and install library

cd src/bitbot
go install

echo build and install console
cd console
go install

echo build and install robot
cd ..
cd robot
go install

echo copy files to bin folder
cd ../../../bin
copy  /Y console.exe ..\dist\
copy  /Y robot.exe ..\dist\

REM back to original folder
cd ..

echo Done.