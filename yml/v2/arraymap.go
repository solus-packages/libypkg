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
	"errors"
	"gopkg.in/yaml.v3"
	"sort"
)

type ArrayMap map[string]string

var ErrInvalidMap = errors.New("ArrayMap must be a single string or an array of key value pairs")

const (
	DefaultPackage = "^"
)

func (m ArrayMap) MarshalYAML() (out interface{}, err error) {
	switch len(m) {
	case 0:
		err = ErrInvalidMap
	case 1:
		if len(m[DefaultPackage]) == 0 {
			err = ErrInvalidMap
			return
		}
		out = m[DefaultPackage]
	case 2:
		main := m[DefaultPackage]
		if len(main) == 0 {
			err = ErrInvalidMap
			return
		}
		nodes := make([]yaml.Node, 0)
		node := yaml.Node{
			Kind:  yaml.ScalarNode,
			Value: main,
		}
		nodes = append(nodes, node)
		var names []string
		for name := range m {
			if name != DefaultPackage {
				names = append(names, name)
			}
		}
		sort.Strings(names)
		for _, name := range names {
			node := yaml.Node{
				Kind: yaml.MappingNode,
			}
			key := yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: name,
			}
			value := yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: m[name],
			}
			node.Content = append(node.Content, &key, &value)
			nodes = append(nodes, node)
		}
		out = nodes
	}
	return
}

func (am *ArrayMap) UnmarshalYAML(value *yaml.Node) error {
	switch value.Kind {
	case yaml.ScalarNode:
		v := value.Value
		if len(v) == 0 {
			return ErrInvalidMap
		}
		*am = make(ArrayMap)
		(*am)[DefaultPackage] = v
	case yaml.SequenceNode:
		if len(value.Content) == 0 {
			return ErrInvalidMap
		}
		main := value.Content[0]
		if main.Kind != yaml.ScalarNode || len(main.Value) == 0 {
			return ErrInvalidMap
		}
		m := make(ArrayMap)
		m[DefaultPackage] = main.Value
		for _, node := range value.Content[1:] {
			if node.Kind != yaml.MappingNode || len(node.Content) != 2 {
				return ErrInvalidMap
			}
			k := node.Content[0]
			if k.Kind != yaml.ScalarNode || len(k.Value) == 0 || k.Value == DefaultPackage {
				return ErrInvalidMap
			}
			v := node.Content[1]
			if v.Kind != yaml.ScalarNode || len(v.Value) == 0 {
				return ErrInvalidMap
			}
			m[k.Value] = v.Value
		}
		*am = m
	default:
		return ErrInvalidMap
	}
	return nil
}
