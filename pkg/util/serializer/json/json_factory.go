package json

import "encoding/json"

type MetaFactory struct{}

// Interpret interprets the data
func (j *MetaFactory) Interpret(_ []byte) error {
	return nil
}

// Encode encodes the object into data
func (j *MetaFactory) Encode(obj interface{}) ([]byte, error) {
	return json.Marshal(obj)
}

// Decode decodes the data into the object
func (j *MetaFactory) Decode(data []byte, obj interface{}) error {
	return json.Unmarshal(data, obj)
}
