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

type ArrayListMap map[string][]string

var ErrInvalidListMap = errors.New("ArrayListMap must be a list of strings or a map of lists of strings")

func (m ArrayListMap) MarshalYAML() (out interface{}, err error) {
	if len(m) == 0 {
		err = ErrInvalidListMap
		return
	}
	nodes := make([]yaml.Node, 0)
	main := m[DefaultPackage]
	if len(main) > 0 {
		for _, pkg := range main {
			node := yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: pkg,
			}
			nodes = append(nodes, node)
		}
	}
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
			Kind: yaml.SequenceNode,
		}
		for _, pkg := range m[name] {
			child := yaml.Node{
				Kind:  yaml.ScalarNode,
				Value: pkg,
			}
			value.Content = append(value.Content, &child)
		}
		node.Content = append(node.Content, &key, &value)
		nodes = append(nodes, node)
	}
	out = nodes
	return
}

func (am *ArrayListMap) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.SequenceNode {
		return ErrInvalidListMap
	}
	if len(value.Content) == 0 {
		return ErrInvalidListMap
	}
	for _, node := range value.Content {
		switch node.Kind {
		case yaml.ScalarNode:
			(*am)[DefaultPackage] = append((*am)[DefaultPackage], node.Value)
		case yaml.MappingNode:
			if len(node.Content) != 2 {
				return ErrInvalidListMap
			}
			k := node.Content[0]
			if k.Kind != yaml.ScalarNode || len(k.Value) == 0 || k.Value == DefaultPackage {
				return ErrInvalidListMap
			}
			v := node.Content[1]
			if v.Kind != yaml.SequenceNode || len(v.Content) == 0 {
				return ErrInvalidListMap
			}
			for _, n := range v.Content {
				if n.Kind != yaml.ScalarNode || len(n.Value) == 0 {
					return ErrInvalidListMap
				}
				(*am)[k.Value] = append((*am)[k.Value], n.Value)
			}
		default:
			return ErrInvalidListMap
		}
	}
	return nil
}
