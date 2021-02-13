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
	"os"
)

// PackageYML is the v3 representation of the Package YML specification
type PackageYML struct {
	Name         string      `yaml:"name"`
	Version      string      `yaml:"version"`
	Release      uint        `yaml:"release"`
	Source       []Source    `yaml:"source"`
	Homepage     string      `yaml:"homepage,omitempty"`
	License      Licenses    `yaml:"license"`
	Component    ArrayMap    `yaml:"component"`
	Summary      ArrayMap    `yaml:"summary"`
	Description  ArrayMap    `yaml:"description"`
	Flags        BuildFlags  `yaml:",omitempty,inline"`
	Environment  string      `yaml:"environment,omitempty"`
	Dependencies PackageDeps `yaml:"dependencies,inline,omitempty"`
	Stages       BuildStages `yaml:",inline"`
	Patterns     interface{} `yaml:"patterns,omitempty"`
	f            *os.File
}

// Load populates a v2 PackageYML by reading in the contents from a specific filepath
func (p *PackageYML) Load(path string, mode int) error {
	// p.Flags = DefaultBuildFlags
	f, err := os.OpenFile(path, mode, 00644)
	if err != nil {
		return err
	}
	p.f = f
	dec := yaml.NewDecoder(p.f)
	return dec.Decode(p)
}

// Save writes any changes to this PackageYML to the currently open file descriptor
func (p *PackageYML) Save() error {
	// Seek to the beginning of the file
	_, err := p.f.Seek(0, 0)
	if err != nil {
		return err
	}
	// clear the contents of the file
	err = p.f.Truncate(0)
	if err != nil {
		return err
	}
	// write out the new contents to the file
	enc := yaml.NewEncoder(p.f)
	return enc.Encode(&p)
}

// Close close the file descriptor for this PackageYML
func (p *PackageYML) Close() {
	_ = p.f.Close()
}
