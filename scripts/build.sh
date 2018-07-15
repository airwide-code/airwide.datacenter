#!/usr/bin/env bash

# todo(yumcoder) change dc ip
# sed -i '/ipAddress = /c\ipAddress = 127.0.0.1' a.txt
# todo(yumcoder) change folder path for nbfs

docker start mysql-docker redis-docker etcd-docker
export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH
export GOPATH=/home/airwide
airwide_datacenter="/home/airwide/src/github.com/airwide-code/airwide.datacenter"

echo "build frontend ..."
cd ${airwide_datacenter}/access/frontend
go get
go build
./frontend &

echo "build auth_key ..."
cd ${airwide_datacenter}/access/auth_key
go get
go build
.//auth_key &

echo "build sync ..."
cd ${airwide_datacenter}/push/sync
go get
go build
./sync &

echo "build nbfs ..."
cd ${airwide_datacenter}/nbfs/nbfs
go get
go build
./nbfs &

echo "build biz_server ..."
cd ${airwide_datacenter}/biz_server
go get
go build
./biz_server &

echo "build session ..."
cd ${airwide_datacenter}/access/session
go get
go build
./session &

echo "***** wait *****"
wait
