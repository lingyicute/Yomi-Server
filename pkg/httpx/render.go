// Copyright 2022 Papercraft Authors
//  All rights reserved.
//
// Author: @lingyicute
//

package httpx

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/pkg/httpx/render"

	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// JSONError
// JSON serializes the given struct as JSON into the response body.
// It also sets the Content-Type as "application/json".
func JSONError(w http.ResponseWriter, err error) {
	var (
		rErr *mtproto.TLRpcError
	)

	switch {
	case errors.As(err, &rErr):
		httpx.WriteJson(w, http.StatusOK, render.JSON{
			Ok:          false,
			ErrorCode:   rErr.Code(),
			Description: rErr.Message(),
		})
	default:
		httpx.WriteJson(w, http.StatusOK, render.JSON{
			Ok:          false,
			ErrorCode:   500,
			Description: err.Error(),
		})
	}
}

func JSON(w http.ResponseWriter, data interface{}) {
	// r, err := jsonx.Marshal(data)
	var (
		err error
		r   []byte
	)
	switch data.(type) {
	case proto.Message:
		r, err = protojson.MarshalOptions{UseProtoNames: true}.Marshal(data.(proto.Message))
		if err != nil {
			JSONError(w, err)
		}
	default:
		r, err = json.Marshal(data)
	}

	if err != nil {
		JSONError(w, err)
	} else {
		httpx.WriteJson(w, http.StatusOK, render.JSON{
			Ok:     true,
			Result: r,
		})
	}
}
