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
)

// BuildFlags are special options that configure the build process
type BuildFlags struct {
	AutoDep    DefaultTrue  `yaml:"autodep,omitempty"`
	AVX2       DefaultFalse `yaml:"avx2,omitempty"`
	Clang      DefaultTrue  `yaml:"clang,omitempty"`
	CCache     DefaultFalse `yaml:"ccache,omitempty"`
	Debug      DefaultTrue  `yaml:"debug,omitempty"`
	Devel      DefaultFalse `yaml:"devel,omitempty"`
	Emul32     DefaultFalse `yaml:"emul32,omitempty"`
	Extract    DefaultTrue  `yaml:"extract,omitempty"`
	LAStrip    DefaultTrue  `yaml:"lastrip,omitempty"`
	Networking DefaultFalse `yaml:"networking,omitempty"`
	Optimize   []string     `yaml:"optimize,omitempty"`
	Strip      DefaultTrue  `yaml:"strip,omitempty"`
}

var (
	ErrNotABool = errors.New("Not a valid boolean string")
)

type DefaultTrue struct {
	Valid bool
	Bool  bool
}

func (dt DefaultTrue) MarshalYAML() (out interface{}, err error) {
	node := yaml.Node{
		Kind: yaml.ScalarNode,
	}
	if dt.Valid && !dt.Bool {
		node.Value = "no"
	}
	out = node
	return
}

func (dt *DefaultTrue) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return ErrNotABool
	}
	switch value.Value {
	case "no", "NO", "No", "False", "false":
		(*dt).Valid = true
		(*dt).Bool = false
	case "yes", "YES", "Yes", "True", "true":
		(*dt).Valid = true
		(*dt).Bool = true
	default:
		return ErrNotABool
	}
	return nil
}

type DefaultFalse struct {
	Valid bool
	Bool  bool
}

func (df DefaultFalse) MarshalYAML() (out interface{}, err error) {
	node := yaml.Node{
		Kind: yaml.ScalarNode,
	}
	if df.Valid && df.Bool {
		node.Value = "yes"
	}
	out = node
	return
}

func (df *DefaultFalse) UnmarshalYAML(value *yaml.Node) error {
	if value.Kind != yaml.ScalarNode {
		return ErrNotABool
	}
	switch value.Value {
	case "no", "NO", "No", "False", "false":
		(*df).Valid = true
		(*df).Bool = false
	case "yes", "YES", "Yes", "True", "true":
		(*df).Valid = true
		(*df).Bool = true
	default:
		return ErrNotABool
	}
	return nil
}
