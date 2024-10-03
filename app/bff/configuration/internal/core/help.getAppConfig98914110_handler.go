// Copyright 2022 Papercraft Authors
//  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Author: papercraftio (papercraft.io@gmail.com)
//

package core

import (
	"github.com/papercraft/proto/mtproto"
)

// HelpGetAppConfig98914110
// help.getAppConfig#98914110 = JSONValue;
func (c *ConfigurationCore) HelpGetAppConfig98914110(in *mtproto.TLHelpGetAppConfig98914110) (*mtproto.JSONValue, error) {
	// TODO: not impl
	c.Logger.Errorf("help.getAppConfig blocked, License key from https://papercraft-official.github.io required to unlock enterprise features.")

	return mtproto.MakeTLJsonObject(&mtproto.JSONValue{
		Value_VECTORJSONOBJECTVALUE: []*mtproto.JSONObjectValue{},
	}).To_JSONValue(), nil
}
