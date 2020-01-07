package descriptor

import (
	jsoniter "github.com/json-iterator/go"
	"gopkg.in/yaml.v3"
)

/*
 * Deserialize server description based on JSON
 */

func ParseJson(b []byte) (Service, error) {
	var s Service
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	err := json.Unmarshal(b, &s)
	return s, err
}

func ParseYaml(b []byte) (Service, error) {
	var s Service
	err := yaml.Unmarshal(b, &s)
	return s, err
}
