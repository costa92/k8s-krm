package serializer

type MeteType string

const (
	MeteTypeYaml MeteType = "yaml"
	MeteTypeJSON MeteType = "json"
	MeteTypeToml MeteType = "toml"
)

type Encoder interface {
	Encode(obj interface{}) ([]byte, error)
}

type Decoder interface {
	Decode(data []byte, obj interface{}) error
}

type Serializer interface {
	Encoder
	Decoder
}

type (
	Codec          Serializer
	JSONSerializer Serializer // nolint:golint,unused
	YamlSerializer Serializer // nolint:golint,unused
	TomlSerializer Serializer // nolint:golint,unused
)
