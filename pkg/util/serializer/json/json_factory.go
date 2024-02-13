package json

import "encoding/json"

type MetaFactory struct{}

// Interpret interprets the data
func (j *MetaFactory) Interpret(data []byte) error {
	return nil
}

// Encode encodes the object into data
func (j *MetaFactory) Encode(obj interface{}) ([]byte, error) {
	marshal, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	return marshal, nil
}

// Decode decodes the data into the object
func (j *MetaFactory) Decode(data []byte, obj interface{}) error {
	if err := json.Unmarshal(data, obj); err != nil {
		return err
	}
	return nil
}
