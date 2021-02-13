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
	"dev.getsol.us/source/ypkg-source/spec/shared"
	"gopkg.in/yaml.v3"
	"os"
)

// PackageDeps includes the dependencies required for the Build, Check, and Run Stages
type PackageDeps struct {
	Build []string      `yaml:"builddeps,omitempty"`
	Run   []interface{} `yaml:"rundeps,omitempty"`
}

// Convert turns a v2 PackageDeps into the intermediate "shared" PackageDeps
func (p PackageDeps) Convert() shared.PackageDeps {
	d := shared.NewPackageDeps()
	// Set build dependencies if not empty
	if len(p.Build) > 0 {
		d.Build[shared.BaseLabel] = p.Build
	}
	// Figure out rundeps for all sub-packages
	baseRun := make([]string, 0)
	// For each entry in the rundeps
	for _, r := range p.Run {
		switch v1 := r.(type) {
		case string:
			// strings get added to the base rundeps
			// e.g:
			// - one
			// - two
			// - three
			baseRun = append(baseRun, v1)
		case map[string][]interface{}:
			// maps contain lists of rundeps for the sub-packages
			// e.g:
			// - sub:
			//		- one
			//		- two
			//		- three
			for k, v2 := range v1 {
				part := make([]string, 0)
				for _, s := range v2 {
					part = append(part, s.(string))
				}
				d.Run[k] = part
			}
		default:
			// nothing else is allowed here
			panic(v1)
		}
	}
	// set base package rundeps if not empty
	if len(baseRun) > 0 {
		d.Run[shared.BaseLabel] = baseRun
	}
	return d
}
