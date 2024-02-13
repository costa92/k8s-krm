package serializer

import (
	"fmt"
	"testing"
)

// Path: pkg/util/serializer/serializer.go
// Compare this snippet from pkg/util/serializer/serializer.go:
// package serializer
//

type testMetaFactory struct{}

func (f *testMetaFactory) Interpret(data []byte) error {
	return nil
}

func (f *testMetaFactory) Encode(obj interface{}) ([]byte, error) {
	fmt.Println(obj)
	return nil, nil
}

func (f *testMetaFactory) Decode(data []byte, obj interface{}) error {
	return nil
}

func Test_Run(t *testing.T) {
	opts := CodecFactoryOptions{
		MetaType: MeteTypeYaml,
	}
	codeFactory := NewCodecFactory(opts)
	if codeFactory.Coder == nil {
		t.Errorf("Expected a non-nil value, got nil")
	}
	codeFactory.Encode("test")
}
