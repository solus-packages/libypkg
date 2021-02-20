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

func TestPatternsMarshalEmpty(t *testing.T) {
	expected := "name: golang\n"
	var out strings.Builder
	enc := yaml.NewEncoder(&out)
	value := struct {
		Name     string       `yaml:"name"`
		Patterns ArrayListMap `yaml:",omitempty"`
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

func TestPatternsMarshalSingle(t *testing.T) {
	expected := `name: golang
patterns:
    - /usr/bin/go
    - /usr/share/golang
`
	value := struct {
		Name     string       `yaml:"name"`
		Patterns ArrayListMap `yaml:",omitempty"`
	}{
		Name: "golang",
		Patterns: ArrayListMap{
			DefaultPackage: []string{
				"/usr/bin/go",
				"/usr/share/golang",
			},
		},
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

func TestPatternsMarshalMultiple(t *testing.T) {
	expected := `name: golang
patterns:
    - /usr/bin/go
    - /usr/share/golang
    - devel:
        - /usr/include
        - /usr/share/golang/include
`
	value := struct {
		Name     string       `yaml:"name"`
		Patterns ArrayListMap `yaml:",omitempty"`
	}{
		Name: "golang",
		Patterns: ArrayListMap{
			DefaultPackage: []string{
				"/usr/bin/go",
				"/usr/share/golang",
			},
			"devel": []string{
				"/usr/include",
				"/usr/share/golang/include",
			},
		},
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

func TestPatternsUnmarshalEmpty(t *testing.T) {
	input := "name: golang"
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	value := struct {
		Name     string       `yaml:"name"`
		Patterns ArrayListMap `yaml:",omitempty"`
	}{
		Patterns: make(ArrayListMap),
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if value.Name != "golang" {
		t.Fatalf("expected '%s', found: %s", "golang", value.Name)
	}
}

func TestPatternsUnmarshalSingle(t *testing.T) {
	input := `name: golang
patterns:
    - /usr/bin/go
    - /usr/share/golang
`
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	value := struct {
		Name     string       `yaml:"name"`
		Patterns ArrayListMap `yaml:",omitempty"`
	}{
		Patterns: make(ArrayListMap),
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if value.Name != "golang" {
		t.Fatalf("expected '%s', found: %s", "golang", value.Name)
	}
	if len(value.Patterns) != 1 {
		t.Fatalf("expected 1 package, found: %d", len(value.Patterns))
	}
	paths := value.Patterns[DefaultPackage]
	if len(paths) != 2 {
		t.Fatalf("expected 2 paths, found: %d", len(paths))
	}
	if paths[0] != "/usr/bin/go" {
		t.Errorf("expected '%s', found: %s", "/usr/bin/go", paths[0])
	}
	if paths[1] != "/usr/share/golang" {
		t.Errorf("expected '%s', found: %s", "/usr/share/golang", paths[0])
	}
}

func TestPatternsUnmarshalMultiple(t *testing.T) {
	input := `name: golang
patterns:
    - /usr/bin/go
    - /usr/share/golang
    - devel:
        - /usr/include
        - /usr/share/golang/include
`
	in := strings.NewReader(input)
	dec := yaml.NewDecoder(in)
	value := struct {
		Name     string       `yaml:"name"`
		Patterns ArrayListMap `yaml:",omitempty"`
	}{
		Patterns: make(ArrayListMap),
	}
	if err := dec.Decode(&value); err != nil {
		t.Fatalf("Expected no error, found: %s", err)
	}
	if value.Name != "golang" {
		t.Fatalf("expected '%s', found: %s", "golang", value.Name)
	}
	if len(value.Patterns) != 2 {
		t.Fatalf("expected 2 package, found: %d", len(value.Patterns))
	}
	paths := value.Patterns[DefaultPackage]
	if len(paths) != 2 {
		t.Fatalf("expected 2 paths, found: %d", len(paths))
	}
	if paths[0] != "/usr/bin/go" {
		t.Errorf("expected '%s', found: %s", "/usr/bin/go", paths[0])
	}
	if paths[1] != "/usr/share/golang" {
		t.Errorf("expected '%s', found: %s", "/usr/share/golang", paths[0])
	}
	paths = value.Patterns["devel"]
	if len(paths) != 2 {
		t.Fatalf("expected 2 paths, found: %d", len(paths))
	}
	if paths[0] != "/usr/include" {
		t.Errorf("expected '%s', found: %s", "/usr/include", paths[0])
	}
	if paths[1] != "/usr/share/golang/include" {
		t.Errorf("expected '%s', found: %s", "/usr/share/golang/include", paths[0])
	}
}