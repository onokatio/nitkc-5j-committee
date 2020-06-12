#!/bin/bash

NAME=委員会割り当て

GOOS=(linux darwin windows)
GOARCH=(amd64 386)
for os in ${GOOS[@]}; do
	for arch in ${GOARCH[@]}; do
		if [ "$os" == "windows" ];then
			ext=".exe"
		else
			ext=".bin"
		fi

		if [ "$os" == "darwin" ];then
			os_string="Mac-OSX"
		else
			os_string=$os
		fi

		if [ "$arch" == "amd64" ];then
			arch_string="64bit"
		else
			arch_string="32bit"
		fi

		GOOS=$os GOARCH=$arch go build -o dst/$os_string-${arch_string}-${NAME}${ext} main.go
	done
done
