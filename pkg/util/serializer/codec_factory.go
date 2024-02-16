package serializer

import (
	"github.com/costa92/krm/pkg/util/serializer/json"
	"github.com/costa92/krm/pkg/util/serializer/toml"
	"github.com/costa92/krm/pkg/util/serializer/yaml"
)

type CodecFactory struct {
	Coder Codec
}

type CodecFactoryOptions struct {
	MetaType MeteType
}

// Encode encodes the object into data
func (c *CodecFactory) Encode(obj interface{}) ([]byte, error) {
	return c.Coder.Encode(obj)
}

// Decode decodes the data into the object
func (c *CodecFactory) Decode(data []byte, obj interface{}) error {
	return c.Coder.Decode(data, obj)
}

// NewCodecFactory returns a new codec factory
func NewCodecFactory(opts CodecFactoryOptions) *CodecFactory {
	coder := getMetaFactory(opts.MetaType)
	return &CodecFactory{
		Coder: coder,
	}
}

// MeteType represents the type of meta
var metaFactoryMap = map[MeteType]Codec{
	MeteTypeYaml: &yaml.MetaFactory{},
	MeteTypeJSON: &json.MetaFactory{},
	MeteTypeToml: &toml.MetaFactory{},
}

// getMetaFactory returns the meta factory
func getMetaFactory(metaType MeteType) Codec {
	return metaFactoryMap[metaType]
}

// NewJSONSerializer returns a new json serializer
func NewJSONSerializer() *CodecFactory {
	return NewCodecFactory(CodecFactoryOptions{MetaType: MeteTypeJSON})
}

// NewYamlSerializer returns a new yaml serializer
func NewYamlSerializer() *CodecFactory {
	return NewCodecFactory(CodecFactoryOptions{MetaType: MeteTypeYaml})
}

// NewTomlSerializer returns a new toml serializer
func NewTomlSerializer() *CodecFactory {
	return NewCodecFactory(CodecFactoryOptions{MetaType: MeteTypeToml})
}
