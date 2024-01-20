package main

import (
	"fmt"
	"github.com/costa92/krm/pkg/version"
	"time"
)

func main() {
	for {
		fmt.Println("Hello, world!")
		time.Sleep(10 * time.Second)
		version.CheckNewVersion()
	}

}
