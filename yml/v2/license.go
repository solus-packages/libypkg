package v2

import (
	"gopkg.in/yaml.v3"
)

type Licenses []yaml.Node

var ErrNotLicense = errors.New("Licenses must be a single string or an array of strings")

func (l Licenses) MarshalYAML() (out interface{}, err error) {
	switch len(l) {
	case 0:
		err = ErrNotLicense
	case 1:
		if l[0].Kind != yaml.ScalarNode {
			err = ErrNotLicense
		}
		out = l[0]
	case 2:
		for _, node := range l {
			if node.Kind != yaml.ScalarNode {
				err = ErrNotLicense
			}
		}
		out = []yaml.Node(l)
	}
	return
}
