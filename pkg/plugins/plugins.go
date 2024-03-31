package plugins

import "io"

type Factor func(config io.Reader) (Interface, error)

func Register(name string, f Factor) {
	//registry[name] = f
}
