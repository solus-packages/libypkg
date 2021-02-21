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

// PackageDeps includes the dependencies required for the Build, Check, and Run Stages
type PackageDeps struct {
	Replaces  ArrayListMap `yaml:"replaces,omitempty"`
	Conflicts ArrayListMap `yaml:"conflicts,omitempty"`
	Build     []string     `yaml:"builddeps,omitempty"`
	Run       ArrayListMap `yaml:"rundeps,omitempty"`
}
