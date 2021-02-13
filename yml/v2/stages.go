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

// BuildStages represent the scripted commands to execute for each stage of the build process
type BuildStages struct {
	Setup   string `yaml:"setup,omitempty"`
	Build   string `yaml:"build,omitempty"`
	Profile string `yaml:"profile,omitempty"`
	Check   string `yaml:"check,omitempty"`
	Install string `yaml:"install"`
}

// Convert turns a v2 BuildStages into the intermediate "shared" BuildStages
func (s BuildStages) Convert() shared.BuildStages {
	return shared.BuildStages{
		Setup:   s.Setup,
		Build:   s.Build,
		Profile: s.Profile,
		Check:   s.Check,
		Install: s.Install,
	}
}
