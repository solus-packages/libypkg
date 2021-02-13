//
// Copyright © 2021 Solus Project <copyright@getsol.us>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package v2

import (
	"gopkg.in/yaml.v3"
)

type Licenses []yaml.Node

var ErrNotLicense = errors.New("Licenses must be a single string or an array of strings")

func (l Licenses) MarshalYAML() (out interface{}, err error) {
	switch len(l) {
	case 0:
		err = ErrNotLicense
	case 1:
		if l[0].Kind != yaml.ScalarNode {
			err = ErrNotLicense
		}
		out = l[0]
	case 2:
		for _, node := range l {
			if node.Kind != yaml.ScalarNode {
				err = ErrNotLicense
			}
		}
		out = []yaml.Node(l)
	}
	return
}
