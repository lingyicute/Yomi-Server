/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'mtprotoc'
 *
 * Copyright 2022 Papercraft Authors.
 *  All rights reserved.
 *
 * Author: papercraftio (papercraft.io@gmail.com)
 */

package service

import (
	"context"

	"github.com/papercraft/proto/mtproto"
	"github.com/lingyicute/papercraft-server/app/bff/nsfw/internal/core"
)

// AccountSetContentSettings
// account.setContentSettings#b574b16b flags:# sensitive_enabled:flags.0?true = Bool;
func (s *Service) AccountSetContentSettings(ctx context.Context, request *mtproto.TLAccountSetContentSettings) (*mtproto.Bool, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.setContentSettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountSetContentSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.setContentSettings - reply: %s", r)
	return r, err
}

// AccountGetContentSettings
// account.getContentSettings#8b9b4dae = account.ContentSettings;
func (s *Service) AccountGetContentSettings(ctx context.Context, request *mtproto.TLAccountGetContentSettings) (*mtproto.Account_ContentSettings, error) {
	c := core.New(ctx, s.svcCtx)
	c.Logger.Debugf("account.getContentSettings - metadata: %s, request: %s", c.MD, request)

	r, err := c.AccountGetContentSettings(request)
	if err != nil {
		return nil, err
	}

	c.Logger.Debugf("account.getContentSettings - reply: %s", r)
	return r, err
}
