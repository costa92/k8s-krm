package serializer

import "testing"

// Path: pkg/runtime/serializer/codec_factory_test.go

type testMetaFactory struct{}

func (f *testMetaFactory) Interpret(data []byte) error {
	return nil
}

func Test_Run(t *testing.T) {
}
