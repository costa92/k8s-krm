package options

import (
	"fmt"
	"testing"

	"github.com/costa92/krm/pkg/server"
)

func Test_Apply(t *testing.T) {
	ops := NewSecureServingOptions()
	cfg := server.NewConfig()
	if err := ops.ApplyTo(cfg); err != nil {
		t.Fatal(err)
	}
	data := ops.Address()
	fmt.Println(data)
}
