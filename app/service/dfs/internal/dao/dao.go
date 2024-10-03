/*
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright (c) 2021-present,  Papercraft Studio (https://papercraft.io).
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package dao

import (
	"github.com/lingyicute/papercraft-server/app/service/dfs/internal/config"
	"github.com/lingyicute/papercraft-server/app/service/dfs/internal/minio_util"
	idgen_client "github.com/lingyicute/papercraft-server/app/service/idgen/client"

	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/core/stores/kv"
	"github.com/zeromicro/go-zero/zrpc"
)

type Dao struct {
	minio *minio.Core
	idgen_client.IDGenClient2
	ssdb kv.Store
}

func New(c config.Config) *Dao {
	return &Dao{
		minio:        minio_util.MustNewMinioClient(&c.Minio),
		IDGenClient2: idgen_client.NewIDGenClient2(zrpc.MustNewClient(c.IdGen)),
		ssdb:         kv.NewStore(c.SSDB),
	}
}

func NewDFSHelper(minio *minio_util.MinioConfig, idgen zrpc.RpcClientConf, ssdb kv.KvConf) *Dao {
	return &Dao{
		minio:        minio_util.MustNewMinioClient(minio),
		IDGenClient2: idgen_client.NewIDGenClient2(zrpc.MustNewClient(idgen)),
		ssdb:         kv.NewStore(ssdb),
	}
}

func NewMinioHelper(minio *minio_util.MinioConfig) minio_util.MinioHelper {
	return &Dao{
		minio: minio_util.MustNewMinioClient(minio),
	}
}
