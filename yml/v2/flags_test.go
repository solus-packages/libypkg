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
	"strings"
	"testing"
)

func TestBuildFlagsMarshalEmpty(t *testing.T) {
	expected := "name: golang\n"
	var out strings.Builder
	enc := yaml.NewEncoder(&out)
	value := struct {
		Name  string     `yaml:"name"`
		Flags BuildFlags `yaml:",omitempty,inline"`
	}{
		Name: "golang",
	}
	if err := enc.Encode(value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if result := out.String(); result != expected {
		t.Fatalf("Expected '%s' result, found: %s", expected, result)
	}
}

func TestBuildFlagsMarshalSingle(t *testing.T) {
	expected := `name: golang
autodep: no
`
	value := struct {
		Name  string     `yaml:"name"`
		Flags BuildFlags `yaml:",omitempty,inline"`
	}{
		Name: "golang",
		Flags: BuildFlags{
			AutoDep: "no",
		},
	}
	var out strings.Builder
	enc := yaml.NewEncoder(&out)
	if err := enc.Encode(value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	println(out.String())
	if result := out.String(); result != expected {
		t.Fatalf("Expected %s, found: %s", expected, result)
	}
}

/*
func TestBuildFlagsMarshalMultiple(t *testing.T) {
	expected := `component:
    - system.devel
    - devel: programming.devel
`
	am := make(BuildFlags)
	am[DefaultPackage] = "system.devel"
	am["devel"] = "programming.devel"
	value := struct {
		Components BuildFlags `yaml:"component"`
	}{
		Components: am,
	}
	var out strings.Builder
	enc := yaml.NewEncoder(&out)
	if err := enc.Encode(value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if result := out.String(); result != expected {
		t.Fatalf("Expected %s, found: %s", expected, result)
	}
}

func TestBuildFlagsUnmarshalEmpty(t *testing.T) {
	input := "component:"
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		Components BuildFlags `yaml:"component"`
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if l := len(value.Components); l != 0 {
		t.Fatalf("expected exactly 0 entries, found: %d", l)
	}
}

func TestBuildFlagsUnmarshalSingle(t *testing.T) {
	input := "component: system.devel"
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		Components BuildFlags `yaml:"component"`
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if l := len(value.Components); l != 1 {
		t.Fatalf("expected exactly one entry, found: %d", l)
	}
	if v := value.Components[DefaultPackage]; v != "system.devel" {
		t.Errorf("expected 'system.devel', found: %s", v)
	}
}

func TestBuildFlagsUnmarshalMultiple(t *testing.T) {
	input := `component:
     - system.devel
     - devel: programming.devel
`
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		Components BuildFlags `yaml:"component"`
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if l := len(value.Components); l != 2 {
		t.Fatalf("expected exactly two entries, found: %d", l)
	}
	if v := value.Components[DefaultPackage]; v != "system.devel" {
		t.Errorf("expected 'system.devel', found: %s", v)
	}
	if v := value.Components["devel"]; v != "programming.devel" {
		t.Errorf("expected 'programming.devel', found: %s", v)
	}
}
*/
