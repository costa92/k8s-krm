package plugins

import "io"

type Factor func(config io.Reader) (Interface, error)
