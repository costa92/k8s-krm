package toml

import (
	"bytes"

	"github.com/BurntSushi/toml"
)

type MetaFactory struct{}

// Interpret interprets the data
// onlint:revive
func (t *MetaFactory) Interpret(_ []byte) error {
	return nil
}

// Encode encodes the object into data
func (t *MetaFactory) Encode(obj interface{}) ([]byte, error) {
	var buf bytes.Buffer
	err := toml.NewEncoder(&buf).Encode(obj)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// Decode decodes the data into the object
func (t *MetaFactory) Decode(data []byte, obj interface{}) error {
	err := toml.Unmarshal(data, obj)
	if err != nil {
		return err
	}
	return nil
}
