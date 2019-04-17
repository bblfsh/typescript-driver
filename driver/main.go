package main

import (
	_ "github.com/bblfsh/typescript-driver/driver/impl"
	"github.com/bblfsh/typescript-driver/driver/normalizer"

	"github.com/bblfsh/sdk/v3/driver/server"
)

func main() {
	server.Run(normalizer.Transforms)
}
