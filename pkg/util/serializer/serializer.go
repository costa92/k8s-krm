package serializer

type MeteType string

const (
	MeteTypeYaml MeteType = "yaml"
	MeteTypeJson MeteType = "json"
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
	JsonSerializer Serializer
	YamlSerializer Serializer
	TomlSerializer Serializer
)
