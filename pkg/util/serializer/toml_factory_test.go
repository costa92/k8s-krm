package serializer

import (
	"fmt"
	"os"
	"testing"
)

func Test_tomlEncode(t *testing.T) {
	metaFactory := NewTomlSerializer()
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

func openToml() []byte {
	dataBytes, err := os.ReadFile("test.toml")
	if err != nil {
		fmt.Println("读取文件失败：", err)
		return nil
	}
	return dataBytes
}

func Test_TomlDecode(t *testing.T) {
	metaFactory := NewTomlSerializer()
	test := Test{}
	dataBytes := openToml()
	err := metaFactory.Decode(dataBytes, &test)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(test.Server)
}
