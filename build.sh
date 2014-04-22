#!/bin/bash
##############################
#####Setting Environments#####
echo "Setting Environments"
set -e
export PATH=$PATH:$GOPATH/bin:$HOME/bin:$GOROOT/bin
##############################
######Install Dependence######
echo "Installing Dependence"
go get github.com/Centny/Cny4go
go get github.com/Centny/igtest
##############################
#########Running Test#########
rm -f *.out
echo "mode: set" >a.out

echo "Running Unit Test"
pkgs="\
 github.com/Centny/igtest_ex/pkga\
 github.com/Centny/igtest_ex/pkgb\
"
for p in $pkgs;
do
 if [ "$1" = "-u" ];then
  go get -u $p
 fi
 go test -v -coverprofile c.out $p
 mrepo t.out a.out c.out
 cp t.out a.out
done

echo "Running Integration Test"
cpkgs="\
 github.com/Centny/igtest_ex/pkga\
,github.com/Centny/igtest_ex/pkgb\
,github.com/Centny/igtest_ex/srv\
"
go test -v -coverprofile c.out -coverpkg $cpkgs github.com/Centny/igtest_ex/srv
mrepo t.out a.out c.out
cp t.out a.out

cp $GOPATH/src/github.com/Centny/igtest_ex/srv/srv.json .
ig-repo srv.json igtest_ex

##############################
#####Create Coverage Report###
echo "Create Coverage Report"
gocov convert a.out > coverage.json
cat coverage.json | gocov-xml -b $GOPATH/src > coverage.xml
cat coverage.json | gocov-html coverage.json > coverage.html

