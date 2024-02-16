package yaml

import (
	"gopkg.in/yaml.v3"
)

type MetaFactory struct{}

// Interpret interprets the data
func (y *MetaFactory) Encode(obj interface{}) ([]byte, error) {
	marshal, err := yaml.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

// Decode decodes the data into the object
func (y *MetaFactory) Decode(data []byte, obj interface{}) error {
	return yaml.Unmarshal(data, obj)
}
