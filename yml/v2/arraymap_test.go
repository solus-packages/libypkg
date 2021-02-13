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

func TestArrayMapMarshalEmpty(t *testing.T) {
	var am ArrayMap
	var out strings.Builder
	enc := yaml.NewEncoder(&out)
	if err := enc.Encode(am); err != ErrInvalidMap {
		t.Fatalf("Expected ErrNotLicense, found: %s", err)
	}
	if result := out.String(); len(result) != 0 {
		t.Fatalf("Expected empty result, found: %s", result)
	}
}

func TestArrayMapMarshalSingle(t *testing.T) {
	expected := "component: system.devel\n"
	am := make(ArrayMap)
	am[DefaultPackage] = "system.devel"
	value := struct {
		Components ArrayMap `yaml:"component"`
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

func TestArrayMapMarshalMultiple(t *testing.T) {
	expected := `component:
    - system.devel
    - devel: programming.devel
`
	am := make(ArrayMap)
	am[DefaultPackage] = "system.devel"
	am["devel"] = "programming.devel"
	value := struct {
		Components ArrayMap `yaml:"component"`
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

func TestArrayMapUnmarshalEmpty(t *testing.T) {
	input := "component:"
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		Components ArrayMap `yaml:"component"`
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if l := len(value.Components); l != 0 {
		t.Fatalf("expected exactly 0 entries, found: %d", l)
	}
}

func TestArrayMapUnmarshalSingle(t *testing.T) {
	input := "component: system.devel"
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		Components ArrayMap `yaml:"component"`
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

func TestArrayMapUnmarshalMultiple(t *testing.T) {
	input := `component:
     - system.devel
     - devel: programming.devel
`
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	var value struct {
		Components ArrayMap `yaml:"component"`
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
