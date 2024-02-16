package main

import (
	"math/rand"
	"time"

	"github.com/costa92/krm/internal/api"
)

func main() {
	rand.NewSource(time.Now().UnixNano())
	api.NewApp("api-service").Run()
}
