package serializer

import (
	"fmt"
	"os"
	"testing"
)

func Test_Encode(t *testing.T) {
	metaFactory := NewYamlSerializer()
	test := Test{
		Server: Server{
			Mode:    "test",
			Healthz: true,
		},
	}
	encode, err := metaFactory.Encode(test)
	if err != nil {
		return
	}
	fmt.Println(string(encode))
}

type Test struct {
	Server Server `json:"server" yaml:"server" toml:"server"`
}

type Server struct {
	Mode    string `json:"mode" yaml:"mode" toml:"mode"`
	Healthz bool   `json:"healthz" yaml:"healthz" toml:"healthz"`
}

func openYaml() []byte {
	dataBytes, err := os.ReadFile("test.yaml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return nil
	}
	return dataBytes
}

func Test_Decode(t *testing.T) {
	metaFactory := NewYamlSerializer()
	test := Test{}
	dataBytes := openYaml()
	err := metaFactory.Decode(dataBytes, &test)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(test.Server)
}
