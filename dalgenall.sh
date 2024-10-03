#!/usr/bin/env bash

PWD=`pwd`
PAPERCRAFTAPP=${PWD}"/app"
echo ${PWD}

cd ${PAPERCRAFTAPP}/messenger/msg/internal/dal/dalgen
./dalgen_all.sh

cd ${PAPERCRAFTAPP}/messenger/sync/internal/dal/dalgen
./dalgen_all.sh

cd ${PAPERCRAFTAPP}/bff/authorization/internal/dal/dalgen
./dalgen_all.sh

cd ${PAPERCRAFTAPP}/service/biz/chat/internal/dal/dalgen
./dalgen_all.sh

cd ${PAPERCRAFTAPP}/service/biz/message/internal/dal/dalgen
./dalgen_all.sh

cd ${PAPERCRAFTAPP}/service/biz/user/internal/dal/dalgen
./dalgen_all.sh

cd ${PAPERCRAFTAPP}/service/biz/updates/internal/dal/dalgen
./dalgen_all.sh

cd ${PAPERCRAFTAPP}/service/biz/dialog/internal/dal/dalgen
./dalgen_all.sh

cd ${PAPERCRAFTAPP}/service/biz/username/internal/dal/dalgen
./dalgen_all.sh

cd ${PAPERCRAFTAPP}/service/authsession/internal/dal/dalgen
./dalgen_all.sh

cd ${PAPERCRAFTAPP}/service/media/internal/dal/dalgen
./dalgen_all.sh

