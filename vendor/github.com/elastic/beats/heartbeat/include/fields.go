// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package include

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("heartbeat", "fields.yml", Asset); err != nil {
		panic(err)
	}
}

// Asset returns asset data
func Asset() string {
	return "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"
}