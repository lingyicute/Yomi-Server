#!/usr/bin/env bash

PWD=`pwd`
PAPERCRAFTAPP=${PWD}"/app"
INSTALL=${PWD}"/papercraftd"

echo "build idgen ..."
cd ${PAPERCRAFTAPP}/service/idgen/cmd/idgen
go build -o ${INSTALL}/bin/idgen

echo "build status ..."
cd ${PAPERCRAFTAPP}/service/status/cmd/status
go build -o ${INSTALL}/bin/status

echo "build dfs ..."
cd ${PAPERCRAFTAPP}/service/dfs/cmd/dfs
go build -o ${INSTALL}/bin/dfs

echo "build media ..."
cd ${PAPERCRAFTAPP}/service/media/cmd/media
go build -o ${INSTALL}/bin/media

echo "build authsession ..."
cd ${PAPERCRAFTAPP}/service/authsession/cmd/authsession
go build -o ${INSTALL}/bin/authsession

echo "build biz ..."
cd ${PAPERCRAFTAPP}/service/biz/biz/cmd/biz
go build -o ${INSTALL}/bin/biz

echo "build msg ..."
cd ${PAPERCRAFTAPP}/messenger/msg/cmd/msg
go build -o ${INSTALL}/bin/msg

echo "build sync ..."
cd ${PAPERCRAFTAPP}/messenger/sync/cmd/sync
go build -o ${INSTALL}/bin/sync

echo "build bff ..."
cd ${PAPERCRAFTAPP}/bff/bff/cmd/bff
go build -o ${INSTALL}/bin/bff

echo "build session ..."
cd ${PAPERCRAFTAPP}/interface/session/cmd/session
go build -o ${INSTALL}/bin/session

echo "build gateway ..."
cd ${PAPERCRAFTAPP}/interface/gateway/cmd/gateway
go build -o ${INSTALL}/bin/gateway

echo "build gnetway ..."
cd ${PAPERCRAFTAPP}/interface/gnetway/cmd/gnetway
go build -o ${INSTALL}/bin/gnetway

#echo "build httpserver ..."
#cd ${PAPERCRAFTAPP}/interface/httpserver/cmd/httpserver
#go build -o ${INSTALL}/bin/httpserver
