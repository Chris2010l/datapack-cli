package src

import (
	"fmt"

	"github.com/goccy/go-yaml"
)

type DatapackConfig struct {
	Package string `yaml:"package"`
	Steps   []Step `yaml:"steps"`
}

type Step struct {
	Goto   string   `yaml:"goto,omitempty"`
	Create []Create `yaml:"create,omitempty"`
}

type Create struct {
	Path   string `yaml:"path"`
	Type   string `yaml:"type"`
	Import string `yaml:"import"`
}

func ParseYaml(data_structure []byte) (DatapackConfig, error) {

	var unmarshaled_data DatapackConfig

	if err := yaml.Unmarshal([]byte(data_structure), &unmarshaled_data); err != nil {
		fmt.Printf("%s", err)
		return DatapackConfig{}, err
	}
	return unmarshaled_data, nil
}
