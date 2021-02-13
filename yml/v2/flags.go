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

// BuildFlags are special options that configure the build process
type BuildFlags struct {
	AutoDep    string   `yaml:"autodep,omitempty"`
	AVX2       string   `yaml:"avx2,omitempty"`
	Clang      string   `yaml:"clang,omitempty"`
	CCache     string   `yaml:"ccache,omitempty"`
	Debug      string   `yaml:"debug,omitempty"`
	Devel      string   `yaml:"devel,omitempty"`
	Emul32     string   `yaml:"emul32,omitempty"`
	Extract    string   `yaml:"extract,omitempty"`
	LAStrip    string   `yaml:"lastrip,omitempty"`
	Networking string   `yaml:"networking,omitempty"`
	Optimize   []string `yaml:"optimize,omitempty"`
	Strip      string   `yaml:"strip,omitempty"`
}

// Convert turns the v2 BuildFlags into the intermediate "shared" BuildFlags
func (f BuildFlags) Convert() shared.BuildFlags {
	fs := shared.BuildFlags{
		AutoDep:    f.AutoDep,
		AVX2:       f.AVX2,
		Clang:      f.Clang,
		CCache:     f.CCache,
		Debug:      f.Debug,
		Devel:      f.Devel,
		Emul32:     f.Emul32,
		Extract:    f.Extract,
		LAStrip:    f.LAStrip,
		Networking: f.Clang,
		Optimize:   f.Optimize,
		Strip:      f.Strip,
	}
	return fs
}
