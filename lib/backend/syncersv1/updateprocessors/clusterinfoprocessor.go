// Copyright (c) 2017 Tigera, Inc. All rights reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package updateprocessors

import (
	"reflect"

	apiv2 "github.com/projectcalico/libcalico-go/lib/apis/v2"
	"github.com/projectcalico/libcalico-go/lib/backend/model"
	"github.com/projectcalico/libcalico-go/lib/backend/watchersyncer"
)

// Create a new NewClusterInfoUpdateProcessor.
func NewClusterInfoUpdateProcessor() watchersyncer.SyncerUpdateProcessor {
	return NewConfigUpdateProcessor(
		reflect.TypeOf(apiv2.ClusterInformationSpec{}),
		DisallowAnnotations,
		func(node, name string) model.Key { return model.HostConfigKey{Hostname: node, Name: name} },
		func(name string) model.Key { return model.GlobalConfigKey{Name: name} },
		nil,
	)
}
