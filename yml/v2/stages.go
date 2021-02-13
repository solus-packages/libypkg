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
