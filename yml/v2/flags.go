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
