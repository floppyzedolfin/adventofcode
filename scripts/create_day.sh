#! /usr/bin/env bash

# This script should be called from the root of the repo.

if [ ! -f ./Makefile ]; then
    exit 1
fi

if [ $# -ne 1 ]; then
    echo "missing parameter day"
    exit 1
fi

# cast to int
day=$(($1))
# print to the right format
day=$(printf "%02d" ${day})

mkdir 2020/dec${day}
cp -r 2020/decXX/* 2020/dec${day}/.
mv  2020/dec${day}/decXX.go 2020/dec${day}/dec${day}.go
mv  2020/dec${day}/decXX_test.go 2020/dec${day}/dec${day}_test.go
sed -i -e "s/XX/${day}/g" 2020/dec${day}/dec${day}.go
sed -i -e "s/XX/${day}/g" 2020/dec${day}/dec${day}_test.go
