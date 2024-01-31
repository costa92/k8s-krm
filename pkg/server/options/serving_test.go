package options

import (
	"fmt"
	"testing"

	"github.com/costa92/krm/pkg/server"
)

func Test_Apply(t *testing.T) {
	ops := NewSecureServingOptions()
	cfg := server.NewConfig()
	cfg.SecureServing.Addr = "0.0.0.0"
	cfg.SecureServing.Port = 8000
	err := ops.ApplyTo()
	if err != nil {
		panic(err)
	}
	data := ops.Address()
	fmt.Println(data)
}
